#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Injects the version and builds the Go app with watch.
#
# Examples
#
#   ./scripts/dev.sh
#
# Returns exit code 0.
function main() {
  local arch
  local build_path
  local os
  local version

  set_vars

  arch=$(go env GOARCH)
  os=$(go env GOOS)
  build_path="${os}-${arch}"

  # get the version in the version file
  version=$(<VERSION)

  printf "%b starting app...\n" "${INFO_PREFIX}"
  CompileDaemon \
    -build="go build -ldflags=-X=main.Version=${version} -o ${PWD}${BIN_DIR}/${build_path}/koinonos ./cmd/koinonos/main.go" \
    -command="${PWD}${BIN_DIR}/${build_path}/koinonos"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$@"
