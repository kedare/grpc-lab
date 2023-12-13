# gRPC presentation

# Setup

## Install Go protobuf and grpc toolchain

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
```

# To generate the protobuf interfaces

```shell

# Generate Go interface
protoc -I proto --go_out gomonitor --go-grpc_out gomonitor proto/MonitoringService.proto

# Generate .NET interface
```