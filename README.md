# go-test-server

![GitHub Actions Status](https://github.com/grimoh/go-test-server/actions/workflows/push.yaml/badge.svg)

## Usage

### local

```sh
go run cmd/server/main.go
```

#### Docker

```sh
docker build -t go-test-server .
docker run -p 8080:8080 go-test-server
```

#### Kubernetes

Prepare Docker Image.

```
docker build -t go-test-server .
```
Create K8s Cluster using [kind](https://github.com/kubernetes-sigs/kind) and load local Docker Image.

```sh
kind create cluster
kubectl cluster-info --context kind-kind
kind load docker-image go-test-server
```

```sh
kubectl apply -f k8s/raw
kubectl port-forward service/go-test-server 8080:8080 -n test-server
```
