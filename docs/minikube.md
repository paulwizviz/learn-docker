# Using minikube and kubectl

This is a an implementation using minikube based on this example [https://minikube.sigs.k8s.io/docs/start/](https://minikube.sigs.k8s.io/docs/start/)

## Minikube operations

1. Start minikube.
   * With default vm `minikube start`
   * With virtual box `minikube start --driver=virtualbox`
1. Opening dashboard. `minikube dashboard`
1. Addons. 
   * List addons. `minikube addons list`
   * Enabling addons. `minikube addons enable <addons name>`
   * Disabline addons. `minikube addons disable <addons name>`
1. VM operations.
   * Stop vm. `minikube stop`
   * Delete vm. `minikube delete`

## Using Kubectl to create a basic application

In this scenario, we are creating an app based on an image named `k8s.gcr.io/echoserver:1.4`. The steps involved are:

1. Create deployment call **hello-node** based on image k8s.gcr.io/echoserver:1.4. `kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4`
1. Create a service Expose deployment named **hello-node**, exposing port 8080 to the outside world. `kubectl expose deployment hello-node --type=LoadBalancer --port=8080`
1. Provide access to an applications. `kubectl port-forward service/hello-node 7080:8080`

## Inspecting and managing applications

Deployment related
1. View deployment. `kubectl get deployments`
1. View pod. `kubectl get pods`
1. Delete deployment. `kubectl delete deployment hello-node`

Service related
1. Delete services. `kubectl delete service hello-node`
1. View service. `kubectl get services`