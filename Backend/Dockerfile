# Use the official Golang image based on Alpine Linux for a lightweight container
FROM golang:1.23-alpine3.19

# Set the working directory inside the container
WORKDIR /src/app

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy the Go module files and download dependencies separately to leverage Docker's caching
# Separate Dependency Installation:
   #  Copying go.mod and go.sum before the source code (COPY . .) allows Docker to cache the dependency download step, making future builds faster if dependencies don’t change.


COPY go.mod go.sum ./
#go.sum contains checksums for each dependency, ensuring that the exact versions used are verified against these checksums to avoid tampering.
RUN go mod tidy   

# Copy the rest of the source code
COPY . .

# Specify the command to start the application using Air
# CMD ["air"]
