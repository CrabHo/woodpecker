#!/bin/bash

set -e
set -x

VERSION=$DRONE_TAG

if [ -z "$VERSION" ]; then
  VERSION=${DRONE_COMMIT_SHA:0:8}
fi

echo "Building $VERSION"

go build -ldflags '-extldflags "-static" -X github.com/laszlocph/woodpecker/version.Version='${VERSION} -o release/drone-server github.com/laszlocph/woodpecker/cmd/drone-server
GOOS=linux GOARCH=amd64 CGO_ENABLED=0         go build -ldflags '-X github.com/laszlocph/woodpecker/version.Version='${VERSION} -o release/drone-agent             github.com/laszlocph/woodpecker/cmd/drone-agent
go build -ldflags '-extldflags "-static" -X main.version='${VERSION} -o release/drone ./cli/drone
