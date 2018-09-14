#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

echo "==> Building protocol buffers for $(basename $(dirname $(dirname "$SOURCE")))"
protoc  \
    -I. \
    --go_out=plugins=micro:$GOPATH/src/github.com/lukasjarosch/educonn-platform \
    --micro_out=:$GOPATH/src/github.com/lukasjarosch/educonn-platform \
    user-api/proto/*.proto
