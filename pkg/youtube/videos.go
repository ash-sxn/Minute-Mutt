package youtube

import (
	"errors"

	"google.golang.org/api/youtube/v3"
)

// Video represents a YouTube video with its ID and title.
type Video struct {
	ID    string
	Title string
}

// FetchLatestVideos fetches the latest videos for a given channel ID.
func FetchLatestVideos(service *youtube.Service, channelID string, maxResults int64) ([]Video, error) {
	// First, retrieve the channel to get the Uploads playlist ID
	channelCall := service.Channels.List([]string{"contentDetails"}).Id(channelID)
	channelResponse, err := channelCall.Do()
	if err != nil {
		return nil, err
	}
	if len(channelResponse.Items) == 0 {
		return nil, errors.New("channel not found")
	}
	uploadsPlaylistID := channelResponse.Items[0].ContentDetails.RelatedPlaylists.Uploads

	// Now, use the PlaylistItems.list method to get the latest videos from the Uploads playlist
	playlistCall := service.PlaylistItems.List([]string{"contentDetails", "snippet"}).
		PlaylistId(uploadsPlaylistID).
		MaxResults(maxResults)

	response, err := playlistCall.Do()
	if err != nil {
		return nil, err
	}

	var videos []Video
	for _, item := range response.Items {
		videos = append(videos, Video{
			ID:    item.ContentDetails.VideoId,
			Title: item.Snippet.Title,
		})
	}

	return videos, nil
}
