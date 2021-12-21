#!/bin/bash

COMMAND=$1

. ./scripts/common.sh

case $COMMAND in
    "start")
        docker-compose -f ./deployment/compose1/docker-compose.yaml up -d
        ;;
    "stop")
        docker-compose -f ./deployment/compose1/docker-compose.yaml down
        ;;
    *)
        echo "Usage; $0 <Command>"
        echo
        echo "Command:"
        echo " start   docker-compose test"
        echo " stop    docker-compose test"
        ;;
esac