#!/bin/sh

export NONROOT_IMAGE_NAME="learn-docker/nonroot:current"
export RESTSERVER_IMAGE_NAME="learn-docker/restserver:current"
export SOCKET_IMAGE_NAME="learn-docker/socket:current"

function image(){
    local cmd=$1
    case $cmd in
        "build:nonroot")
            docker-compose -f ./build/builder.yaml build nonroot
            ;;
        "build:rest")
            docker-compose -f ./build/builder.yaml build rest
            ;;
        "build:socket")
            docker-compose -f ./build/builder.yaml build socket
            ;;
        "build")
            docker-compose -f ./build/builder.yaml build
            ;;
        "clean:nonroot")
            docker rmi -f ${NONROOT_IMAGE_NAME}
            ;;
        "clean:rest")
            docker rmi -f ${RESTSERVER_IMAGE_NAME}
            ;;
        "clean:socket")
            docker rmi -f ${SOCKET_IMAGE_NAME}
            ;;
        "clean")
            docker rmi -f ${NONROOT_IMAGE_NAME}
            docker rmi -f ${RESTSERVER_IMAGE_NAME}
            docker rmi -f ${SOCKET_IMAGE_NAME}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        *)
            echo "Usage: $0 image [command]

command:
    build:nonroot   image for nonroot container
    build:rest      image with restserver
    build:socket    image with socket
    build           all images
    clean:nonroot   clear image for nonroot container
    clean:rest      clear image with restserver
    clean:socket    clear image with socket
    clean           all images"
            ;;
    esac
}