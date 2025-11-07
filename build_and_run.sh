#!/bin/bash

set -e

# configuration
APP_NAME="core-service"
IMAGE_TAG="latest"
PORT=8002

# build
echo "Building Docker image for $APP_NAME..."
docker build -t ${APP_NAME}:${IMAGE_TAG} .
echo "Build complete: ${APP_NAME}:${IMAGE_TAG}"

echo "Running container..."
# stop any existing container with same name
docker rm -f ${APP_NAME} 2>/dev/null || true

# run
docker run -d \
  --name ${APP_NAME} \
  -p ${PORT}:${PORT} \
  --network chirp-network \
  ${APP_NAME}:${IMAGE_TAG}

echo "${APP_NAME} is running on http://localhost:${PORT}"
