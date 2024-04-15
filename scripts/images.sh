#!/bin/bash

export NONROOT_IMAGE_NAME="learn-docker/nonroot:current"
export RESTSERVER_IMAGE_NAME="learn-docker/restserver:current"


function build_images(){
    docker-compose -f ./build/builder.yaml build
}
        
function clean_images(){
    docker rmi -f ${GHA_IMAGE_NAME}
    docker rmi -f ${NONROOT_IMAGE_NAME}
    docker rmi -f ${RESTSERVER_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function image(){
    local cmd=$1
    case $cmd in
        "build")
            build_images
            ;;
        "clean")
            clean_images
            ;;
        *)
            echo "$0 image [command]

command:
    build  images
    clean  images"
    esac
}
