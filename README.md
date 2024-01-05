# Minute-Mutt

Minute-Mutt is a Golang application designed to help combat YouTube addiction by automatically downloading new videos from subscribed YouTube channels. By using the YouTube API and yt-dlp, Minute-Mutt ensures that you can watch your favorite content without the distractions of the YouTube platform.

## Features

- **Automatic Downloads**: Downloads new videos from your subscribed channels using the YouTube API.
- **Local Storage**: Saves videos locally, avoiding the need to surf YouTube and potentially get distracted.
- **Docker Support**: Includes a Dockerfile for easy deployment and a Docker image hosted on Docker Hub (`ash191245141/minute-mutt-cron`).
- **Cron Scheduling**: Uses cron to schedule regular checks and downloads at specified times (because I have free internet between 12:00 AM and 6:00 AM).
- **Customizable Quality**: Allows setting the maximum resolution for downloaded videos.

## Installation

To use Minute-Mutt, you can either run it locally or use the provided Docker image(recommended).

### Local Installation

1. Clone the repository to your local machine.
2. Ensure you have Golang, yt-dlp, and ffmpeg installed.
3. Obtain a `client_secret.json` file for YouTube API authentication:
   - Please don't misuse my key, it's a restricted key with a limited quota per day for the free tier
   - You can create your own `client_secret.json` by setting up a project in the Google Developers Console and enabling the YouTube Data API v3.
   - Alternatively, you can contact me to add your email to the authorized users list. Reach out on Twitter at [@ash_sxn](https://twitter.com/ash_sxn) or email me at ash.191245141@gmail.com.
5. Place the `client_secret.json` file in the root directory of the project.
6. Run `go run main.go` and follow the prompts to authenticate with the YouTube API.

### Docker Installation

1. Pull the Docker image from Docker Hub: `docker pull ash191245141/minute-mutt-cron`
2. Run the Docker container with the required environment variables:`docker run -it -e MAX_RESOLUTION="1080" -e CRON_SCHEDULE="0 0-6 * * *" -v <local-download-location>:/watch ash191245141/minute-mutt-cron:1.0`

Replace `MAX_RESOLUTION` and `CRON_SCHEDULE` with your preferred settings, and replace `<local-download-locatoin>` with the directory where you want downloaded videos to be saved, like `docker run -it -e MAX_RESOLUTION="1080" -e CRON_SCHEDULE="*/2 * * * *" -v ~/watch:/watch ash191245141/minute-mutt-cron:1.0` will download videos to the `~/watch` directory with 1080p quality and check for new videos every 2 minutes(get familiar with cron to understand this) and download them to the `~/watch` directory if available.

## Usage

After installation, Minute-Mutt will automatically check for new videos from your subscribed channels based on the CRON_SCHEDULE. Downloaded videos will be saved to the specified output directory.

### Configuration

- `MAX_RESOLUTION`: Sets the maximum resolution for downloaded videos (e.g., "1080").
- `CRON_SCHEDULE`: Defines when the program should check for and download new videos (e.g., "*/2 * * * *").
- `OUTPUT_DIR`: Sets the directory where downloaded videos will be saved (default is `/watch` in the Docker container).

## Logs

Logs are saved to `/var/log/cron.log` in the Docker container, which includes any errors encountered during the download process.

## Contributing

Contributions to Minute-Mutt are welcome. Please feel free to fork the repository, make changes, and submit a pull request.

## License

Minute-Mutt is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- [yt-dlp](https://github.com/yt-dlp/yt-dlp): Command-line program to download videos from YouTube and other video sites.
- [YouTube API](https://developers.google.com/youtube/v3): API services provided by YouTube for programmatic access to YouTube content.

