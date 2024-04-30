#!/bin/bash

export NONROOT_IMAGE_NAME="learn-docker/nonroot:current"
export RESTSERVER_IMAGE_NAME="learn-docker/restserver:current"

function build_nonroot(){
    docker-compose -f ./build/builder.yaml build nonroot
}

function build_restserver(){
    docker-compose -f ./build/builder.yaml build rest
}

function build_images(){
    docker-compose -f ./build/builder.yaml build
}

function clean_nonroot(){
    docker rmi -f ${NONROOT_IMAGE_NAME}
}

function clean_rest(){
    docker rmi -f ${RESTSERVER_IMAGE_NAME}
}

function clean_images(){
    clean_nonroot
    clean_rest
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function image(){
    local cmd=$1
    case $cmd in
        "build:nonroot")
            build_nonroot
            ;;
        "build:rest")
            build_restserver
            ;;
        "build")
            build_images
            ;;
        "clean:nonroot")
            clean_nonroot
            ;;
        "clean:rest")
            clean_rest
            ;;
        "clean")
            clean_images
            ;;
        *)
            echo "$0 image [command]

command:
    build:nonroot
    build:rest 
    build           all images 
    clean:nonroot
    clean:rest
    clean           all images"
    esac
}
