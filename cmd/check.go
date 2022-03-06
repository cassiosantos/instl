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
		pathString := os.Getenv("PATH")
		pterm.Debug.Println(pathString)
		pathParts := strings.Split(pathString, ":")
		homeDir, _ := os.UserHomeDir()
		pathParts = append(pathParts, homeDir+"/bin")
		pathParts = append(pathParts, "~/bin")

		pterm.DefaultSection.Println("Permissions")
		pterm.Info.Printfln("Has administrative rights: %v\n", isadmin.Check())
		var writablePathLocations []string
		for _, path := range pathParts {
			hasPerm := internal.CheckPermissionsToDir(path)
			style := pterm.NewStyle(pterm.FgRed)
			if hasPerm {
				style = pterm.NewStyle(pterm.FgGreen)
				writablePathLocations = append(writablePathLocations, path)
			}
			style.Printfln("Can write to %s: %v", path, hasPerm)
		}

		pterm.DefaultSection.Println("Possible Installation Locations")
		locations := []string{homeDir + "/bin", homeDir + "/.local/bin", "/usr/local/bin", "/usr/bin", "/bin"}

		if len(writablePathLocations) > 0 {
			// Check if pathParts contains any of the possible installation locations
			for _, path := range writablePathLocations {
				for _, loc := range locations {
					if strings.Contains(path, loc) {
						pterm.Info.Printfln("Found possible installation location: %s", path)
					}
				}
			}
		} else {
			pterm.Warning.Println("No possible installation locations found")
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
