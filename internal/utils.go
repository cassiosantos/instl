package internal

import (
	"io/fs"
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
	spinner.Success(resolvedMsg)

	return nil
}

func FindBinary(path string) (binaryPath, binaryName string, err error) {
	var dotExeFiles []string
	var noExtFiles []string

	blacklist := []string{"LICENSE", "LICENCE"}

	filepath.Walk(path, func(currentPath string, info fs.FileInfo, err error) error {
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

		if ext == ".exe" {
			dotExeFiles = append(dotExeFiles, currentPath)
		}

		if !strings.Contains(file, ".") {
			noExtFiles = append(noExtFiles, currentPath)
		}

		return nil
	})

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

	return
}
