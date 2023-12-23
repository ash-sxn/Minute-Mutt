package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ash-sxn/Minute-Mutt/pkg/auth"
	"google.golang.org/api/youtube/v3"
)

// Channel holds the name and ID of a YouTube channel
type Channel struct {
	Name string
	ID   string
}

func main() {
	const youtubeScope = youtube.YoutubeReadonlyScope
	client := auth.GetClient(youtubeScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	var allSubscribedChannels []Channel // Slice to store all subscribed channels

	// Function to recursively fetch all pages of subscribed channels
	var fetchSubscriptions func(pageToken string) error
	fetchSubscriptions = func(pageToken string) error {
		call := service.Subscriptions.List([]string{"snippet"}).Mine(true).MaxResults(50)
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}

		response, err := call.Do()
		if err != nil {
			return err
		}

		for _, item := range response.Items {
			allSubscribedChannels = append(allSubscribedChannels, Channel{
				Name: item.Snippet.Title,
				ID:   item.Snippet.ResourceId.ChannelId,
			})
		}

		if response.NextPageToken != "" {
			return fetchSubscriptions(response.NextPageToken)
		}

		return nil
	}

	// Start fetching subscriptions from the first page
	err = fetchSubscriptions("")
	if err != nil {
		log.Fatalf("Error fetching subscriptions: %v", err)
	}

	// Create a CSV file to save the subscribed channels
	file, err := os.Create("subscribed_channels.csv")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header to the CSV file
	header := []string{"Channel Name", "Channel ID"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Error writing header to CSV: %v", err)
	}

	// Write the channel data to the CSV file
	for _, channel := range allSubscribedChannels {
		record := []string{channel.Name, channel.ID}
		if err := writer.Write(record); err != nil {
			log.Fatalf("Error writing record to CSV: %v", err)
		}
	}

	fmt.Println("Subscribed channel names and IDs have been saved to subscribed_channels.csv")
}
