#!/bin/bash

find ${1} -name '*.proto' -print0 |
    while IFS= read -r -d '' file; do
        path=${file%/*}
        package=${path##*/}
        filename=${file%.*}
        protoc -I. --go_out=. --go-grpc_out=require_unimplemented_servers=false:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. $file
        mockgen -source=${filename}_grpc.pb.go -destination=${filename}_grpc.pb.mock.go -package $package
    done