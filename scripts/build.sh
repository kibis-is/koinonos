#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Builds the Go app with the version fro the VERSION file injected and packages the binary.
#
# See https://go.dev/doc/install/source#environment for a list of supported OS/arch combos.
#
# $1 - [optional] The OS and architecture to build to. They must be formatted in as "<OS>-<ARCH>".
#
# Examples
#
#   ./scripts/build.sh # uses the default set at `go env GOOS` & `go env GOARCH`
#   ./scripts/build.sh "darwin-amd64"
#   ./scripts/build.sh "darwin-arm64"
#   ./scripts/build.sh "linux-amd64"
#   ./scripts/build.sh "linux-arm64"
#   ./scripts/build.sh "windows-amd64"
#   ./scripts/build.sh "windows-arm64"
#
# Returns exit code 0.
function main() {
  local arch
  local build_path
  local os
  local version

  set_vars

  # get the version in the version file
  version=$(<VERSION)

  if [[ -n "${1}" ]]; then
    IFS='-' read -r os arch <<< "${1}"
  fi

  if [[ -z "${os}" ]]; then
    os=$(go env GOOS)
  fi

  if [[ -z "${arch}" ]]; then
    arch=$(go env GOARCH)
  fi

  build_path="${os}-${arch}"

  go build \
    -ldflags="-X main.Version=${version}" \
    -o "${PWD}${BIN_DIR}/${build_path}/koinonos" \
    ./cmd/koinonos/main.go

  printf "%b build at \"%b\"\n" "${INFO_PREFIX}" "${PWD}/${BIN_DIR}/${build_path}/koinonos"

  # create the dist directory if it does not exist
  if [ ! -d "${DIST_DIR}" ]; then
    printf "%b no \"%b\" directory found, creating a new one \n" "${INFO_PREFIX}" "${DIST_DIR}"

    mkdir -p "${DIST_DIR}"
  fi

  tar -czf "${PWD}/${DIST_DIR}/koinonos-${os}-${arch}-${version}.tar.gz" -C "${PWD}/${BIN_DIR}/${build_path}" koinonos

  printf "%b package at \"%b\" \n" "${INFO_PREFIX}" "${PWD}/${DIST_DIR}/koinonos-${os}-${arch}-${version}.tar.gz"

  exit 0
}

# and so, it begins...
main "$@"
