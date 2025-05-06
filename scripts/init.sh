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

  # install commitlint
  go install github.com/conventionalcommit/commitlint@latest

  printf "%b setting git hooks \n" "${INFO_PREFIX}"
  git config core.hooksPath "${PWD}/.commitlint/hooks"

  env_example_file=".env.example"

  if [[ -f "${env_example_file}" ]]; then
    printf "%b creating file: .env \n" "${INFO_PREFIX}"

    # copies the .env.example (uses no clobber [-n] so it doesn't overwrite any existing .env)
    cp -n ".env.example" ".env"
  fi

  exit 0
}

# and so, it begins...
main
