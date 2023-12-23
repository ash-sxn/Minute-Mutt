package youtube

import (
	"google.golang.org/api/youtube/v3"
)

// Channel holds the name and ID of a YouTube channel
type Channel struct {
	Name string
	ID   string
}

// FetchSubscriptions fetches all pages of subscribed channels and returns a slice of Channels
func FetchSubscriptions(service *youtube.Service) ([]Channel, error) {
	var allSubscribedChannels []Channel
	var fetch func(pageToken string) error

	fetch = func(pageToken string) error {
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
			return fetch(response.NextPageToken)
		}

		return nil
	}

	err := fetch("")
	if err != nil {
		return nil, err
	}

	return allSubscribedChannels, nil
}
