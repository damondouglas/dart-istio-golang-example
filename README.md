# About

This example combines [istio](https://istio.io), [grpc](https://grpc.io) with a simple [dart](https://dart.dev) front end.  This experiment stages future work using the same technologies for a [flutter](https://flutter.dev) application using gRPC backend services hosted on [kubernetes](https://kubernetes.io).

# Prerequisites

This README assumes the following are setup on your machine.

1. [docker](https://docker.com)
2. [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
3. [istio](https://istio.io/docs/setup/getting-started/)
4. [dart](https://dart.dev)
5. [go](https://golang.org/)

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

## Build kubernetes artifacts

```
$ kubetcl apply -f manifest/gateway.yaml
$ kubetcl apply -f manifest/server.yaml
```

## Acquire minikube gateway address

See [determining the ingress IP and Ports](https://istio.io/docs/tasks/traffic-management/ingress/ingress-control/#determining-the-ingress-ip-and-ports) from the Istio documentation for more information.

```
$ export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
$ export GATEWAY=$(minikube ip):$INGRESS_PORT
```

## Launch web server

Run the following commands to open the web front end.  Filling out anything in the input and pressing Submit should show a string reversal proving that you've set up everything correctly.  Your front end web app is calling your backend gRPC.

```
$ docker run -it -p 5000:5000 -e PORT=5000 -e ENDPOINT=http://$GATEWAY echo-web:1.0
$ open "http://$(minikube ip):5000"
```
