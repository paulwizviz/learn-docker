#!/bin/bash

COMMAND=$1

. ./scripts/common.sh

case $COMMAND in
    "build")
        docker-compose -f ./build/shared/builder.yaml build
        ;;
    "clean")
        docker rmi -f ${RESTSERVER_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo "Usage: $0 <Command>"
        echo
        echo "Command:"
        echo " build      image"
        echo " clean      images"
        exit 1
        ;;
esac