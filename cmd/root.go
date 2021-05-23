package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mholt/archiver/v3"
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/instl-sh/instl/internal"
)

var rootCmd = &cobra.Command{
	Use:   "instl [username/repo]",
	Short: "Instl is an installer that installs GitHub projects on your system with a single command.",
	Long: `Instl is an installer that installs GitHub projects on your system with a single command.  
Additionally, Instl provides a server that generates dynamic scripts that install a GitHub project.  

Official docs: https://docs.instl.sh

To use the server you can use the following commands:
  
**Windows**  

    iwr instl.sh/username/reponame/windows | iex  
  
**macOS**  

    curl -sSL instl.sh/username/reponame/macos | sudo bash   
  
**Linux**  

    curl -sSL instl.sh/username/reponame/linux | sudo bash  
  
(Replace username and reponame with the GitHub project you want to install)  

Read more about the web installer here: https://docs.instl.sh/#/web-installer
  
These commands can be executed from any system and install the respective GitHub project.  
You can also provide these commands to your users to make your GitHub project easily installable.`,
	Version: "v1.5.0", // <---VERSION---> This comment enables auto-releases on version change!
	Example: "instl instl-sh/instl",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a GitHub repo to install\nExample: instl user/repo")
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if runtime.GOOS != "windows" {
			repoArg := args[0]

			repoArg = strings.TrimPrefix(repoArg, "https://github.com/")
			repoArg = strings.TrimPrefix(repoArg, "github.com/")
			repoArgParts := strings.Split(repoArg, "/")

			pterm.Info.Printfln("Instl needs administrative permissions to write to %s and %s.\n"+
				"If you have installed instl, you can use: "+pterm.Green("sudo instl %s")+"\n"+
				"If you used the web installer, you can use "+pterm.Green("curl -fsSL instl.sh/%s/%s | sudo bash"), pterm.Magenta("/usr/local/lib"), pterm.Magenta("/usr/local/bin"),
				strings.Join(repoArgParts, "/"), strings.Join(repoArgParts, "/"), runtime.GOOS)
			return errors.New("permission denied")
		}

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
		pterm.Printf(pterm.Cyan("                       %s\n\n"), cmd.Version)
		pterm.Info.Printf("instl.sh is an automated installer for GitHub projects.\nWe do not own https://github.com/%s.\n", pterm.Magenta(repoName))
		pterm.Println()
		pterm.DefaultHeader.Printf("Running installer for github.com/%s", repoName)
		pterm.Println()

		// Request latest GitHub asset and it's assets.
		err := internal.MakeSpinner("Getting asset metadata from latest release...", func() (string, error) {
			repoTmp, err := internal.ParseRepository(repoArg)
			internal.Repo = repoTmp
			if err != nil {
				return "", err
			}
			var assetCount int
			internal.Repo.ForEachAsset(func(release internal.Asset) {
				assetCount++
			})

			return pterm.Sprintf("Found %d assets in latest release!", assetCount), nil
		})
		if err != nil {
			return err
		}

		// Detect right asset for system.
		var asset internal.Asset
		err = internal.MakeSpinner("Detecting right asset for machine...", func() (string, error) {
			pterm.Debug.Println("Your system:", runtime.GOOS, runtime.GOARCH)
			asset = internal.DetectRightAsset(internal.Repo)
			return pterm.Sprintf("Found an asset which seems to fit to your system:"), nil
		})
		if err != nil {
			return err
		}

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
		installPath := internal.GetInstallPath(internal.Repo.User, internal.Repo.Name) + "/" + asset.Name
		installDir := internal.GetInstallPath(internal.Repo.User, internal.Repo.Name)
		pterm.Debug.PrintOnError(os.RemoveAll(installDir))
		pterm.Warning.PrintOnError(os.MkdirAll(installDir, 0755))

		// Downloading asset.
		pterm.Fatal.PrintOnError(internal.DownloadFile(installPath, asset.DownloadURL))
		pterm.Success.Printf("Downloaded %s\n", asset.Name)

		// Installing asset.
		pterm.Fatal.PrintOnError(archiver.Unarchive(installPath, installDir))
		pterm.Warning.PrintOnError(os.Remove(installPath))
		internal.AddToPath(installDir, internal.Repo.Name)

		// Success message.
		pterm.Success.Printfln("%s was installed successfully!\nYou might have to restart your terminal session to use %s.", internal.Repo.Name, internal.Repo.Name)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empry strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "d", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "only outputs errors")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
	pterm.Error = *pterm.Error.WithShowLineNumber(false)
}
