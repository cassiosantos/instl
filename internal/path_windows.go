package internal

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"golang.org/x/sys/windows/registry"
)

func GetInstallPath(username, programName string) string {
	basePath, _ := os.UserHomeDir()
	basePath += pterm.Sprintf(`/instl/%s/%s/`, username, programName)
	basePath = filepath.Clean(basePath)
	os.MkdirAll(basePath, 0755)

	return basePath
}

// AddToPath adds a value to the global system path environment variable.
func AddToPath(path, filename string) {
	pterm.Debug.Printfln("Adding %s to path", path)

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		panic(err)
	}
	defer k.Close()

	oldPath, _, _ := k.GetStringValue("Path")

	if strings.Contains(oldPath, path) {
		pterm.Debug.Printfln("Path %s is already in the system path", path)

		return
	}

	err = k.SetStringValue("Path", oldPath+";"+path)
	if err != nil {
		pterm.Fatal.Println(err)
	}
	pterm.Debug.Printfln("Added to path")
}
