#!/usr/bin/env bash

DOCKER_NAME=indoquran-api

docker stop $DOCKER_NAME
docker rm $DOCKER_NAME
docker rmi $DOCKER_NAME

docker run --name redis -d -p 6379:6379 redis redis-server --appendonly yes
docker run -itd -p 27017:27017 --name mongo -v /c/Temp/data/docker/mongo:/bitnami/mongodb bitnami/mongodb
docker build --build-arg ENV=development . -t $DOCKER_NAME
docker run -it -d -p 8080:8080 --env ENV=development --name $DOCKER_NAME --link mongo --link redis $DOCKER_NAME

printf "y\n" | docker system prune