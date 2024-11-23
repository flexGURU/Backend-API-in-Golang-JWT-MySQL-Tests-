#!/bin/bash

# Define the 'run' task
run() {
  echo "Running the Go application..."
  go run ./cmd/main.go
}

# Handle command-line arguments (similar to Makefile targets)
case "$1" in
  run)
    run
    ;;
  *)
    echo "Usage: $0 {run}"
    exit 1
    ;;
esac
