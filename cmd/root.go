package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/mholt/archiver/v3"
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/instl-sh/instl/internal"
)

var rootCmd = &cobra.Command{
	Use:   "instl [repo]",
	Short: "Instl can install GitHub projects on your machine. No setup required.",
	Long: `Instl is a CLI tool which detects the right release of a GitHub repository for your system.
It will download the detected release and install the asset files to your computer.
The repositories themself, don't need a setup to be installable with instl. They just need a default release with assets for multiple operating systems.`,
	Example: `instl instl-sh/instl
instl yourName/yourRepo`,
	Version: "v0.0.2", // <---VERSION---> This comment enables auto-releases on version change!
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a GitHub repo to install")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		repoArg := args[0]
		repoArgParts := strings.Split(repoArg, "/")
		repoName := repoArgParts[len(repoArgParts)-2] + "/" + repoArgParts[len(repoArgParts)-1]

		introText, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithStyle("  INSTL", pterm.NewStyle(pterm.FgMagenta))).Srender()
		pterm.Println()
		pterm.Println(strings.TrimRight(introText, "\n"))
		pterm.Printf(pterm.Cyan("                > https://instl.sh\n\n"))
		pterm.Info.Printf("instl.sh is an automated installer for GitHub projects.\nWe do not own https://github.com/%s.\n", pterm.Magenta(repoName))
		pterm.Println()
		pterm.DefaultHeader.Printf("Running installer for github.com/%s", repoName)
		pterm.Println()

		var repo internal.Repository
		internal.MakeSpinner("Getting asset metadata from latest release...", func() string {
			repo = internal.ParseRepository(repoArg)
			var releasesCount int
			repo.ForEachRelease(func(release internal.Release) {
				releasesCount++
			})

			return pterm.Sprintf("Found %d assets in latest release!", releasesCount)
		})

		var release internal.Release
		internal.MakeSpinner("Detecting right asset for machine...", func() string {
			pterm.Debug.Println("Your system:", runtime.GOOS, runtime.GOARCH)
			release = internal.DetectRightRelease(repo)
			return pterm.Sprintf("Found an asset which seems to fit to your system:")
		})
		assetStats, _ := pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
			{"Name", "Last Update", "Download Count"},
			{release.Name, release.UpdatedAt.Format(time.RFC1123), pterm.Sprint(release.DownloadCount)},
		}).Srender()
		pterm.DefaultBox.Println(assetStats)
		installPath := internal.GetInstallPath(repo.User, repo.Name) + "/" + release.Name
		installDir := internal.GetInstallPath(repo.User, repo.Name)
		os.RemoveAll(installDir)
		os.MkdirAll(installDir, 0755)
		err := internal.DownloadFile(installPath, release.DownloadURL)
		if err != nil {
			pterm.Fatal.Println(err)
		}
		pterm.Success.Printf("Downloaded %s\n", release.Name)

		err = archiver.Unarchive(installPath, installDir)
		if err != nil {
			pterm.Fatal.Println(err)
		}

		os.Remove(installPath)
		internal.AddToPath(installDir, repo.Name)

		pterm.Success.Printfln("%s was installed successfully!\nYou might have to restart your terminal session to use %s", repo.Name, repo.Name)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empry strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRootCmd(rootCmd)
	rootCmd.AddCommand(pcli.GetCiCommand())
	rootCmd.SetFlagErrorFunc(pcli.FlagErrorFunc())
	rootCmd.SetGlobalNormalizationFunc(pcli.GlobalNormalizationFunc())
	rootCmd.SetHelpFunc(pcli.HelpFunc())
	rootCmd.SetHelpTemplate(pcli.HelpTemplate())
	rootCmd.SetUsageFunc(pcli.UsageFunc())
	rootCmd.SetUsageTemplate(pcli.UsageTemplate())
	rootCmd.SetVersionTemplate(pcli.VersionTemplate())

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
