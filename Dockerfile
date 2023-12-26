# Start from the official Golang image
FROM golang:1.20

# Install Python, pip, and any other necessary packages
RUN apt-get update && apt-get install -y \
    python3 \
    python3-venv \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create a virtual environment for Python packages
RUN python3 -m venv /opt/venv

# Install yt-dlp in the virtual environment
RUN /opt/venv/bin/pip install --no-cache-dir yt-dlp

# Set the PATH to include the virtual environment's bin directory
ENV PATH="/opt/venv/bin:$PATH"

# Set default environment variables needed for our image
ENV OUTPUT_DIR="/watch"
ENV MAX_RESOLUTION="144"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies that are required
RUN go mod download

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
