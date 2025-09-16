# Build all Go services
$services = @("api-gateway", "users-service", "cars-service", "bookings-service")

foreach ($service in $services) {
    Write-Host "Building $service..." -ForegroundColor Cyan
    
    $servicePath = "./services/$service"
    $imageName = "carcius-rent-car-$($service.ToLower())"
    
    # Build the Docker image
    docker build -t $imageName -f "$servicePath/Dockerfile" $servicePath
    
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Failed to build $service"
        exit 1
    }
    
    Write-Host "Successfully built $service" -ForegroundColor Green
}

Write-Host "\nAll services built successfully!" -ForegroundColor Green
Write-Host "You can now start the application with: docker-compose up" -ForegroundColor Green
