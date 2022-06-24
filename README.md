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

## Disclaimer

* The packages in this project are intended for educational purpose only.
* This package is constantly updated and items may be removed and modified without warning.

## Copyright

Unless otherwise specificed, the copyright in this project are assigned as follows.

Copyright 2022 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.