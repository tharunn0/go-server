# Step 1: Build the Go binary
FROM golang:1.25.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install git if needed
RUN apk add --no-cache git

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application statically with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server .

# Step 2: Create a lightweight runner image
FROM alpine:3.19

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Copy the default .env configuration file
COPY --from=builder /app/.env.example .

# Expose port 8080
EXPOSE 8080

# Run the server
ENTRYPOINT ["./server"]
