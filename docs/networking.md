# Networking

In this section, we examine the operations behind containers networking.

## Multiple Bridge Networking Setup

Here is a working example, [ex3](../deployments/ex3.yaml), demonstrating a setup with multiple bridges. One bridge. The first bridge has an IP address range from 172.28.5.0/24 and the second range from 172.30.5.0/24.

If you shell into `ex3_1` and `curl ex3_2:9090` this will return a message `hello`. It means you can access container `ex3_2` via name from `ex3_1`. 

If you curl from `ex3_1` to `ex3_3`, which is in a different network, it will say host name cannot be resolved.

To shell into `ex3_1`, use the [script](../scripts/ops.sh) run the command `./scripts/ops.sh ex3 shell`.

## Useful References

* [Official documentation](https://docs.docker.com/engine/tutorials/networkingcontainers/)
* [Docker Networking Tutorial (Bridge - None - Host - IPvlan - Macvlan - Overlay)](https://www.youtube.com/watch?v=fBRgw5dyBd4)