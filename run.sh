#!/bin/bash

# Default name in case no parameter is passed
NAME="World"

# Parse the --name parameter
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --name) NAME="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

# Run the Go application with the provided name
go run ./cmd/helloapp --name "$NAME"
