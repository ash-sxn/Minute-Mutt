package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ash-sxn/Minute-Mutt/pkg/auth"
	"github.com/ash-sxn/Minute-Mutt/pkg/downloader"
	"github.com/ash-sxn/Minute-Mutt/pkg/util"
	myYoutube "github.com/ash-sxn/Minute-Mutt/pkg/youtube" // Alias for local youtube package
	"google.golang.org/api/youtube/v3"
)

const toDownloadFilePath = "ToDownload"

func main() {

	// Check if we should download videos
	shouldDownload, err := checkDownloadPreference(toDownloadFilePath)
	if err != nil {
		log.Fatalf("Error checking download preference: %v", err)
	}
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
		videos, err := myYoutube.FetchLatestVideos(service, channel.ID, 5) // Fetch the latest 1 videos
		if err != nil {
			log.Printf("Error fetching latest videos for channel %s: %v", channel.Name, err)
			continue
		}

		for _, video := range videos {
			videoQueue.AddVideo(video)
		}
	}

	outputDir := os.Getenv("OUTPUT_DIR")
	maxResolution := os.Getenv("MAX_RESOLUTION")
	csvHistoryFilename := "pkg/database/history_queue.csv"

	// Load the history of downloaded videos
	downloadedVideos, err := util.LoadDownloadedVideos(csvHistoryFilename)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Error loading downloaded videos history: %v", err)
		}
		// If the file does not exist, proceed with an empty list
		downloadedVideos = make(map[string]struct{})
	}

	// Print the list of videos in the queue and download if not already downloaded
	fmt.Println("Queue of latest videos:")
	for _, video := range videoQueue.Videos {
		fmt.Printf("ID: %s, Title: %s\n", video.ID, video.Title)
		if _, alreadyDownloaded := downloadedVideos[video.ID]; !alreadyDownloaded {
			if shouldDownload {
				// Download the video if the preference is set to true
				err := downloader.DownloadVideo(video.ID, outputDir, maxResolution)
				if err != nil {
					// Print the error message if the download fails
					fmt.Printf("Error downloading video ID %s: %v\n", video.ID, err)
				}
			}
			// Add video to CSV regardless of download preference
			util.AddVideoToCSV(video, csvHistoryFilename)
		} else {
			fmt.Printf("Video %s has already been downloaded. Skipping.\n", video.ID)
		}
	}

	// At the end of the first run, write 'true' to the ToDownload file to enable future downloads
	if !shouldDownload {
		if err := ioutil.WriteFile(toDownloadFilePath, []byte("true"), 0644); err != nil {
			log.Fatalf("Error writing to ToDownload file: %v", err)
		}
	}
}

func checkDownloadPreference(filePath string) (bool, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, it's the first run, and we do not download videos
			return false, nil
		}
		return false, err
	}

	// If the file contains 'true', return true
	return strings.TrimSpace(string(content)) == "true", nil
}
