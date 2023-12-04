#!/usr/bin/env bash

if [[ "${TRACE-0}" == "1" ]]; then set -x; fi
set -Eeuo pipefail
IFS=$'\n\t'
pushd "$(dirname "$0")" >/dev/null

# main
{
    go build
}

popd >/dev/null
