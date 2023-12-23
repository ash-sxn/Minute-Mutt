package youtube

import (
	"google.golang.org/api/youtube/v3"
)

// Video represents a YouTube video with its ID and title.
type Video struct {
	ID    string
	Title string
}

// FetchLatestVideos fetches the latest videos for a given channel ID.
func FetchLatestVideos(service *youtube.Service, channelID string, maxResults int64) ([]Video, error) {
	call := service.Search.List([]string{"id", "snippet"}).
		ChannelId(channelID).
		MaxResults(maxResults).
		Order("date").
		Type("video")

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []Video
	for _, item := range response.Items {
		videos = append(videos, Video{
			ID:    item.Id.VideoId,
			Title: item.Snippet.Title,
		})
	}

	return videos, nil
}
