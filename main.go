package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vodm URL1 [URL2...]")
		fmt.Println("Example: vodm https://example.com/video.mp4")
		os.Exit(1)
	}

	urls := os.Args[1:]

	for _, url := range urls {
		fmt.Printf("Processing link: %s\n", url)

		filename := filepath.Base(url)
		err := downloadFile(filename, url)
		if err != nil {
			fmt.Printf("❌ Error downloading %s: %v\n", url, err)
			continue
		}

		fmt.Printf("✅ Successfully downloaded: %s\n\n", filename)
	}
}

func downloadFile(filename string, url string) error {
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer out.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		progressbar.OptionSetDescription(filename),
		progressbar.OptionSetWidth(15),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionOnCompletion(func() {
			fmt.Println()
		}),
		progressbar.OptionShowCount(),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}
