param (
    [Parameter(Mandatory=$true)]
    [string]$ServiceName
)

$servicePath = "./services/$ServiceName"
$dockerfilePath = "$servicePath/Dockerfile"

# Create go.sum if it doesn't exist
if (-not (Test-Path "$servicePath/go.sum")) {
    Write-Host "Generating go.sum for $ServiceName..."
    docker run --rm -v "${PWD}/$servicePath":/app -w /app golang:1.22-alpine sh -c "go mod tidy"
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Failed to generate go.sum for $ServiceName"
        exit 1
    }
}

# Build the Docker image
Write-Host "Building $ServiceName..."
docker build -t "carcius-rent-car-$($ServiceName.ToLower())" -f $dockerfilePath $servicePath

if ($LASTEXITCODE -eq 0) {
    Write-Host "Successfully built $ServiceName"
} else {
    Write-Error "Failed to build $ServiceName"
    exit 1
}
