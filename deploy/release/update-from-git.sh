#!/usr/bin/env bash
set -euo pipefail

deploy_dir="${NEMO_DEPLOY_DIR:-$(pwd)}"
remote="${NEMO_DEPLOY_REMOTE:-origin}"
branch="${NEMO_DEPLOY_BRANCH:-main}"
run_tests="${NEMO_RUN_TESTS:-true}"
restart_service="${NEMO_RESTART_SERVICE:-nemo-web.service}"
restart_after_update="${NEMO_RESTART_AFTER_UPDATE:-true}"
force_build="${NEMO_FORCE_BUILD:-false}"
lock_path="${NEMO_DEPLOY_LOCK:-/tmp/nemo-git-update.lock}"

mkdir -p "$deploy_dir/.bin"
exec 9>"$lock_path"
if ! flock -n 9; then
	echo "another git update is already running"
	exit 0
fi

cd "$deploy_dir"
if [ ! -d .git ]; then
	echo "$deploy_dir is not a git checkout" >&2
	exit 2
fi

if ! git diff --quiet || ! git diff --cached --quiet; then
	echo "tracked working tree has local changes; refusing automatic git deploy" >&2
	git status --short
	exit 2
fi

before="$(git rev-parse HEAD)"
git fetch --prune "$remote" "$branch"

if git show-ref --verify --quiet "refs/heads/$branch"; then
	git switch "$branch"
else
	git switch -c "$branch" --track "$remote/$branch"
fi

git merge --ff-only "$remote/$branch"
after="$(git rev-parse HEAD)"

if [ "$before" = "$after" ] && [ "$force_build" != "true" ] && [ -x .bin/nemo ] && [ -x .bin/nemo-web ]; then
	echo "checkout unchanged at $after"
	exit 0
fi

if [ "$run_tests" = "true" ]; then
	go test ./...
fi

go build -trimpath -ldflags="-s -w" -o .bin/nemo ./cmd/nemo
go build -trimpath -ldflags="-s -w" -o .bin/nemo-web ./cmd/nemo-web

mkdir -p .deploy-state
printf '%s\n' "$after" > .deploy-state/current-git-commit
printf '%s/%s\n' "$remote" "$branch" > .deploy-state/current-git-source

if [ "$restart_after_update" = "true" ]; then
	systemctl --user daemon-reload || true
	systemctl --user restart "$restart_service"
fi

echo "deployed $remote/$branch at $after"
