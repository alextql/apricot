#!/bin/bash
set -v
set -e

cd /app
go env -w GOPROXY="https://goproxy.io,direct"

CGO_ENABLED=1
GOARCH=amd64 GOOS=linux go build -v -ldflags "-w -s" -o dist/apricot_cgo_linux_amd64 .
GOARCH=arm64 GOOS=linux go build -v -ldflags "-w -s" -o dist/apricot_cgo_linux_arm64 .
