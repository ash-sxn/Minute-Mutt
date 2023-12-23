package util

import (
	"encoding/csv"
	"os"

	"github.com/ash-sxn/Minute-Mutt/pkg/youtube"
)

const downloadedVideosFile = "pkg/database/downloaded_videos.csv"

// VideoQueue represents a queue of videos to be downloaded.
type VideoQueue struct {
	Videos []youtube.Video
}

// NewVideoQueue creates a new VideoQueue.
func NewVideoQueue() *VideoQueue {
	return &VideoQueue{}
}

// AddVideo adds a new video to the queue if it hasn't been downloaded.
func (vq *VideoQueue) AddVideo(video youtube.Video) {
	if !vq.IsDownloaded(video.ID) {
		vq.Videos = append(vq.Videos, video)
	}
}

// RemoveVideo removes a video from the queue after it's been downloaded.
func (vq *VideoQueue) RemoveVideo(videoID string) {
	for i, video := range vq.Videos {
		if video.ID == videoID {
			vq.Videos = append(vq.Videos[:i], vq.Videos[i+1:]...)
			break
		}
	}
}

// IsDownloaded checks if a video has been downloaded by looking it up in the downloaded_videos.csv file.
func (vq *VideoQueue) IsDownloaded(videoID string) bool {
	file, err := os.Open(downloadedVideosFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := csv.NewReader(file)
	for {
		record, err := scanner.Read()
		if err != nil {
			break
		}
		if record[0] == videoID {
			return true
		}
	}

	return false
}

// MarkAsDownloaded adds a video to the downloaded_videos.csv file.
func (vq *VideoQueue) MarkAsDownloaded(video youtube.Video) error {
	file, err := os.OpenFile(downloadedVideosFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{video.ID, video.Title}); err != nil {
		return err
	}

	return nil
}

// SaveToCSV writes the video queue to a CSV file.
func (vq *VideoQueue) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	if err := writer.Write([]string{"ID", "Title"}); err != nil {
		return err
	}

	// Write the video data
	for _, video := range vq.Videos {
		if err := writer.Write([]string{video.ID, video.Title}); err != nil {
			return err
		}
	}

	return nil
}
