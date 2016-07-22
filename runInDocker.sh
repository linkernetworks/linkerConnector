#!/bin/bash
docker run --rm -v "$PWD":/usr/src/linkerConnector -w /usr/src/linkerConnector golang:1.6 /usr/src/linkerConnector/dockerBuild.sh
docker run --rm -v "$PWD":/app ubuntu /app/testSendKafka.sh
