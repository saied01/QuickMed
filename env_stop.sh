#!/bin/bash
set -e

echo "Stopping Docker Compose containers..."
sudo docker compose down

echo "Environment stopped."
