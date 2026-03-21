#!/bin/sh
# Start the asynq worker in the background
./worker &

# Start the HTTP server in the foreground
exec ./main
