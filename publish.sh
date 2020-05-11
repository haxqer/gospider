#!/bin/bash
# macOS compatible

prjName="gspider"

remoteTargetFolder="/data/update/${prjName}/"`date +"%y%m%d"`

ssh -T -t  root@172.31.0.1 'mkdir -p '${remoteTargetFolder}

find . -type f -name "gspider-*.zip"  -exec ls -t {} + | head -n 1 | xargs -I[] scp [] root@172.31.0.1:${remoteTargetFolder}
find . -type f -name "itemsave-*.zip" -exec ls -t {} + | head -n 1 | xargs -I[] scp [] root@172.31.0.1:${remoteTargetFolder}
find . -type f -name "spiderhttp-*.zip" -exec ls -t {} + | head -n 1 | xargs -I[] scp [] root@172.31.0.1:${remoteTargetFolder}

ssh -T -t  root@172.31.0.1 'ls -hal '${remoteTargetFolder}'/*.zip'
