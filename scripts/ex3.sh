#!/bin/sh

export EX3_1_NETWORK_NAME="learn-docker_ex3-1"
export EX3_2_NETWORK_NAME="learn-docker_ex3-2"

function ex3_ops(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployments/ex3.yaml down
            docker network rm ${EX3_1_NETWORK_NAME}
            docker network rm ${EX3_2_NETWORK_NAME}
            ;;
        "shell")
            docker-compose -f ./deployments/ex3.yaml exec -it ex3_1 /bin/sh
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
    shell   into node ex3_1
    start   restserver network
    stop    restserver network"
            ;;
    esac
}

