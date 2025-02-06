package downloader

import (
	"fmt"
)

type DownloadJob struct {
	URL      string
	Filename string
}

type DownloadResult struct {
	Job     DownloadJob
	Error   error
	Success bool
}

func DownloadWithWorkers(urls []string, numWorkers int) {
	// Create channels for jobs and results
	jobs := make(chan DownloadJob, len(urls))
	results := make(chan DownloadResult, len(urls))

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Queue jobs
	for _, url := range urls {
		filename := GetFilename(url)

		jobs <- DownloadJob{
			URL:      url,
			Filename: filename,
		}
	}
	close(jobs)

	// Collect results
	for i := 0; i < len(urls); i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("❌ Error downloading %s: %v\n", result.Job.URL, result.Error)
		} else {
			fmt.Printf("✅ Successfully downloaded: %s\n", result.Job.Filename)
		}
	}
}

// Worker function
func worker(id int, jobs <-chan DownloadJob, results chan<- DownloadResult) {
	for job := range jobs {
		fmt.Printf("Worker %d starting download of %s\n", id, job.URL)

		err := DownloadFile(job.Filename, job.URL)

		results <- DownloadResult{
			Job:     job,
			Error:   err,
			Success: err == nil,
		}
	}
}
