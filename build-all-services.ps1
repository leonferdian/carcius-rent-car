# Build all Go services using the same method that worked for api-gateway
$services = @("users-service", "cars-service", "bookings-service")

foreach ($service in $services) {
    Write-Host "`n[1/4] Building $service..."
    $servicePath = "./services/$service"
    $imageName = "carcius-rent-car-$($service.ToLower())"
    
    # Create a temporary directory for the build context
    $tempDir = Join-Path $env:TEMP "carcius-build-$([System.Guid]::NewGuid())"
    New-Item -ItemType Directory -Path $tempDir -Force | Out-Null

    try {
        Write-Host "  - Copying files to temporary directory..."
        Copy-Item -Path "$servicePath\*" -Destination $tempDir -Recurse -Force

        # Use the simple Dockerfile for building
        if (Test-Path "$tempDir\Dockerfile.simple") {
            $dockerfile = "Dockerfile.simple"
        } else {
            # Create a simple Dockerfile if it doesn't exist
            @"
FROM golang:1.22-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Set up Go modules
ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct \
    GOSUMDB=sum.golang.org

# Copy only go.mod first for better caching
COPY go.mod .

# Download dependencies
RUN go mod download \
    && go mod verify

# Copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o $service .

# Command to run the executable
CMD ["./$service"]
"@ | Out-File -FilePath "$tempDir\Dockerfile.simple" -Encoding utf8
            $dockerfile = "Dockerfile.simple"
        }

        Write-Host "  - Building Docker image..."
        docker build -t $imageName -f "$tempDir\$dockerfile" $tempDir

        if ($LASTEXITCODE -eq 0) {
            Write-Host "  - Successfully built $service as $imageName"
        } else {
            Write-Error "  - Failed to build $service"
            exit 1
        }
    }
    catch {
        Write-Error "  - Error building $service : $_"
        exit 1
    }
    finally {
        # Clean up
        if (Test-Path $tempDir) {
            Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue
        }
    }
}

Write-Host "`nAll services built successfully!"
