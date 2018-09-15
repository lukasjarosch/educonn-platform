#!/usr/bin/env bash

# Set the service's name
SERVICE_NAME=lesson-api

VERSION=$1

LD_FLAGS="-X github.com/lukasjarosch/educonn-platform/${SERVICE_NAME}/internal/platform/config.Version=${VERSION}"

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Delete old binary
echo "==> Removing old $SERVICE_NAME binary..."
rm  -f cmd/$SERVICE_NAME/$SERVICE_NAME

# We want a statically linked binary
export CGO_ENABLED=0

# Build!
echo "==> Building $SERVICE_NAME..."
cd cmd/$SERVICE_NAME
go build -ldflags="${LD_FLAGS}" -a -installsuffix cgo -o $SERVICE_NAME main.go
