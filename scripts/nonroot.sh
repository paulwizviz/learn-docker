#!/bin/bash

COMMAND=$1

. ./scripts/common.sh

export NONROOT_IMAGE_NAME="learn-k8s/nonroot:current"
export NONROOT_CONTAINER_NAME="non-root-container"
export USER_NAME="john"

case $COMMAND in
    "build")
        docker-compose -f ./build/nonroot/builder.yaml build
        ;;
    "clean")
        docker rmi -f ${NONROOT_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "shell")
        docker run -it --rm -w /home/${USER_NAME} $NONROOT_IMAGE_NAME /bin/bash
        ;;
    *)
        echo "Usage; $0 <Command>"
        echo
        echo "Command:"
        echo " shell   into the container"
        ;;
esac