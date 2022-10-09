# external-server

## Usage

### local

```sh
go run cmd/server/main.go
```

#### Docker

```sh
docker build -t external-server .
docker run -p 8080:8080 external-server
```

#### Kubernetes

Prepare Docker Image.

```
docker build -t external-server .
```

Create K8s Cluster using [kind](https://github.com/kubernetes-sigs/kind) and load local Docker Image.

```sh
kind create cluster
kubectl cluster-info --context kind-kind
kind load docker-image external-server
```

##### Kubernetes Manifest(Raw)

Apply K8s manifest and connect to K8s service.

```sh
kubectl apply -f kubernetes/raw
kubectl port-forward service/external-server 8080:8080 -n external-server
```
##### Kubernetes Manifest(HelmChart)

##### Kubernetes Manifest(Helmfile)
