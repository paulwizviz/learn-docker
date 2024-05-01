#!/bin/sh

export NONROOT_CONTAINER_NAME="non-root-container"
export USER_NAME="john"

function ex1_ops(){
    local cmd=$1
    case $cmd in
        "clean")
            docker rm -f ${NONROOT_CONTAINER_NAME}
            ;;
        "shell")
            docker run -it --rm -w /home/${USER_NAME} $NONROOT_IMAGE_NAME /bin/bash
            ;;
        *)
            echo "Usage: $0 ex1 [command]
            
    command:
        clean   remove images and containers
        shell   access container shell"
            ;;
    esac
}