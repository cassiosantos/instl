package cmd

import (
	"errors"
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

    iwr -useb instl.sh/username/reponame/windows | iex  
  
**macOS**  

    curl -fsSL instl.sh/username/reponame/macos | bash   
  
**Linux**  

    curl -fsSL instl.sh/username/reponame/linux | bash  
  
(Replace username and reponame with the GitHub project you want to install)  

Read more about the web installer here: https://docs.instl.sh/#/web-installer
  
These commands can be executed from any system and install the respective GitHub project.  
You can also provide these commands to your users to make your GitHub project easily installable.`,
	Version: "v0.0.14", // <---VERSION---> This comment enables auto-releases on version change!
	Example: "instl instl-sh/instl",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a GitHub repo to install\nExample: instl user/repo")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Get GitHub username and repository from args.
		repoArg := args[0]
		repoArgParts := strings.Split(repoArg, "/")
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
		var repo internal.Repository
		internal.MakeSpinner("Getting asset metadata from latest asset...", func() string {
			repo = internal.ParseRepository(repoArg)
			var assetCount int
			repo.ForEachAsset(func(release internal.Asset) {
				assetCount++
			})

			return pterm.Sprintf("Found %d assets in latest asset!", assetCount)
		})

		// Detect right asset for system.
		var asset internal.Asset
		internal.MakeSpinner("Detecting right asset for machine...", func() string {
			pterm.Debug.Println("Your system:", runtime.GOOS, runtime.GOARCH)
			asset = internal.DetectRightAsset(repo)
			return pterm.Sprintf("Found an asset which seems to fit to your system:")
		})

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
		installPath := internal.GetInstallPath(repo.User, repo.Name) + "/" + asset.Name
		installDir := internal.GetInstallPath(repo.User, repo.Name)
		pterm.Debug.PrintOnError(os.RemoveAll(installDir))
		pterm.Warning.PrintOnError(os.MkdirAll(installDir, 0755))

		// Downloading asset.
		pterm.Fatal.PrintOnError(internal.DownloadFile(installPath, asset.DownloadURL))
		pterm.Success.Printf("Downloaded %s\n", asset.Name)

		// Installing asset.
		pterm.Fatal.PrintOnError(archiver.Unarchive(installPath, installDir))
		pterm.Warning.PrintOnError(os.Remove(installPath))
		internal.AddToPath(installDir, repo.Name)

		// Success message.
		pterm.Success.Printfln("%s was installed successfully!\nYou might have to restart your terminal session to use %s", repo.Name, repo.Name)
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

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
