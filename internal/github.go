package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/tidwall/gjson"
)

// Repository contains infomation about a GitHub repository.
type Repository struct {
	User        string
	Name        string
	URL         string
	ReleasesURL string
	Releases    gjson.Result
}

// Asset is a representation of an asset in a GitHub release.
type Asset struct {
	Name          string
	Size          int64
	DownloadCount int64
	UpdatedAt     time.Time
	DownloadURL   string
	Version       string
	Score         int
}

// ParseRepository parses a repository from a string.
func ParseRepository(repo string) Repository {
	// Parse "https://github.com/name/repo", "github.com/name/repoStr", etc. to "name/repoStr"
	repoParts := strings.Split(repo, "/")
	repo = repoParts[len(repoParts)-2] + "/" + repoParts[len(repoParts)-1]

	r := Repository{
		User:        repoParts[len(repoParts)-2],
		Name:        repoParts[len(repoParts)-1],
		URL:         "https://github.com/" + repo,
		ReleasesURL: pterm.Sprintf("https://api.github.com/repos/%s/releases/latest", repo),
	}

	resp, err := http.Get(r.ReleasesURL)
	if err != nil {
		pterm.Fatal.Println(fmt.Errorf("could not get github releases json: %w", err))
	}
	defer resp.Body.Close()

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		pterm.Fatal.Println(fmt.Errorf("could not get github releases json: %w", err))
	}
	json := string(jsonBytes)

	r.Releases = gjson.Get(json, "assets.#.{name,size,download_count,updated_at,browser_download_url}")

	return r
}

// ForEachAsset iterates over every asset in a release.
func (repo Repository) ForEachAsset(f func(release Asset)) {
	repo.Releases.ForEach(func(key, value gjson.Result) bool {
		release := Asset{
			Name:          value.Get("name").String(),
			Size:          value.Get("size").Int(),
			DownloadCount: value.Get("download_count").Int(),
			UpdatedAt:     value.Get("updated_at").Time(),
			DownloadURL:   value.Get("browser_download_url").String(),
			Version:       strings.Split(value.Get("browser_download_url").String(), "/")[7],
		}

		f(release)

		return true
	})
}
