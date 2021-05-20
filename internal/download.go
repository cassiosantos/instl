package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pterm/pterm"
)

// WriteCounter counts the number of bytes written to it.
type WriteCounter struct {
	Total uint64
	pb    *pterm.ProgressbarPrinter
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.pb.Add(len(p))
	return n, nil
}

// DownloadFile downloads a file to tmp and displays a progressbar.
// The file will be moved to output directory when downloading is finished.
func DownloadFile(output, url string) error {
	path := filepath.Clean(output)
	pterm.Debug.Printf("Downloading to %s\n", path)
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create download path: %w", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return fmt.Errorf("error while downloading file: %w", err)
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	fileSize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	counter.pb, _ = pterm.DefaultProgressbar.WithRemoveWhenDone().WithTitle("Downloading asset").WithTotal(fileSize).Start()
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	pterm.Debug.PrintOnError(os.Chmod(path, 0755))

	out.Close()
	return nil
}
