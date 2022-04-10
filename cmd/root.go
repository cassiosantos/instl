package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/mholt/archiver/v3"
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/installer/instl/internal"
)

var rootCmd = &cobra.Command{
	Use:   "instl [username/repo]",
	Short: "Instl is an installer that can install most GitHub projects on your system with a single command.",
	Long: `Instl is an installer that can install most GitHub projects on your system with a single command.  
Additionally, Instl provides a server that generates dynamic scripts that install a GitHub project.  

Official docs: https://docs.instl.sh

## Web Installer

> The web install command can be used by anyone and does not require anything to be installed.
> Running the web install command will download instl and install the given GitHub project.
> After that, instl will be removed from the system again.

The instl web installer is a single command, which everyone can run, to install a GitHub project.
This is the basic syntax, which will return an install script from our API server:

	                     ┌ The GitHub username of the project
	                     |          ┌ The GitHub repository name of the project
	                     |          |         ┌ The platform, see "Valid Platforms"
	                     |          |         |
	https://instl.sh/:username/:reponame/:platform

### Valid Platforms

| Valid Platforms | Parameter |
|-----------------|-----------|
|     Windows     |  windows |
|      macOS      |  macos  |
|      Linux      |  linux  |

### Running the web installer command

> Different operating systems need different commands to download and run the web installer script.

#### Windows

This command will download and execute the web installer script for windows.
You have to execute it in a powershell terminal.

	iwr -useb instl.sh/:username/:reponame/windows | iex

#### macOS

This command will download and execute the web installer script for macOS.

	curl -fsSL instl.sh/:username/:reponame/macos | bash

#### Linux

This command will download and execute the web installer script for linux.

	curl -fsSL instl.sh/:username/:reponame/linux | bash
`,
	Version: "v1.9.0", // <---VERSION---> This comment enables auto-releases on version change!
	Example: "instl installer/instl",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a GitHub repo to install\nExample: instl user/repo")
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		disableOutput, _ := cmd.PersistentFlags().GetBool("silent")

		if disableOutput {
			pterm.DisableOutput()
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get GitHub username and repository from args.
		repoArg := args[0]

		repoArg = strings.TrimPrefix(repoArg, "https://github.com/")
		repoArg = strings.TrimPrefix(repoArg, "github.com/")

		repoArgParts := strings.Split(repoArg, "/")
		if len(repoArgParts) != 2 {
			return fmt.Errorf("%s is not a valid GitHub repository", repoArg)
		}
		repoName := repoArgParts[len(repoArgParts)-2] + "/" + repoArgParts[len(repoArgParts)-1]

		// Print instl header.
		introText, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithStyle("  INSTL", pterm.NewStyle(pterm.FgMagenta))).Srender()
		pterm.Println()
		pterm.Println(strings.TrimRight(introText, "\n"))
		pterm.Printf(pterm.Cyan("                > https://instl.sh\n"))
		pterm.Printf(pterm.Blue("                       %s\n\n"), cmd.Version)
		pterm.Info.Printfln("instl.sh is an automated installer for GitHub projects.")
		if repoArgParts[0] != "installer" {
			pterm.Info.Printfln("We do not own https://github.com/%s.", pterm.Magenta(repoName))
		}
		pterm.Println()
		pterm.DefaultHeader.Printf("Running installer for github.com/%s", repoName)
		pterm.Println()

		// Request latest GitHub asset and it's assets.
		gettingAssetSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone().Start("Getting asset metadata from latest release...")
		repoTmp, err := internal.ParseRepository(repoArg)
		internal.Repo = repoTmp
		if err != nil {
			return err
		}
		var assetCount int
		internal.Repo.ForEachAsset(func(release internal.Asset) {
			assetCount++
		})
		gettingAssetSpinner.Stop()
		pterm.Debug.Sprintf("Found %d assets in latest release!", assetCount)

		// Detect right asset for system.
		detectAssetSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone().Start("Detecting right asset for machine...")
		var asset internal.Asset
		pterm.Debug.Println("Your system:", runtime.GOOS, runtime.GOARCH)
		asset, err = internal.DetectRightAsset(internal.Repo)
		if err != nil {
			return err
		}
		detectAssetSpinner.Stop()

		// Print asset stats.
		assetData := pterm.TableData{
			{"Name", "Version", "Last Update", "Size"},
			{asset.Name, pterm.Sprint(asset.Version), asset.UpdatedAt.Format("02 Jan 2006"), internal.ReadbleSize(asset.Size)},
		}
		if pterm.PrintDebugMessages {
			assetData[0] = append(assetData[0], "Score")
			assetData[1] = append(assetData[1], pterm.Sprint(asset.Score))
		}
		assetStats, _ := pterm.DefaultTable.WithHasHeader().WithData(assetData).Srender()
		pterm.DefaultBox.Println(assetStats)

		// Making installation ready.
		installPath := internal.GetInstallPath(internal.Repo.Name) + "/" + asset.Name
		installDir := internal.GetInstallPath(internal.Repo.Name)
		pterm.Debug.Printfln("InstallPath: %s\nInstallDir: %s", installPath, installDir)
		pterm.Debug.PrintOnError(os.RemoveAll(installDir))
		pterm.Warning.PrintOnError(os.MkdirAll(installDir, 0755))

		// Downloading asset.
		err = internal.DownloadFile(installPath, asset.DownloadURL)
		if err != nil {
			return err
		}
		pterm.Debug.Printf("Downloaded %s\n", asset.Name)

		// Installing asset.
		err = archiver.Unarchive(installPath, installDir)
		if err != nil {
			pterm.Debug.Println("Could not unarchive asset.\nTrying to install it directly.")
		} else {
			pterm.Warning.PrintOnError(os.Remove(installPath))
		}
		internal.AddToPath(installDir, internal.Repo.Name)

		// Success message.
		pterm.Println() // blank line
		pterm.Success.Printfln("You might have to restart your terminal session to use %s.", pterm.Magenta(internal.Repo.Name))

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		pcli.CheckForUpdates()
		os.Exit(1)
	}

	pcli.CheckForUpdates()
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empry strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "d", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "only outputs errors")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("installer/instl")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
	pterm.Error = *pterm.Error.WithShowLineNumber(false)
}
