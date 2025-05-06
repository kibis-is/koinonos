#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Creates a .env file, if it don't exist.
#
# Examples
#
#   ./scripts/setup.sh
#
# Returns exit code 0.
function main() {
  local env_example_file

  set_vars

  env_example_file=".env.example"

  if [[ ! -f "${env_example_file}" ]]; then
    printf "%b no .env.example at %b, ignoring \n" "${INFO_PREFIX}" "${env_example_file}"
    exit 0
  fi

  printf "%b creating file: .env \n" "${INFO_PREFIX}"
  # copies the .env.example if it doesn't exist (uses no clobber [-n])
  cp -n ".env.example" ".env"

  exit 0
}

# and so, it begins...
main
