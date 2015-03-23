#!/bin/bash -e

ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/demoapp"

if [ ! -h gopath/src/${REPO_PATH} ]; then
        mkdir -p gopath/src/${ORG_PATH}
        ln -s ../../../.. gopath/src/${REPO_PATH} || exit 255
fi

export GOBIN=${PWD}/bin
export GOPATH=${PWD}/gopath

eval $(go env)

if [ ${GOOS} = "linux" ]; then
        echo "Building demoapp..."
        go build -o ${GOBIN}/demoapp ${REPO_PATH}
else
        echo "Not on Linux - skipping demoapp build"
fi
