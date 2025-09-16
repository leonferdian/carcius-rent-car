@echo off
echo Building API Gateway...
cd services\api-gateway
docker build -t carcius-rent-car-api-gateway .

if %ERRORLEVEL% EQU 0 (
    echo API Gateway built successfully!
) else (
    echo Failed to build API Gateway.
)

cd ..\..
