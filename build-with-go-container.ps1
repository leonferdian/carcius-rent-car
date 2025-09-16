$ErrorActionPreference = "Stop"

# Parameters
$serviceName = "api-gateway"
$servicePath = "./services/$serviceName"
$imageName = "carcius-rent-car-$($serviceName.ToLower())"

# Create a temporary directory for the build context
$tempDir = Join-Path $env:TEMP "carcius-build-$([System.Guid]::NewGuid())"
New-Item -ItemType Directory -Path $tempDir -Force | Out-Null

try {
    Write-Host "[1/4] Copying files to temporary directory..."
    Copy-Item -Path "$servicePath\*" -Destination $tempDir -Recurse -Force

    Write-Host "[2/4] Initializing Go modules..."
    docker run --rm -v "${tempDir}:/app" -w /app golang:1.22-alpine sh -c "
        go mod init temp-module 2>/dev/null || true
        go mod tidy
        go mod download
        go mod verify
    "

    Write-Host "[3/4] Building Docker image..."
    docker build -t $imageName -f "$tempDir/Dockerfile.simple" $tempDir

    Write-Host "[4/4] Build completed successfully!"
    Write-Host "Image name: $imageName"
}
catch {
    Write-Error "Build failed: $_"
    exit 1
}
finally {
    # Clean up
    if (Test-Path $tempDir) {
        Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue
    }
}
