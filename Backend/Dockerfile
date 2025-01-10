# Use an official Golang image for building
FROM golang:1.23-alpine as builder

# Set the working directory
WORKDIR /app

# Install necessary tools and dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal base image for the final container
FROM alpine:3.19

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file (Optional: only for local testing)
# COPY .env . 

# Expose the port your app runs on
EXPOSE 3000

# Command to run the application
CMD ["./main"]
