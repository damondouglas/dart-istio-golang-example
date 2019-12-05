# About

This example combines [istio](https://istio.io), [grpc](https://grpc.io) with a simple [dart](https://dart.dev) front end.  This experiment stages future work using the same technologies for a [flutter](https://flutter.dev) application hosted on [kubernetes](https://kubernetes.io) using grpc backend services.

# Prerequisites

1. [docker](https://docker.com)
2. [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
3. [istio](https://istio.io/docs/setup/getting-started/)
4. [dart](https://dart.dev)
5. [go](https://golang.org/)

**IF YOU NEED HELP, PLEASE FEEL FREE TO CONTACT ME, THIS 💩 IS HARD SO DON'T FEEL LIKE YOU NEED TO BELONG TO SOME COOL CLUB TO ASK QUESTIONS**

# Setup

## Set docker to point to minikube

```
$ eval $(minikube docker-env)
```

## Build Echo service

The [echo service](server/server.go) simple reverses a string.  It simply does this to prove that a backend service is processing a payload provided by the frontend.

```
$ make build-server
```

## Build Echo Web

The [front end](client/dart/app) is built simply using dart from protobuf [generated dart client code](client/dart/lib/src/generated).

```
$ make build-web
```

## Kubernetes artifacts

```
kubetcl 
```

# Optional development tools for your interested but not required

## Proxy

To facilitate local testing and development, [proxy](proxy) simulates what istio does locally allowing the front end web to communicate with the backend service.

```
$ make build-proxy
$ make run-proxy
```

## Go Client

To quickly test whether the server is running, a local go client may be used:

```
$ go run client/go/client.go $(minikube ip):9090 asdf
```

