# Use the official Golang image based on Alpine Linux for a lightweight container
FROM golang:1.23-alpine3.19 as builder

# Set the working directory inside the container
WORKDIR /src/app

# Install Air for live reloading (optional for local development)
RUN go install github.com/air-verse/air@latest

# Copy the Go module files and download dependencies separately to leverage Docker's caching
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application from the 'cmd/api' directory
RUN go build -o app ./cmd/api

# Use a minimal base image for the final container
FROM alpine:3.19

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /src/app/app .

# Expose the default port (Render will dynamically assign a port via the PORT environment variable)
EXPOSE 3000

# Command to run the application (Render will pass the $PORT automatically)
CMD ["./app"]
