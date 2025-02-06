package file

import (
	"fmt"
	"os"
	"strings"
)

func GetUrlsFromFile(filename string) []string {
	fmt.Printf("Processing text file: %s\n", filename)
	urls, err := readURLsFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading URLs from file %s: %s\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("Found %d URLs in file\n", len(urls))

	return urls
}

func ArgumentIsFile(arg string) bool {
	return strings.HasSuffix(arg, ".txt") && IsFile(arg)
}

func IsFile(path string) bool {
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
