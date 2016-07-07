#!/bin/bash
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 /usr/src/myapp/dockerBuild.sh
docker run --rm -v "$PWD":/app ubuntu myapp 
