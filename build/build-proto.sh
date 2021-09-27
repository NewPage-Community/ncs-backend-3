#!/bin/bash

# Install tools
apt update
apt install -y protobuf-compiler

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/{protoc-gen-grpc-gateway,protoc-gen-openapiv2} \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/golang/mock/mockgen

export PATH=$PATH:/$GOPATH/bin

find ${1} -name '*.proto' -print0 |
    while IFS= read -r -d '' file; do
        path=${file%/*}
        package=${path##*/}
        filename=${file%.*}
        protoc \
            --go_out=paths=source_relative:. \
            --go-grpc_out=paths=source_relative:. \
            --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
            --openapiv2_out=logtostderr=true:. \
            $file
        mockgen -source=${filename}_grpc.pb.go -destination=${filename}_grpc.pb.mock.go -package $package
    done