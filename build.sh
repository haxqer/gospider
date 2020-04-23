#!/bin/sh

shDir=$(cd `dirname $0`; pwd)
strFDate=`date +"%y%m%d"`
#prjName=$(basename "$PWD")

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
for goos in darwin linux
do
    for goarch in amd64
    do
        GOOS=${goos} GOARCH=${goarch} go build -o ./gspider_${goos}-${goarch}
        GOOS=${goos} GOARCH=${goarch} go build -o ./itemsave_${goos}-${goarch} ./services/itemsave
    done
done

zip -r gspider-${strFDate}.zip  gspider_linux-amd64  docs conf/app.dist.ini README.md
zip -r itemsave-${strFDate}.zip itemsave_linux-amd64 docs conf/app.dist.ini README.md

echo '----------------------------------------------'
unzip -l gspider-${strFDate}.zip
unzip -l itemsave-${strFDate}.zip
