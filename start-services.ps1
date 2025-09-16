# Stop and remove any existing containers
Write-Host "Stopping and removing any existing containers..." -ForegroundColor Yellow
docker-compose down --remove-orphans

# Start the database first
Write-Host "\nStarting PostgreSQL database..." -ForegroundColor Cyan
docker-compose up -d postgres

# Wait for PostgreSQL to be ready
Write-Host "\nWaiting for PostgreSQL to be ready..." -ForegroundColor Cyan
$maxRetries = 30
$retryCount = 0
$dbReady = $false

while ($retryCount -lt $maxRetries -and -not $dbReady) {
    try {
        $result = docker-compose exec -T postgres pg_isready -U postgres -d carrental
        if ($LASTEXITCODE -eq 0) {
            $dbReady = $true
            Write-Host "PostgreSQL is ready!" -ForegroundColor Green
            break
        }
    } catch {
        # Ignore errors and retry
    }
    
    $retryCount++
    Write-Host "." -NoNewline -ForegroundColor Yellow
    Start-Sleep -Seconds 1
}

if (-not $dbReady) {
    Write-Error "\nFailed to connect to PostgreSQL after $maxRetries attempts"
    exit 1
}

# Start all services
Write-Host "\nStarting all services..." -ForegroundColor Cyan
docker-compose up -d

# Show status
Write-Host "\nServices status:" -ForegroundColor Cyan
docker-compose ps

Write-Host "\nApplication is starting up..." -ForegroundColor Green
Write-Host "Frontend will be available at http://localhost:3000" -ForegroundColor Green
Write-Host "API Gateway is available at http://localhost:8080" -ForegroundColor Green

# Show logs
Write-Host "\nTailing logs (Ctrl+C to stop watching logs)..." -ForegroundColor Cyan
Start-Sleep -Seconds 2
docker-compose logs -f
