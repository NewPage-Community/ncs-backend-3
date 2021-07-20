#!/bin/bash

# Get last build success commit
url="${CI_SERVER_URL}/api/v4/projects/${CI_PROJECT_ID}/pipelines?status=success&ref=master&private_token=${CI_PRIVATE_TOKEN}"
json=$(curl -s ${url})
if [[ "$json" != "[]" && "$json" != "" ]]; then
    last_commit=$(echo ${json} | python -c "import json; import sys; obj=json.load(sys.stdin); print obj[0]['sha'].encode('utf-8')")
else
    last_commit="HEAD^"
fi

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