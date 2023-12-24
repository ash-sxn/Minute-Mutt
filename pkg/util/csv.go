package util

import (
	"encoding/csv"
	"os"

	myYoutube "github.com/ash-sxn/Minute-Mutt/pkg/youtube"
)

// WriteChannelsToCSV writes the slice of Channels to a CSV file
func WriteChannelsToCSV(channels []myYoutube.Channel, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {

		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Channel Name", "Channel ID"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, channel := range channels {
		record := []string{channel.Name, channel.ID}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// AddVideoToCSV appends a single video to the end of a CSV file.
func AddVideoToCSV(video myYoutube.Video, filename string) error {
	// Open the file with the O_APPEND and O_WRONLY flags to append data to the file and write only.
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the video data
	if err := writer.Write([]string{video.ID, video.Title}); err != nil {
		return err
	}

	return nil
}

// LoadDownloadedVideos loads the history of downloaded videos from a CSV file.
func LoadDownloadedVideos(filename string) (map[string]struct{}, error) {
	downloaded := make(map[string]struct{})
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if len(record) > 0 {
			videoID := record[0]
			downloaded[videoID] = struct{}{}
		}
	}

	return downloaded, nil
}
