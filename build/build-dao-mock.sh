#!/bin/bash

find ${1} -name 'dao.go' -print0 |
    while IFS= read -r -d '' file; do
        path=${file%/*}
        package=${path##*/}
        filename=${file%.*}
        mockgen -source=${filename}.go -destination=${filename}.mock.go -package $package
    done