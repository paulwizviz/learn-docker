#!/bin/sh

export EX3_NETWORK_NAME="learn-docker_ex3"

function ex3_ops(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployments/ex3.yaml down
            docker rm -f ${EX3_CONTAINER_NAME}
            docker volume rm ${EX3_NETWORK_NAME}
            ;;
        "start")
            docker-compose -f ./deployments/ex3.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/ex3.yaml down
            ;;
        *)
            echo "$0 ex3 [command]

command:
    clean   stop network and clear container
    start   restserver network
    stop    restserver network"
            ;;
    esac
}

