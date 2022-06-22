#!/usr/bin/env bash

. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

if [[ -z "${DB_URL}" ]]; then
    echo_info "Not set DB_URL yet"
else
    echo_info "Using db url: $DB_URL"
fi

# Go to project root dir to make sure that we can call other scripts correctly
cd "$PROJECT_DIR"

cmd=$*
shift
echo_info "Run migrate command: $cmd $options"
./bin/migrate -verbose -database "$DB_URL" -path ./db/migrations/ $cmd 
cd $WORKING_DIR
