package downloader

import (
	"fmt"
	"os/exec"
)

// DownloadVideo uses yt-dlp to download a video by its ID and saves it to the specified output directory.
// maxResolution is the maximum allowed resolution (e.g., "1080p").
// startTimeStr and endTimeStr are strings representing the start and end times of the download window in "HH:MM" format.
func DownloadVideo(videoID, outputDir, maxResolution string) error {
	// Construct the yt-dlp command
	cmd := exec.Command(
		"yt-dlp",
		"-S", fmt.Sprintf("res:%s", maxResolution),
		"--merge-output-format", "mp4", // This line ensures the output is in MP4 format
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
