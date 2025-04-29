# Use Go 1.19 as base image
FROM golang:1.19.5-alpine

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o metric_service

# Expose the service port
EXPOSE 8080

# Command to run the built binary
CMD ["./metric_service"]
