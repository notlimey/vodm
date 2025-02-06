package main

import (
	"fmt"
	"github.com/notlimey/vodm/internal/downloader"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vodm URL1 [URL2...]")
		fmt.Println("Example: vodm https://example.com/video.mp4")
		fmt.Println("Example: vodm ./videos.txt")
		os.Exit(1)
	}

	firstArg := os.Args[1]

	// Check if it's a file and has .txt extension
	if strings.HasSuffix(firstArg, ".txt") && isFile(firstArg) {
		fmt.Printf("Processing text file: %s\n", firstArg)
		urls, err := readURLsFromFile(firstArg)
		if err != nil {
			fmt.Printf("Error reading URLs from file %s: %s\n", firstArg, err)
			os.Exit(1)
		}

		fmt.Printf("Found %d URLs in file\n", len(urls))

		for _, url := range urls {
			fmt.Printf("Processing URL: %s\n", url)
			filename := filepath.Base(url)
			err := downloader.DownloadFile(filename, url)
			if err != nil {
				fmt.Printf("❌ Error downloading %s: %v\n", url, err)
				continue
			}
			fmt.Printf("✅ Successfully downloaded: %s\n\n", url)
		}
	} else {
		// Process arguments as URLs
		urls := os.Args[1:]
		for _, url := range urls {
			fmt.Printf("Processing URL: %s\n", url)
			filename := filepath.Base(url)
			err := downloader.DownloadFile(filename, url)
			if err != nil {
				fmt.Printf("❌ Error downloading %s: %v\n", url, err)
				continue
			}
			fmt.Printf("✅ Successfully downloaded: %s\n\n", url)
		}
	}
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func readURLsFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			urls = append(urls, line)
		}
	}

	return urls, nil
}
