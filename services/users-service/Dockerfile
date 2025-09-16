# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod .
RUN go mod tidy

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o users-service .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/users-service .

# Expose port
EXPOSE 8081

# Command to run the executable
CMD ["./users-service"]
