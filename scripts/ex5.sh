#!/bin/sh

export EX5_CONTAINER_NAME="ex5-container"
export EX5_NETWORK_1="ex5_net_1"
export EX5_NETWORK_2="ex5_net_2"

function ex5_ops(){
    local cmd=$1
    case $cmd in
        "clean")
            docker rm -f ${EX5_CONTAINER_NAME}
            docker network rm -f ${EX5_NETWORK_NAME}
            ;;
        "shell")
            docker-compose -f ./deployments/ex5.yaml run -it --rm ex5 /bin/sh
            ;;
        *)
            echo "Usage: $0 ex5 [command]
            
    command:
        clean   remove images and containers
        shell   access container shell"
            ;;
    esac
}