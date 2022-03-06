package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
)

// MakeSpinner creates a spinner.
func MakeSpinner(msg string, f func() (string, error)) error {
	spinner, _ := pterm.DefaultSpinner.Start(msg)
	resolvedMsg, err := f()
	if err != nil {
		spinner.RemoveWhenDone = true
		spinner.Stop()
		return err
	}

	spinner.RemoveWhenDone = msg == ""

	spinner.Success(resolvedMsg)

	return nil
}

// FindBinary tries to find a binary in the release asset.
func FindBinary(path string) (binaryPath, binaryName string, err error) {
	var dotExeFiles []string
	var noExtFiles []string
	var detectedFiles []string

	blacklist := []string{"LICENSE", "LICENCE"}

	err = filepath.Walk(path, func(currentPath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(currentPath)
		file := filepath.Base(currentPath)

		for _, s := range blacklist {
			if strings.Contains(file, s) {
				return nil
			}
		}

		detectedFiles = append(detectedFiles, currentPath)

		if ext == ".exe" {
			dotExeFiles = append(dotExeFiles, currentPath)
		}

		if !strings.Contains(file, ".") {
			noExtFiles = append(noExtFiles, currentPath)
		}

		return nil
	})
	if err != nil {
		return "", "", err
	}

	if runtime.GOOS == "windows" {
		if len(dotExeFiles) == 1 {
			return filepath.Dir(dotExeFiles[0]), filepath.Base(dotExeFiles[0]), nil
		}
	} else {
		pterm.Debug.Printfln("NoExtFiles: %#v", noExtFiles)
		if len(noExtFiles) == 1 {
			return filepath.Dir(noExtFiles[0]), filepath.Base(noExtFiles[0]), nil
		}
	}

	for i, file := range detectedFiles {
		detectedFiles[i] = filepath.Clean(strings.Join(strings.Split(file, "instl")[1:], ""))
	}

	pterm.Error.Println("We could not find a binary file inside the release asset.")
	CreateIssue(fmt.Sprintf("[Binary Detection] `%s/%s`", Repo.User, Repo.Name), fmt.Sprintf("Detected files: \n```go\n%#v\n```\n", detectedFiles))
	return "", "", errors.New("could not find binary file in asset")
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
