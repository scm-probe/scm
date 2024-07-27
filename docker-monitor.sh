#!/bin/bash

# Check if the path to the docker-compose file is provided
if [ -z "$1" ]; then
  echo "Error: No docker-compose file path provided."
  echo "Usage: $0 /path/to/docker-compose.yaml"
  exit 1
fi

DOCKER_COMPOSE_FILE=$1

# Check if the provided file exists
if [ ! -f "$DOCKER_COMPOSE_FILE" ]; then
  echo "Error: File '$DOCKER_COMPOSE_FILE' does not exist."
  exit 1
fi


# Get the container IDs
CONTAINER_IDS=$(sudo docker-compose -f "$DOCKER_COMPOSE_FILE" ps -q)

# Print the container IDs
echo "Started containers with IDs:"
echo "$CONTAINER_IDS"

FIRST_CONTAINER_ID=$(echo "$CONTAINER_IDS" | head -n 1)

PIDS=$(sudo docker top $FIRST_CONTAINER_ID | awk 'NR>1 {print $2}')

echo "Process IDs:"
echo "$PIDS"

FIRST_PID=$(echo "$PIDS" | head -n 1)

echo ${FIRST_PID}

sudo ./main -id ${FIRST_PID}
