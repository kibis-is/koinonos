#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Injects the version and runs the Go app.
#
# $1 - [optional] a version to inject, otherwise the version from the VERSION file is read.
#
# Examples
#
#   ./scripts/dev.sh # reads the version in the VERSION file
#   ./scripts/dev.sh "1.2.3"
#
# Returns exit code 0.
function main() {
  local version

  set_vars

  # get the version in the version file
  version=$(<VERSION)

  # if the version argument exists, use it instead of the one on file
  if [ -n "$1" ]; then
    version="$1"
  fi

  printf "%b starting app...\n" "${INFO_PREFIX}"
  CompileDaemon -build="go build -ldflags=-X=main.Version=${version} -o ${BUILD_DIR}/koinonos ./cmd/koinonos/main.go" -command="${BUILD_DIR}/koinonos"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$@"
