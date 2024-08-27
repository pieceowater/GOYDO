FROM golang:1.23-alpine AS build

LABEL authors="pieceowater"

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o goydo main.go

# Use a minimal base image for the final build
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the build stage
COPY --from=build /app/goydo .

# Expose the port on which the app runs
EXPOSE 3000

# Run the Go application
CMD ["./goydo"]