# Use the official Golang image for building
FROM golang:1.23.4 AS builder

# Set working directory
WORKDIR /

# Copy everything to the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o myapp

# Use a minimal base image for running the application
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port (change it if necessary)
EXPOSE 8080

# Run the application
CMD ["./myapp"]