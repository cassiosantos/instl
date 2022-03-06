package internal

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"strings"
)

func FindShellProfileFiles() []string {
	homeDir, _ := os.UserHomeDir()
	shellProfileFileNames := []string{
		homeDir + "/" + ".bashrc",
		homeDir + "/" + ".zshrc",
		homeDir + "/" + ".bash_profile",
		homeDir + "/" + ".profile",
	}

	var shellProfileFiles []string

	// check if the files exist in the user home dir
	for _, shellProfileFileName := range shellProfileFileNames {
		if FileExists(shellProfileFileName) {
			pterm.Debug.Println("Found shell profile file: " + homeDir + "/" + shellProfileFileName)
			shellProfileFiles = append(shellProfileFiles, shellProfileFileName)
		}
	}

	pterm.Debug.Println("shell profile files:", shellProfileFiles)

	return shellProfileFiles
}

func AppendPathToShellProfileFiles() error {
	shellProfileFiles := FindShellProfileFiles()
	homeDir, _ := os.UserHomeDir()

	for _, shellProfileFile := range shellProfileFiles {
		pterm.Debug.Println("Checking shell profile file:", shellProfileFile)
		dat, err := os.ReadFile(shellProfileFile)
		if err != nil {
			return err
		}

		if strings.Contains(string(dat), fmt.Sprintf(`export PATH="$PATH:%s`, homeDir+"/.local/bin")) {
			pterm.Debug.Println("path already exists in", shellProfileFile)
			continue
		}

		// Append to env path
		pterm.Debug.Println("Appending path to", shellProfileFile)
		f, err := os.OpenFile(shellProfileFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		_, err = f.WriteString(fmt.Sprintf("\n"+`export PATH="$PATH:%s"`+"\n", homeDir+"/.local/bin"))
		if err != nil {
			return err
		}
		f.Close()
	}

	return nil
}
