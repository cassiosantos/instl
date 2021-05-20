package internal

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
)

func CreateIssue(title, body string) {
	baseURL := "https://github.com/instl-sh/instl/issues/new?title=%s&body=%s"
	baseURL = fmt.Sprintf(baseURL, escapeURL(title), escapeURL(body))
	pterm.Debug.Printfln("Issue URL: %s", baseURL)
	var open bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to report this to the instl team?",
		Default: true,
	}, &open, survey.WithValidator(survey.Required))

	if open {
		openbrowser(baseURL)
	}
}

func escapeURL(text string) string {
	return url.PathEscape(text)
}

func openbrowser(url string) {
	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows":
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		_ = exec.Command("open", url).Start()
	}
}
