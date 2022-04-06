#!/bin/bash

#set -x

# Get last build success commit
url="https://git.new-page.xyz/api/v4/projects/${CI_PROJECT_ID}/pipelines?status=success&ref=master&private_token=${CI_PRIVATE_TOKEN}"
json=$(curl ${url})
exitCode=$?
if [[ $exitCode -ne 0 ]]; then
    echo "${json}"
    exit $exitCode
fi

last_commit=$(echo ${json} | python2.7 -c "import json; import sys; obj=json.load(sys.stdin); print obj[0]['sha'].encode('utf-8')")

# Get diff
files=$(git diff ${last_commit} HEAD --name-only --diff-filter=ACM | grep -E -i "\.(go|bazel)$")

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