#!/bin/sh

shDir=$(cd `dirname $0`; pwd)
strFDate=`date +"%y%m%d"`
#prjName=$(basename "$PWD")

ver=$(git rev-parse --short HEAD)

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
for goos in darwin linux
do
    for goarch in amd64
    do
        GOOS=${goos} GOARCH=${goarch} go build -o ./gspider_${goos}-${goarch}
        GOOS=${goos} GOARCH=${goarch} go build -o ./itemsave_${goos}-${goarch} ./services/itemsave
        GOOS=${goos} GOARCH=${goarch} go build -o ./spiderhttp_${goos}-${goarch} ./services/spiderhttp
    done
done

rm -f gspider-${strFDate}-${ver}.zip
rm -f itemsave-${strFDate}-${ver}.zip
rm -f spiderhttp-${strFDate}-${ver}.zip

zip -r gspider-${strFDate}-${ver}.zip    gspider_linux-amd64    conf/app.gspider.ini
zip -r itemsave-${strFDate}-${ver}.zip   itemsave_linux-amd64   conf/app.itemsave.ini
zip -r spiderhttp-${strFDate}-${ver}.zip spiderhttp_linux-amd64 conf/app.spiderhttp.ini

echo '----------------------------------------------'
unzip -l gspider-${strFDate}-${ver}.zip
unzip -l itemsave-${strFDate}-${ver}.zip
unzip -l spiderhttp-${strFDate}-${ver}.zip
