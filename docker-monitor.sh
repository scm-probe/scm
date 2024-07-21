#!/bin/bash

# Start the containers in detached mode
sudo docker-compose -f /home/saarthak/Desktop/scm-probe/scm-prom/monitor.docker-compose.yaml up -d

# Get the container IDs
CONTAINER_IDS=$(sudo docker-compose -f /home/saarthak/Desktop/scm-probe/scm-prom/monitor.docker-compose.yaml ps -q)

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




