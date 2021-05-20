package internal

import "github.com/pterm/pterm"

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
