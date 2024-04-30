#!/bin/bash

if [ "$(basename $(realpath .))" != "learn-docker" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/network.sh
fi

COMMAND=$1
SUBCOMMAND1=$2

export USER_NAME="john"

case $COMMAND in
    "image")
        image $SUBCOMMAND1
        ;;
    "nonroot")
        nonroot_op $SUBCOMMAND1
        ;;
    "rest")
        rest_network $SUBCOMMAND1
        ;;
    "clean")
        image        clean
        nonroot_op   clean
        rest_network stop
        ;;
    *)
        echo "$0 [command]

command:
    image    related operations - build and clean
    nonroot  executing container with nonroot user
    rest     network"
        ;;
esac