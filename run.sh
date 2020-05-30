#!/usr/bin/env bash

docker stop go-gin
docker rm go-gin
docker rmi go-gin

docker build . -t go-gin
docker run -i -t -p 8080:8080 --name go-gin go-gin