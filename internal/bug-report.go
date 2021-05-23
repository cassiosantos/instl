package internal

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"github.com/pterm/pterm"
)

func CreateIssue(title, body string) {
	baseURL := "https://github.com/instl-sh/instl/issues/new?title=%s&body=%s"
	baseURL = fmt.Sprintf(baseURL, escapeURL(title), escapeURL(body))
	pterm.Debug.Printfln("Issue URL: %s", baseURL)

	pterm.Info.Printf("Do you want to report this to the instl team? %s ", pterm.Gray("[Y/n]"))

	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()

	if char != 'n' {
		openbrowser(baseURL)
	}
	pterm.Println()
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
