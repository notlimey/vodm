package main

import (
	"fmt"
	"github.com/notlimey/vodm/internal/arguments"
	"github.com/notlimey/vodm/internal/downloader"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vodm URL1 [URL2...]")
		fmt.Println("Example: vodm https://example.com/video.mp4")
		fmt.Println("Example: vodm ./videos.txt")
		os.Exit(1)
	}

	args := arguments.ParseArguments(os.Args[1:])

	if args.Flags.Concurrent {
		fmt.Println("Using concurrent downloads")
	}

	for _, url := range args.Urls {
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
