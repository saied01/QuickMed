#!/bin/bash
set -e # salir si hay error

# --- Chequear si Docker está corriendo ---
if ! systemctl is-active --quiet docker; then
  echo "Docker not initialised. Initoating Docker..."
  sudo systemctl start docker
  echo "Docker initiated."
else
  echo "Docker already running."
fi

echo "Initializing Docker Compose..."
sudo docker compose up -d

echo "Waiting for PostgreSQL to be ready..."

# Obtener datos de conexión (ajusta si tu docker-compose los cambia)
PG_CONTAINER_NAME="quick-med-postgres-1"
PG_USER="devuser"
PG_DB="devdb"
PG_PASSWORD="devpass"

# Loop hasta que psql pueda conectarse
until sudo docker exec -e PGPASSWORD=$PG_PASSWORD $PG_CONTAINER_NAME psql -U $PG_USER -d $PG_DB -c '\q' >/dev/null 2>&1; do
  echo "PostgreSQL not ready yet. Waiting..."
  sleep 2
done

echo "PostgreSQL is ready!"
#
# echo "Starting Go server in background..."
# cd backend
# # go server en background, logs en server.log
# nohup go run main.go >server.log 2>&1 &
# cd ..

echo "Environment ready."
