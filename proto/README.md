#Blog Protos

##Preface
This repository contains all the protobuf files used by the GRPC service and gateway.

##Prerequisites
https://github.com/protocolbuffers/protobuf

## Using the proto files
To generate the Go code required by the service and gateway, you will need to compile the proto files.
###Compiling the proto files
Run:
```bash
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. \
  blog_api.proto
```
This will generate the code required for the GRPC service and the gateway.