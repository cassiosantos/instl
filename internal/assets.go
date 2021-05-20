package internal

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

var blacklistRegexes = []string{
	`\.sha256$`,
	`\.sum$`,
	`\.md5$`,
	`\.txt$`,
}

var validGoarches = []string{
	"386",
	"amd64",
	"amd64p32",
	"arm",
	"armbe",
	"arm64",
	"arm64be",
	"ppc64",
	"ppc64le",
	"mips",
	"mipsle",
	"mips64",
	"mips64le",
	"mips64p32",
	"mips64p32le",
	"ppc",
	"riscv",
	"riscv64",
	"s390",
	"s390x",
	"sparc",
	"sparc64",
	"wasm",
}

var validGooses = []string{
	"aix",
	"android",
	"darwin",
	"dragonfly",
	"freebsd",
	"hurd",
	"illumos",
	"js",
	"linux",
	"nacl",
	"netbsd",
	"openbsd",
	"plan9",
	"solaris",
	"windows",
	"zos",
}

// DetectRightAsset tries to detect the right asset for the current machine.
func DetectRightAsset(repo Repository) Asset {
	goos := runtime.GOOS

	var (
		windowsAssets []Asset
		linuxAssets   []Asset
		darwinAssets  []Asset
	)

	windowsRegex := generateMultiRegex("windows", `\.exe$`)
	linuxRegex := generateMultiRegex("linux")
	darwinRegex := generateMultiRegex("darwin")

	repo.ForEachAsset(func(release Asset) {
		name := release.Name

		for _, v := range validGoarches {
			if v != runtime.GOARCH {
				blacklistRegexes = append(blacklistRegexes, v)
			}
		}

		if generateMultiRegex(blacklistRegexes...).MatchString(name) {
			return
		}

		switch {
		case goos == "windows" && windowsRegex.MatchString(name):
			windowsAssets = append(windowsAssets, release)
		case goos == "linux" && linuxRegex.MatchString(name):
			linuxAssets = append(linuxAssets, release)
		case goos == "darwin" && darwinRegex.MatchString(name):
			darwinAssets = append(darwinAssets, release)
		}
	})

	switch goos {
	case "windows":
		return findBestRelease(analyzeMultiReleases(&windowsAssets))
	case "linux":
		return findBestRelease(analyzeMultiReleases(&linuxAssets))
	case "darwin":
		return findBestRelease(analyzeMultiReleases(&darwinAssets))
	}

	return Asset{}
}

func generateMultiRegex(parts ...string) *regexp.Regexp {
	return regexp.MustCompile(`(?mi)` + strings.Join(parts, "|"))
}

func findBestRelease(m map[*Asset]int) (r Asset) {
	var highest int
	for release, i := range m {
		if i > highest {
			r = *release
			r.Score = i
		}
	}

	return
}

func analyzeMultiReleases(releases *[]Asset) map[*Asset]int {
	counted := map[*Asset]int{}
	for _, release := range *releases {
		re := generateGoarchRegex()
		counted[&release] = len(re.FindAllString(release.Name, -1))
	}
	return counted
}

func generateGoarchRegex(parts ...string) *regexp.Regexp {
	parts = append(parts, runtime.GOARCH)

	if runtime.GOARCH == "amd64" {
		parts = append(parts, "64")
		parts = append(parts, "64bit")
		parts = append(parts, "x64")
	}

	if runtime.GOARCH == "386" {
		parts = append(parts, "32")
		parts = append(parts, "32bit")
		parts = append(parts, "86")
		parts = append(parts, "x86")
	}

	return generateMultiRegex(parts...)
}

// ReadbleSize returns a human readble size.
func ReadbleSize(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
