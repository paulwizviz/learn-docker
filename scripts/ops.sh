#!/bin/bash

if [ "$(basename $(realpath .))" != "go-docker" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/ex1.sh
    . ./scripts/ex2.sh
    . ./scripts/ex3.sh
    . ./scripts/ex4.sh
    . ./scripts/ex5.sh
fi

COMMAND=$1
SUBCOMMAND1=$2

case $COMMAND in
    "clean")
        ex1_ops clean
        ex2_ops clean
        ex3_ops clean
        ex4_ops clean
        image clean        
        ;;
    "ex1")
        ex1_ops $SUBCOMMAND1
        ;;
    "ex2")
        ex2_ops $SUBCOMMAND1
        ;;
    "ex3")
        ex3_ops $SUBCOMMAND1
        ;;
    "ex4")
        ex4_ops $SUBCOMMAND1
        ;;
    "ex5")
        ex5_ops $SUBCOMMAND1
        ;;
    "image")
        image $SUBCOMMAND1
        ;;
    *)
        echo "$0 [command]

command:
    clean  clear project of unused Docker artefacts
    ex1    working example to demonstrate non-root container
    ex2    demonstrate container with REST server
    ex3    demonstrate multi bridge networking example
    ex4    demonstrate logging example
    image  operations to build and clean images"
        ;;
esac