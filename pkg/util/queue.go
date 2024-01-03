package util

import (
	"encoding/csv"
	"os"

	"github.com/ash-sxn/Minute-Mutt/pkg/youtube"
)

const downloadedVideosFile = "pkg/database/history_queue.csv"

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
