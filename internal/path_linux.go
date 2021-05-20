package internal

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

// GetInstallPath returns the path, where instl will install the project to.
func GetInstallPath(username, programName string) string {
	basePath := pterm.Sprintf("/usr/local/lib/instl/%s/%s", username, programName)
	basePath = filepath.Clean(basePath)
	os.MkdirAll(basePath, 0755)

	return basePath
}

// AddToPath adds a value to the global system path environment variable.
func AddToPath(path, filename string) {
	path, binaryName, err := FindBinary(path)
	pterm.Error.PrintOnError(err)
	if err != nil {
		os.Exit(1)
	}

	pterm.Debug.Printfln("Path: %s, Binary: %s", path, binaryName)

	err = os.Symlink(path+"/"+binaryName, "/usr/local/bin/"+binaryName)
	if err != nil {
		pterm.Debug.Println("Symlink already exists. This is not a problem, the old one will work too.")
	}
}
