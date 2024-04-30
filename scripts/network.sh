#!/bin/bash

export NONROOT_CONTAINER_NAME="non-root-container"
export USER_NAME="john"

function nonroot_op(){
    local cmd=$1
    case $cmd in
        "shell")
            docker run -it --rm -w /home/${USER_NAME} $NONROOT_IMAGE_NAME /bin/bash
            ;;
        "clean")
            docker rm -f ${NONROOT_CONTAINER_NAME}
            ;;
        *)
            echo "$0 nonroot [command]

command:
    shell  into a docker
    clean  container"
            ;;
    esac
}

export REST_NETWORK_NAME="learn-docker_rest"

function rest_network(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployments/restserver.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/restserver.yaml down
            ;;
        *)
            echo "$0 rest [command]

command:
    start   restserver network
    stop    restserver network"
            ;;
    esac
}