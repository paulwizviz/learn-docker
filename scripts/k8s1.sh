#!/bin/bash


COMMAND=$1

function message(){
    echo "Usage: $0 <Command>"
    echo
    echo "Command:"
    echo " cluster-clean    destroy cluster"
    echo " cluster-start    start cluster"
    echo " deploy       set deployment in minikube"
    echo " service      set service in minikube"
}

case $COMMAND in
    "cluster-clean")
        minikube stop
        minikube delete
        ;;
    "cluster-start")
        minikube start --driver=virtualbox
        ;;
    "deploy")
        kubectl apply -f ./deployment/k8s1/deployment.yaml
        ;;
    "service")
        kubectl apply -f ./deployment/k8s1/service.yaml
        ;;
    *)
        message
        exit 1
        ;;
esac