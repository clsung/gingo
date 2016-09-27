#!/usr/bin/env bash

docker build -t clsung/gingo_build build/
docker run -d --name gingo_build clsung/gingo_build
docker cp gingo_build:/go/bin/gingo run/
docker cp gingo_build:/go/bin/gingo_nopool run/
docker stop gingo_build
docker rm gingo_build
docker rmi clsung/gingo_build
docker build -t clsung/gingo run/
