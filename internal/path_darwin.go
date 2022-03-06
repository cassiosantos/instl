package internal

import (
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
)

// GetInstallPath returns the path, where instl will install the project to.
func GetInstallPath(programName string) string {
	homeDir, _ := os.UserHomeDir()
	basePath := pterm.Sprintf(homeDir+"/.local/bin/_instl/%s", programName)
	basePath = filepath.Clean(basePath)
	os.MkdirAll(basePath, 0755)

	return basePath
}

// AddToPath adds a value to the global system path environment variable.
func AddToPath(path, filename string) {
	homeDir, _ := os.UserHomeDir()
	path, binaryName, err := FindBinary(path)
	pterm.Error.WithShowLineNumber(false).PrintOnError(err)
	if err != nil {
		os.Exit(1)
	}

	pterm.Debug.Printfln("Path: %s, Binary: %s", path, binaryName)
	pterm.Fatal.PrintOnError(AppendPathToShellProfileFiles())

	err = os.Symlink(path+"/"+binaryName, homeDir+"/.local/bin/"+Repo.Name)
	if err != nil {
		pterm.Debug.Println("Symlink already exists. This is not a problem, the old one will work too.")
	}
}
