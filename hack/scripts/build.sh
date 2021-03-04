#!/bin/bash
set -v
set -e

CGO_ENABLED=0
GOARCH=amd64
ORI_GOPROXY=$(go env GOPROXY)
go env -w GOPROXY="https://goproxy.io,direct"

GOOS=darwin go build -v -o dist/apricot_darwin_amd64 .
GOOS=linux go build -v -ldflags "-w -s" -o dist/apricot_linux_amd64 .
GOOS=windows go build -v -ldflags "-w -s" -o dist/apricot_windows_amd64.exe .
GOOS=linux GOARCH=arm64 go build -v -ldflags "-w -s" -o dist/apricot_linux_arm64 .

go env -w GOPROXY="${ORI_GOPROXY}"
