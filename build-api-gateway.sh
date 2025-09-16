#!/bin/bash
set -e

cd services/api-gateway

echo "Building API Gateway..."
docker build -t carcius-rent-car-api-gateway .

echo "API Gateway built successfully!"
