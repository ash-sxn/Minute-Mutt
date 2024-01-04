# Builder stage for ffmpeg
FROM debian:bullseye-slim as ffmpeg-builder

# Install curl and ca-certificates to download files, and xz-utils to extract the tarball
RUN apt-get update && apt-get install -y curl ca-certificates xz-utils && \
    rm -rf /var/lib/apt/lists/*

# Download and extract the ffmpeg static build
RUN curl -L https://github.com/yt-dlp/FFmpeg-Builds/releases/download/autobuild-2024-01-02-14-09/ffmpeg-N-113171-g85b8d59ec7-linux64-gpl.tar.xz | \
    tar xJ --strip-components=1 -C /usr/local/bin

# Builder stage for Go application
FROM golang:1.20-bullseye as go-builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all the dependencies that are required
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN go build -o /main .

# Final stage
FROM python:3.10-slim-bullseye

# Install curl and other necessary utilities
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

# Install yt-dlp
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/download/2023.12.30/yt-dlp -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp

# Copy ffmpeg and ffprobe binaries from the ffmpeg-builder stage
COPY --from=ffmpeg-builder /usr/local/bin/bin/ffmpeg /usr/local/bin/ffmpeg
COPY --from=ffmpeg-builder /usr/local/bin/bin/ffprobe /usr/local/bin/ffprobe

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the go-builder stage
COPY --from=go-builder /main /app/main

# Set default environment variables needed for our image
ENV OUTPUT_DIR="/watch"
ENV MAX_RESOLUTION="144"
ENV CRON_SCHEDULE="59 23 31 2 *"

# Copy the script that will set up the cron job
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Install cron and clean up in one layer to reduce image size
RUN apt-get update && apt-get install -y cron && \
    rm -rf /var/lib/apt/lists/*

# Start the entrypoint script
CMD ["/entrypoint.sh"]
