package downloader

import (
	"fmt"
)

type DownloadManager struct {
	NumWorkers int
	Jobs       chan DownloadJob
	Results    chan DownloadResult
	Done       chan bool
}

func NewDownloadManager(numWorkers int) *DownloadManager {
	return &DownloadManager{
		NumWorkers: numWorkers,
		Jobs:       make(chan DownloadJob, 100),
		Results:    make(chan DownloadResult, 100),
		Done:       make(chan bool),
	}
}

func (dm *DownloadManager) Start() {
	// Start workers
	for w := 1; w <= dm.NumWorkers; w++ {
		go worker(w, dm.Jobs, dm.Results)
	}

	// Start result collector
	go dm.collectResults()
}

func (dm *DownloadManager) AddJob(url string) {
	filename := GetFilename(url)
	dm.Jobs <- DownloadJob{
		URL:      url,
		Filename: filename,
	}
}

func (dm *DownloadManager) collectResults() {
	var completed int
	total := len(dm.Jobs)

	for result := range dm.Results {
		completed++
		if result.Error != nil {
			fmt.Printf("❌ Error downloading %s: %v\n", result.Job.URL, result.Error)
		} else {
			fmt.Printf("✅ Successfully downloaded: %s\n", result.Job.Filename)
		}

		if completed == total {
			dm.Done <- true
			return
		}
	}
}
