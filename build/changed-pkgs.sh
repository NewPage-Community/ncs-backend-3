#!/bin/bash

files=$(git diff HEAD^ HEAD --name-only --diff-filter=ACM | grep -E -i "\.(go|bazel)$")

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