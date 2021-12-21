# Overview

A collection of examples demonstrating techniques to:

* build containers;
* orchestrate docker and kubernetes networks

## Project structure

* `build` folder contains scripts to build executable examples
* `cmd` folder contains Go codes for building big executable applications
* `deployment` folder contains scripts to deploy on docker-compose, k8s or any other container orchestration
* `examples` folder contains single `main` to demonstrate small example
* `internal` folder contains library codes
* `scripts` folder contains bash scripts to help you build and run executable examples

## Topics

* [Using Docker for CI/CD](./docs/cicd.md)
* [Interacting with Docker deamon](./docs/dockerops.md)
* [Local docker compose deployment](./docs/compose1.md)
* [Minimal k8s](./docs/k8s1.md)
* [MiniKube](./docs/minikube.md)

## References

* [Kubernetes Components](https://kubernetes.io/docs/concepts/overview/components/)
* [The Kubernetes API](https://kubernetes.io/docs/concepts/overview/kubernetes-api/)
* [Working with Kubernetes Objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/)
* [Cluster Architecture](https://kubernetes.io/docs/concepts/architecture/)
