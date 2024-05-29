# Dockerfile for Go Backend

# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Expose port 4000
EXPOSE 4000

# Command to run the Go application
CMD ["/docker-gs-ping"]
