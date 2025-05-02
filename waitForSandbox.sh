#!/bin/bash

echo "Waiting for Docker daemon to be ready..."
MAX_ATTEMPTS=30
ATTEMPT=0

while [ $ATTEMPT -lt $MAX_ATTEMPTS ]; do
  echo "Attempt $((ATTEMPT+1))/$MAX_ATTEMPTS: Checking Docker daemon..."
  
  if docker -H tcp://docker-sandbox:2375 info > /dev/null 2>&1; then
    echo "Docker daemon is ready and running!"
    exit 0
  fi
  
  ATTEMPT=$((ATTEMPT+1))
  echo "Docker daemon not ready yet. Waiting 1 second..."
  sleep 1
done

echo "ERROR: Docker daemon failed to start within the timeout period."
exit 1