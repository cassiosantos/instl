package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "instl",
	Short: "Instl can install GitHub projects on your machine. No setup required.",
	Long: `Instl is a CLI tool which detects the right release of a GitHub repository for your system.
It will download the detected release and install the asset files to your computer.
The repositories themself, don't need a setup to be installable with instl. They just need a default release with assets for multiple operating systems.`,
	Version: "v0.0.1", // <---VERSION---> This comment enables auto-releases on version change!
	// 	Run: func(cmd *cobra.Command, args []string) {   },
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
