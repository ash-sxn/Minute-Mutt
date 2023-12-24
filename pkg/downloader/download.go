package downloader

import (
	"fmt"
	"os/exec"
	"time"
)

// DownloadVideo uses yt-dlp to download a video by its ID and saves it to the specified output directory.
// maxResolution is the maximum allowed resolution (e.g., "1080p").
// startTimeStr and endTimeStr are strings representing the start and end times of the download window in "HH:MM" format.
func DownloadVideo(videoID, outputDir, maxResolution, startTimeStr, endTimeStr string) error {
	// Get the current time
	currentTime := time.Now()

	// Parse the start and end times for today's date
	startTime, err := time.Parse("15:04", startTimeStr)
	if err != nil {
		return fmt.Errorf("failed to parse start time: %v", err)
	}
	endTime, err := time.Parse("15:04", endTimeStr)
	if err != nil {
		return fmt.Errorf("failed to parse end time: %v", err)
	}

	// Adjust the start and end times to today's date
	startTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), startTime.Hour(), startTime.Minute(), 0, 0, currentTime.Location())
	endTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), endTime.Hour(), endTime.Minute(), 0, 0, currentTime.Location())

	// If the end time is before the start time, it means the end time is on the next day
	if endTime.Before(startTime) {
		endTime = endTime.Add(24 * time.Hour)
	}

	// Check if the current time is within the download window
	if currentTime.Before(startTime) || currentTime.After(endTime) {
		return fmt.Errorf("current time is outside the download window")
	}

	// Construct the yt-dlp command
	cmd := exec.Command(
		"yt-dlp",
		"-S", fmt.Sprintf("res:%s", maxResolution),
		"-P", outputDir,
		videoID,
	)

	// Run the yt-dlp command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error downloading video: %w\nOutput: %s", err, string(output))
	}

	fmt.Printf("Video downloaded: %s\n", videoID)
	return nil
}
