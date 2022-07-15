#!/bin/bash

d=$(cd "$(dirname "$0")"; pwd)
pushd $d/..

dist=dist
rm -rf ${dist} && true
mkdir ${dist} && true

rsrc -manifest elevation.manifest -o elevation.syso
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${dist}/elevation.exe -v -ldflags "-w -s"

popd
