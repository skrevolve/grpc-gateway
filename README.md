# go-grpc-api-gateway-microservice

## Configuration Description

1. API Gateway
    - [**api-gateway**](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/api-gateway)

2. gRPC Service
    - [**auth-svc**](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/auth-svc)
    - [**product-svc**](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/product-svc)
    - [**order-svc**](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/order-svc)

<!-- https://grpc-ecosystem.github.io/grpc-gateway/ -->

## Installing Protobuffer

### Linux (Ubuntu)

```sh
sudo apt install -y protobuf-compiler
```

### MacOS

```sh
brew install protobuff
```

### Windows

[**Releases Protocl Buffers v2x.x list**](https://github.com/protocolbuffers/protobuf/releases)

```sh
> protoc-2x.x-win32.zip
> protoc-2x.x-win64.zip
```

## GRPC and Protobuffer package dependencies

```sh
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/grpc
```

## Make Resource

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```

## 🐞Issues

### 1. [plugins are not supported : grpc #1070](https://github.com/golang/protobuf/issues/1070)

```sh
protoc protoc/route_guide.proto --go_out=plugins=grpc:.
ERR : --go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```

### 2. [Could not import Golang package : VScode](https://stackoverflow.com/questions/58518588/vscode-could-not-import-golang-package)

Use [**go work**](https://go.dev/doc/tutorial/workspaces) in project root

```sh
cd go-grpc-api-gateway-microservice
go work init
go work use ./api-gateway ./auth-svc ./order-svc ./product-svc
```
