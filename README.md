# gRPC Golang Demo

This is a repository containing demos from this [GRPC Golang course](http://bit.ly/grpc-golang-github).

Forked from https://github.com/simplesteph/grpc-go-course.

<img src="https://grpc.io/img/logos/grpc-icon-color.png" width="150" alt="gRPC logo">

## Content

- Greeting Service
- Calculator Service
- Unary, Server Streaming, Client Streaming, BiDi Streaming
- Error Handling, Deadlines, SSL Encryption
- Blog API CRUD w/ MongoDB

## Requirements

- Go.
- Protocol Buffers compiler (`protoc`).
  ```
  $ brew install protobuf
  ```
- Go plugins for `protoc`.
  ```
  $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
  $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
  $ export PATH="$PATH:$(go env GOPATH)/bin"
  ```

## Usage

### Greeting Service

1. Run the server.
```
make run-server
```

2. Run the client.
```
make run-client
```

### Calculator Service

1. Run the server.
```
make run-server-calc
```

2. Run the client.
```
make run-client-calc
```

This one has several examples of gRPC APIs:
- Unary RPC: Sum and SquareRoot
- Client streaming: ComputeAverage
- Server streaming: PrimeNumberDecomposition
- Bidirectional streaming: FindMaximum

### Notes

If you modify a `.proto` file, you will need to recreate the generated files:
```
sh generate.sh
```
