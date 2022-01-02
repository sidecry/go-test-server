# go-test-server
[![GitHub Actions Status](https://github.com/grimoh/test-server/workflows/Go/badge.svg?branch=master)](https://github.com/grimoh/test-server/actions)

## Usage
### local

```sh
$ go run main.go
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
