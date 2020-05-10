#!/bin/bash

changed=${1}
files=""

function changed_files() {
    files=$(git diff HEAD^ HEAD --name-only --diff-filter=ACM | grep -E -i ".go")
}

function all_files() {
    files=$(find ./app -type f -name 'BUILD.bazel' | xargs grep -l 'push_docker')
    files=${files#./}
}

if [[ "${changed}" = "changed" ]]; then
    changed_files
else
    all_files
fi

pkgs=""
for file in ${files}; do
    pkg="${file%/*}"
    if [ $? -eq 0 ]; then
        if [[ "${pkgs}" = "" ]]; then
            pkgs="${pkg}"
        else
            pkgs="${pkgs}\n${pkg}"
        fi
    fi
done
echo -e "${pkgs}" >/tmp/.pkgs
pkgs=$(sort /tmp/.pkgs | uniq)
echo -e "${pkgs}"