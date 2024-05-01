#!/bin/bash

if [ "$(basename $(realpath .))" != "learn-docker" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/ex1.sh
    . ./scripts/ex2.sh
fi

COMMAND=$1
SUBCOMMAND1=$2

case $COMMAND in
    "clean")
        ex1_ops clean
        ex2_ops clean        
        ;;
    "ex1")
        ex1_ops $SUBCOMMAND1
        ;;
    "ex2")
        ex2_ops $SUBCOMMAND1
        ;;
    "image")
        image $SUBCOMMAND1
        ;;
    *)
        echo "$0 [command]

command:
    clean  clear project of unused Docker artefacts
    ex1    working example to demonstrate non-root container
    ex2    demonstrate bridge networking example"
        ;;
esac