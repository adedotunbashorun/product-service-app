# Use an official Golang image as a base image
FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Install air from the new repository path
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the entrypoint to use air for live-reloading
CMD ["air"]

