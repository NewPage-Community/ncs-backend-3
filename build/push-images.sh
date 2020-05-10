#!/bin/bash

function push_changed() {
    pkgs=$(build/changed-pkgs.sh)
    exitCode=$?
    if [[ ${exitCode} -ne 0 ]]; then
        echo "build/get-pkgs.sh fail"
        exit ${exitCode}
    fi
    if [[ "${pkgs}" = "" ]]; then
        echo "No changed packages"
        exit 0
    fi

    echo -e "Packages for pushing:\n${pkgs}\n"

    paths=""
    for pkg in ${pkgs}; do
        if [[ "${paths}" != "" ]]; then
            paths="${paths} union allpaths(//app/..., //${pkg}:all)"
        else
            paths="allpaths(//app/..., //${pkg}:all)"
        fi
    done

    echo "Pushing image to docker..."
    query=$(bazel query "${paths}")
    bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 -k $(echo -e "${query}" | grep push)
}

function push_all() {
    echo "Pushing all image to docker..."
    query=$(bazel query //app/...)
    bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 -k $(echo -e "${query}" | grep push)
}

if [[ "${1}" = "changed" ]]; then
    push_changed
else
    push_all
fi
