# Learn how to use proto

## Define a proto target in a `.proto` file
// TODO

## Gen golang code from `.proto` file
1. install protoc and protoc-gen-go
```cmd
scoop install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
2. run the command
```cmd
protoc --proto_path=proto --go_out=app proto/*.proto
```

## Build a client and a server transport data via protobuf
// TODO


# Reference
- [how to gen go code from proto](https://developers.google.com/protocol-buffers/docs/reference/go-generated#package)
- [The import path must contain at least one period ('.') or forward slash ('/') character](https://github.com/techschool/pcbook-go/issues/3)