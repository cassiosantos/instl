package cmd

import (
	"github.com/atomicgo/isadmin"
	"github.com/installer/instl/internal"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Performs various tests to ensure instl is working correctly on the host machine",
	Long: `The check command performs various tests to ensure instl is working correctly on the host machine.
Use the -d (debug) flag to see more detailed output.`,
	Run: func(cmd *cobra.Command, args []string) {
		pterm.DefaultSection.Println("PATH Environment Variable")
		pathString := os.Getenv("PATH")
		pterm.Debug.Println(pathString)
		pathParts := strings.Split(pathString, ":")
		pterm.NewBulletListFromStrings(pathParts, "").Render()

		pterm.DefaultSection.Println("Permissions")
		pterm.Printfln("Has administrative rights: %v", isadmin.Check())
		for _, path := range pathParts {
			hasPerm := internal.CheckPermissionsToDir(path)
			pterm.Printfln("Can write to %s: %v", path, hasPerm)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
