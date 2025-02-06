package downloader

import (
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
)

func GetFilename(url string) string {
	ext := filepath.Ext(url)
	if ext == "" {
		ext = ".mp4"
	}
	return fmt.Sprintf("download_%s%s", uuid.New(), ext)
}
