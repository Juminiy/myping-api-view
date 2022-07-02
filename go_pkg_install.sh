#!/bin/bash
go mod tidy
go fmt
LATEST_VERSION='v1.0.4'
go get -u github.com/Juminiy/myping@$LATEST_VERSION