#!/bin/bash
TARGET=$1

if [ -z "$TARGET" ]; then
  echo "Usage: ./run-bench.sh [gin|hertz]"
  exit 1
fi

RESULT_DIR="results"
mkdir -p $RESULT_DIR

echo "POST,CPU,MEM" > ${RESULT_DIR}/${TARGET}_resource.csv

echo "🚀 Starting benchmark for $TARGET..."

cd ../$TARGET && podman-compose up -d --build
echo "Wait 10s for warmup..."
sleep 10

echo "Monitoring resources..."
(while true; do 
    podman stats --no-stream --format "{{.CPUPerc}},{{.MemUsage}}" ${TARGET}-gin-server-1 >> ../benchmark/${RESULT_DIR}/${TARGET}_resource.csv 2>/dev/null; 
    sleep 1; 
done) &
MONITOR_PID=$!

echo "Running k6 attack..."
k6 run -e BASE_URL=http://localhost:8080 ../benchmark/k6-script.js --summary-export=../benchmark/${RESULT_DIR}/${TARGET}_result.json

kill $MONITOR_PID
podman-compose down

echo "✅ Finished. Results: ${RESULT_DIR}/${TARGET}_result.json, ${RESULT_DIR}/${TARGET}_resource.csv"
