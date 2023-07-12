#!/bin/bash

export IMAGE_NAME=learn-docker/gha

COMMAND=$1

case $COMMAND in
    "build")
        docker-compose -f ./build/gha/builder.yaml build
        ;;
    "clean")
        docker rmi -f ${IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "run")
        docker-compose -f ./deployment/gha/docker-compose.yaml up
        ;;
    "stop")
        docker-compose -f ./deployment/gha/docker-compose.yaml down
        ;;
    *)
        echo "Usage: $0 <Command>"
        echo
        echo "Command:"
        echo "  build   image"
        echo "  clean   image"
        echo "  run     container"
        echo "  stop    container"
        ;;
esac