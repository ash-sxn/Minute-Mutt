package downloader

import (
	"fmt"
	"os"
	"time"

	youtube "github.com/kkdai/youtube/v2"
)

// DownloadVideo downloads a video by its ID and saves it to the specified output directory.
// maxResolution is the maximum allowed resolution (e.g., "1080p").
// startTime and endTime define the time window during which downloading is allowed.
func DownloadVideo(videoID string, outputDir string, maxResolution string, startTime, endTime time.Time) error {
	currentTime := time.Now()

	// Check if current time is within the download window
	if currentTime.Before(startTime) || currentTime.After(endTime) {
		return fmt.Errorf("current time is outside the download window")
	}

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		return fmt.Errorf("error fetching video info: %w", err)
	}

	// Find the best format that doesn't exceed the max resolution
	format := video.Formats.
		WithAudioChannels().
		Filter(youtube.FormatResolutionKey, func(resolution string) bool {
			return resolution <= maxResolution
		}).
		FindBest(youtube.FormatResolutionKey)[0]

	outputFile, err := os.Create(fmt.Sprintf("%s/%s.mp4", outputDir, video.Title))
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer outputFile.Close()

	err = client.Download(video, &format, outputFile)
	if err != nil {
		return fmt.Errorf("error downloading video: %w", err)
	}

	fmt.Printf("Video downloaded: %s\n", video.Title)
	return nil
}
