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
