# Interacting with Docker

Docker deployment comsists of a client, a deamon, and a registry.

![Docker architecture](../assets/img/docker-arch.svg)
Source: [Docker overview](https://docs.docker.com/get-started/overview/)

You use the client to interact with the deamon. There are three types of clients:

* A command line tool.
* Using Unix socket.
* Using an Software Development Kit (SDK).

The deamon provides a (RESTful) API. The API specification is [here](https://docs.docker.com/engine/api/v1.41/).

## Interacting via Unix socket

