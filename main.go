package main

import (
	"fmt"
	"log"

	"github.com/ash-sxn/Minute-Mutt/pkg/auth"
	"github.com/ash-sxn/Minute-Mutt/pkg/util"
	myYoutube "github.com/ash-sxn/Minute-Mutt/pkg/youtube" // Alias for local youtube package
	"google.golang.org/api/youtube/v3"
)

func main() {
	// Corrected to use the alias for the official YouTube package
	const youtubeScope = youtube.YoutubeReadonlyScope
	client := auth.GetClient(youtubeScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	allSubscribedChannels, err := myYoutube.FetchSubscriptions(service)
	if err != nil {
		log.Fatalf("Error fetching subscriptions: %v", err)
	}

	videoQueue := util.NewVideoQueue()

	for _, channel := range allSubscribedChannels {
		videos, err := myYoutube.FetchLatestVideos(service, channel.ID, 1) // Fetch the latest 1 videos
		if err != nil {
			log.Printf("Error fetching latest videos for channel %s: %v", channel.Name, err)
			continue
		}

		for _, video := range videos {
			videoQueue.AddVideo(video)
		}
	}

	// Print the list of videos in the queue
	fmt.Println("Queue of latest videos:")
	for _, video := range videoQueue.Videos {
		fmt.Printf("ID: %s, Title: %s\n", video.ID, video.Title)
	}

	// Save the queue to a csv file
	csvFilename := "pkg/database/video_queue.csv"
	if err := videoQueue.SaveToCSV(csvFilename); err != nil {
		log.Fatalf("Failed to save video queue to CSV: %v", err)
	}
	fmt.Printf("Saved video queue to %s\n", csvFilename)

	// ... code to download videos and remove them from the queue ...
}
