# go-test-server

![GitHub Actions Status](https://github.com/grimoh/go-test-server/actions/workflows/push.yaml/badge.svg)

## Usage
### local

```sh
$ go run cmd/server/main.go
```

### Docker

```sh
$ docker build -t go-test-server .
$ docker run -p 8080:8080 go-test-server
```

### Kubernetes

```sh
$ kubectl apply -f kubernetes/raw/
```
