#!/bin/sh

shDir=$(cd `dirname $0`; pwd)

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
for goos in darwin freebsd linux
do
    for goarch in amd64
    do
        GOOS=${goos} GOARCH=${goarch} go build -o ./gspider_${goos}-${goarch} >/dev/null 2>&1
    done
done
