#!/bin/sh
echo Building artifact
docker build -t removeit -f build.Dockerfile .

echo Extracting artifact from builder image
docker container create --name extract removeit
docker container cp extract:/go/src/github.com/nych/docker-in-a-nutshell/demos/builder-pattern/main ./main
docker container rm -f extract
docker image rm removeit

echo Building final docker image
docker build -t builder-pattern-demo:latest .
rm -f ./main