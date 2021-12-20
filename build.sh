#!/usr/bin/env bash

set -x

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

GitTag=$(git tag -l -n)
GitBranch=$(git branch --show-current)
GitCommit=$(git log --pretty=oneline -n 1)
BuildTime=$(date +'%Y-%m-%d %H:%M:%S')
GitStatus=$(git status -s)

BuildInfo=" \
    -X 'github.com/qingants/gin-skeleton/pkg/bininfo.GitTag=${GitTag}' \
    -X 'github.com/qingants/gin-skeleton/pkg/bininfo.GitBranch=${GitBranch}' \
    -X 'github.com/qingants/gin-skeleton/pkg/bininfo.GitCommit=${GitCommit}' \
    -X 'github.com/qingants/gin-skeleton/pkg/bininfo.BuildTime=${BuildTime}' \
    -X 'github.com/qingants/gin-skeleton/pkg/bininfo.GitStatus=${GitStatus}' \
"

go build -ldflags "${BuildInfo}"

echo 'build done.'