#!/usr/bin/env bash

CUR_DIR=$(cd $(dirname $0); pwd)
SCRIPT_DIR="${CUR_DIR}/../script"

pushd "${CUR_DIR}" > /dev/null
    go run main.go -source "${SCRIPT_DIR}/contextual.go" > "${SCRIPT_DIR}/generated.go"
popd > /dev/null