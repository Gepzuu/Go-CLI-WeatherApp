# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.18 AS builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
# -o myapp specifies the output name of the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp /myapp

# Run the web service on container startup.
CMD ["/myapp"]