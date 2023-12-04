#!/usr/bin/env bash

if [[ "${TRACE-0}" == "1" ]]; then set -x; fi
set -Eeuo pipefail
IFS=$'\n\t'
pushd "$(dirname "$0")" >/dev/null

function integration_tests() {
    set +e # disable errexit
    for test_file in $(find test -name '.run'); do
        root="$(dirname "$test_file")"
        echo "running $test_file"

        if [[ -f "$root"/.in ]]; then
            ECHO=1 bash "$test_file" \
                <"$root"/.in \
                1>"$root"/.out \
                2>"$root"/.log &
        else
            ECHO=1 bash "$test_file" \
                1>"$root"/.out \
                2>"$root"/.log &
        fi
    done
    set -e
}

# main
{
    if [[ "${1:-}" == "clean" ]]; then
        find test \( -name ".log" -o -name ".out" \) | xargs rm
    else
        go test ./... -v
        integration_tests
    fi
}

popd >/dev/null
