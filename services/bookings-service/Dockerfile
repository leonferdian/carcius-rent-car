# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod .
RUN go mod tidy

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o bookings-service .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/bookings-service .

# Expose port
EXPOSE 8083

# Command to run the executable
CMD ["./bookings-service"]
