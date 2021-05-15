package internal

import "github.com/pterm/pterm"

// MakeSpinner creates a spinner.
func MakeSpinner(msg string, f func() string) {
	spinner, _ := pterm.DefaultSpinner.Start(msg)
	resolvedMsg := f()
	spinner.Success(resolvedMsg)
}
