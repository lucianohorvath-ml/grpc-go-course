# gRPC Golang Demo

This is a repository containing demos from this [GRPC Golang course](http://bit.ly/grpc-golang-github).

<img src="https://grpc.io/img/logos/grpc-icon-color.png" width="150" alt="gRPC logo">

## Content

- Greeting Service
- Calculator Service
- Unary, Server Streaming, Client Streaming, BiDi Streaming
- Error Handling, Deadlines, SSL Encryption
- Blog API CRUD w/ MongoDB

## Requirements

- Go.
- Protocol Buffers compiler (`protoc`):
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
make run-server
```

If you modify a `.proto` file, you will need to recreate the generated files:
```
sh generate.sh
```
