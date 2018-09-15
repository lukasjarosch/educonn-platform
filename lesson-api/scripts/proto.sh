#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd "$DIR/proto"

echo "==> Building protocol buffers for $(basename $(dirname $(dirname "$SOURCE")))"
protoc  \
    -I. \
    --go_out=plugins=micro:$GOPATH/src \
    --micro_out=:$GOPATH/src \
    *.proto
