#!/bin/bash

find ${1} -name '*.proto' -print0 |
    while IFS= read -r -d '' file; do
        path=${file%/*}
        package=${path##*/}
        filename=${file%.*}
        protoc -I. --gofast_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. $file
        mockgen -source=${filename}.pb.go -destination=${filename}.pb.mock.go -package $package
    done