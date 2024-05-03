#!/bin/sh

export EX4_NETWORK_NAME="learn-docker_ex4"
export EX4_1_CONTAINER="ex4_1"
export EX4_2_CONTAINER="ex4_2"
export EX4_SHARED_VOL="ex4_shared_vol"

function ex4_ops(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployments/ex4.yaml down
            docker rm -f ${EX4_1_CONTAINER}
            docker network rm ${EX4_NETWORK_NAME}
            docker volume rm ${EX4_SHARED_VOL}
            ;;
        "shell")
            docker-compose -f ./deployments/ex4.yaml exec -it ex4_1 /bin/sh
            ;;
        "start")
            docker-compose -f ./deployments/ex4.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/ex4.yaml down
            ;;
        *)
            echo "$0 ex4 [command]

command:
    clean   stop network and clear container
    start   ex4 network
    stop    ex4 network"
            ;;
    esac
}

