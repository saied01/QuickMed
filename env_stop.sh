#!/bin/bash
set -e

echo "Stopping Go server..."
pkill -f "go run backend/main.go" || true

echo "Stopping Docker Compose containers..."
sudo docker compose down

echo "✅ Environment stopped."
