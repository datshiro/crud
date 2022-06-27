#!/usr/bin/env bash


. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

echo_info "Test all packages"
go test -v -race ./...

# EXIT_CODE=$?
# cd "$WORKING_DIR" || exit 1
# exit $EXIT_CODE
