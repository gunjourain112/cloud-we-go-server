#!/bin/bash
TARGET=$1

if [ -z "$TARGET" ]; then
  echo "Usage: ./run-bench.sh [gin|hertz]"
  exit 1
fi

PORT=8080

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
RESULT_DIR="${SCRIPT_DIR}/results"
mkdir -p "$RESULT_DIR"
chmod 777 "$RESULT_DIR"

CONTAINER_ID=$(podman ps --filter "name=${TARGET}" --format "{{.ID}}" | head -n 1)

if [ -z "$CONTAINER_ID" ]; then
  echo "Error: No running container found for $TARGET"
  exit 1
fi

echo "Target detected: $TARGET (ID: $CONTAINER_ID, Port: $PORT)"
echo "POST,CPU,MEM" > "${RESULT_DIR}/${TARGET}_resource.csv"

echo "Monitoring resources..."
(while true; do 
    podman stats --no-stream --format "{{.CPUPerc}},{{.MemUsage}}" $CONTAINER_ID >> "${RESULT_DIR}/${TARGET}_resource.csv" 2>/dev/null; 
    sleep 1; 
done) &
MONITOR_PID=$!

echo "Running k6 attack via Podman..."
podman run --rm -i --network=host \
    -v "${SCRIPT_DIR}:/app:Z" \
    -w /app \
    grafana/k6 run -e BASE_URL=http://localhost:${PORT} k6-script.js --summary-export="results/${TARGET}_result.json"

kill $MONITOR_PID
echo "Finished. Results saved in benchmark/results/"
