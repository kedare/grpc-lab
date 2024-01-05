---
runme:
  id: 01HKCFQB6A8M685V3XX2Y0JQP9
  version: v2.0
---

# gRPC presentation

## What is gRPC ?

In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services. As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.

## gRPC

![image](grpc-architecture.svg)

## Protobuf
![image](protobuf.png)

[Protobuf file](proto/MonitoringService.proto)


## Components

### Gomonitor

Small gRPC server made in go that will expose CPU usage metrics of its host machine (either one shot or streamed)

### Blazinfra

Blazor SSR app that will connect to the gRPC server and display the CPU usage metrics to the client browsers

## Usage

### Install Go protobuf and grpc toolchain

```shell {"id":"01HKCFQB6A8M685V3XWPYMS8X1"}
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
```

### To generate the protobuf interfaces

```shell {"id":"01HKCFQB6A8M685V3XWRYBNPTP"}

# Generate Go interface
protoc -I proto/ --go_out gomonitor --go-grpc_out gomonitor proto/MonitoringService.proto

# Generate .NET interface
# Nothing to do, it's taken care automatically by grpc tool at build time
```

### How to use ?

```sh {"id":"01HKCFQB6A8M685V3XWT4NENAZ"}
# to build
make build
```

```sh {"id":"01HKCFQB6A8M685V3XWTWS05YR"}
# To start
make start

Then to go http://localhost:5000/
```

```sh {"id":"01HKCFQB6A8M685V3XWYMG85V3"}
# To stop
make stop
```

```sh {"id":"01HKCFQB6A8M685V3XX2B878GG"}
# to clean
make clean
```

### Improvements

- Go: Implement proper flow control and context management for the streaming to properly terminate the connections (instead of leaving gorouting forever)

## gRPC API gateway

See https://apisix.apache.org/blog/2021/12/30/apisix-proxy-grpc-service/
