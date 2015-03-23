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
        echo "Getting godep..."
	go get github.com/tools/godep
else
        echo "Not on Linux - skipping demoapp build"
fi


if [ ! -e '.git' ]; then
	echo '.git directory not found'
	echo 'git init . && git add -A . && git commit -m "To test godep"'
	exit 
fi

# echo "Making git repository..."
# 
echo "Running godep save..."
${GOBIN}/godep save github.com/nak3/demoapp
