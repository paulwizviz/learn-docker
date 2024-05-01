#!/bin/sh

export EX2_NETWORK_NAME="learn-docker_ex2"
export EX2_CONTAINER_NAME="ex2-container-name"

function ex2_ops(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployments/ex2.yaml down
            docker rm -f ${EX2_CONTAINER_NAME}
            docker volume rm ${EX2_NETWORK_NAME}
            ;;
        "start")
            docker-compose -f ./deployments/ex2.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/ex2.yaml down
            ;;
        *)
            echo "$0 rest [command]

command:
    clean   stop network and clear container
    start   restserver network
    stop    restserver network"
            ;;
    esac
}

