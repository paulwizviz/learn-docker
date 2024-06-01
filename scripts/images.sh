#!/bin/sh

export NONROOT_IMAGE_NAME="go-docker/nonroot:current"
export RESTSERVER_IMAGE_NAME="go-docker/restserver:current"
export SOCKET_IMAGE_NAME="go-docker/socket:current"
export TOOLS_IMAGE_NAME="go-docker/tools:current"

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
        "build:tools")
            docker-compose -f ./build/builder.yaml build tools
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
        "clean:tools")
            docker rmi -f ${TOOLS_IMAGE_NAME}
            ;;
        "clean")
            image clean:nonroot
            image clean:rest
            image clean:socket
            image clean:tools
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        *)
            echo "Usage: $0 image [command]

command:
    build:nonroot   image for nonroot container
    build:rest      image with restserver
    build:socket    image with socket
    build:tools     image with networking tools
    build           all images
    clean:nonroot   clear image for nonroot container
    clean:rest      clear image with restserver
    clean:socket    clear image with socket
    clean:tools     clear image with tools
    clean           all images"
            ;;
    esac
}