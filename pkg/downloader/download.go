package downloader

import (
	"fmt"

	"github.com/ash-sxn/Minute-Mutt/pkg/util" // Replace with your actual module path
)

func DownloadVideo(videoID string) error {
	downloaded, err := util.IsDownloaded(videoID)
	if err != nil {
		return err
	}
	if downloaded {
		fmt.Printf("Video %s has already been downloaded.\n", videoID)
		return nil
	}

	// Your existing download logic here...

	// After successful download:
	err = util.MarkAsDownloaded(videoID)
	if err != nil {
		return err
	}

	fmt.Printf("Video %s downloaded and marked as such.\n", videoID)
	return nil
}
