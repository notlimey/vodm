package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Check if URLs were provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: vodm URL1 [URL2...]")
		fmt.Println("Example: vodm https://example.com/file.mp4")
		os.Exit(1)
	}

	// Get URLs from command line arguments (skip program name at os.Args[0])
	urls := os.Args[1:]

	// Process each URL
	for _, url := range urls {
		fmt.Printf("Downloading: %s\n", url)

		// Get filename from URL
		filename := filepath.Base(url)

		// Download the file
		err := downloadFile(filename, url)
		if err != nil {
			fmt.Printf("Error downloading %s: %v\n", url, err)
			continue
		}

		fmt.Printf("Successfully downloaded: %s\n", filename)
	}
}

func downloadFile(filename string, url string) error {
	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not download file: %v", err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}
