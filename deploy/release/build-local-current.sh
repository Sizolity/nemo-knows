#!/usr/bin/env bash
set -euo pipefail

deploy_dir="${NEMO_DEPLOY_DIR:-$(pwd)}"
ref="${NEMO_DEPLOY_REF:-}"
run_tests="${NEMO_RUN_TESTS:-true}"
allow_dirty="${NEMO_ALLOW_DIRTY:-false}"
restart_service="${NEMO_RESTART_SERVICE:-nemo-web.service}"
restart_after_build="${NEMO_RESTART_AFTER_BUILD:-true}"

cd "$deploy_dir"

if [ -n "$ref" ]; then
	git fetch origin "$ref"
	git switch --detach FETCH_HEAD
fi

if [ "$allow_dirty" != "true" ]; then
	if ! git diff --quiet || ! git diff --cached --quiet; then
		echo "tracked working tree has local changes; set NEMO_ALLOW_DIRTY=true to build anyway" >&2
		exit 2
	fi
fi

if [ "$run_tests" = "true" ]; then
	go test ./...
fi

mkdir -p .bin
go build -trimpath -ldflags="-s -w" -o .bin/nemo ./cmd/nemo
go build -trimpath -ldflags="-s -w" -o .bin/nemo-web ./cmd/nemo-web

if [ "$restart_after_build" = "true" ]; then
	systemctl --user daemon-reload || true
	systemctl --user restart "$restart_service"
fi

echo "built and deployed current checkout"
