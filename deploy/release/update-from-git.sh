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
go_bin="${NEMO_GO:-}"

find_go() {
	if [ -n "$go_bin" ]; then
		if [ -x "$go_bin" ]; then
			printf '%s\n' "$go_bin"
			return 0
		fi
		echo "NEMO_GO is set but not executable: $go_bin" >&2
		return 1
	fi
	if command -v go >/dev/null 2>&1; then
		command -v go
		return 0
	fi
	for candidate in /usr/local/go/bin/go /usr/bin/go /usr/lib/go/bin/go /snap/bin/go; do
		if [ -x "$candidate" ]; then
			printf '%s\n' "$candidate"
			return 0
		fi
	done
	return 1
}

check_go() {
	go_bin="$(find_go)" || {
		echo "go toolchain not found; install Go or set NEMO_GO=/absolute/path/to/go" >&2
		echo "Debian example: sudo apt install -y golang" >&2
		echo "Official tarball example: set NEMO_GO=/usr/local/go/bin/go after installing Go" >&2
		return 127
	}
	if [ ! -f go.mod ]; then
		echo "go.mod not found under $deploy_dir" >&2
		return 2
	fi
	required="$(awk '/^go / {print $2; exit}' go.mod)"
	actual="$("$go_bin" env GOVERSION 2>/dev/null || true)"
	actual="${actual#go}"
	if [ -z "$actual" ]; then
		actual="$("$go_bin" version | awk '{print $3}')"
		actual="${actual#go}"
	fi
	if [ -n "$required" ] && [ -n "$actual" ]; then
		oldest="$(printf '%s\n%s\n' "$required" "$actual" | sort -V | { IFS= read -r line; printf '%s' "$line"; })"
		if [ "$oldest" != "$required" ]; then
			echo "go version too old: found $actual, need $required from go.mod" >&2
			echo "set NEMO_GO to a newer Go binary or update the system Go install" >&2
			return 2
		fi
	fi
	echo "using go: $go_bin ($("$go_bin" version))"
}

ensure_deploy_script_modes() {
	for script in \
		deploy/release/build-local-current.sh \
		deploy/release/update-from-git.sh \
		deploy/release/update-from-github-release.sh \
		deploy/systemd/install-git-updater.sh \
		deploy/systemd/install-release-updater.sh \
		deploy/systemd/install-user-units.sh; do
		if [ -f "$script" ] && [ ! -x "$script" ]; then
			chmod +x "$script"
		fi
	done
}

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
ensure_deploy_script_modes

if [ "$before" = "$after" ] && [ "$force_build" != "true" ] && [ -x .bin/nemo ] && [ -x .bin/nemo-web ]; then
	echo "checkout unchanged at $after"
	exit 0
fi

check_go

if [ "$run_tests" = "true" ]; then
	"$go_bin" test ./...
fi

"$go_bin" build -trimpath -ldflags="-s -w" -o .bin/nemo ./cmd/nemo
"$go_bin" build -trimpath -ldflags="-s -w" -o .bin/nemo-web ./cmd/nemo-web

mkdir -p .deploy-state
printf '%s\n' "$after" > .deploy-state/current-git-commit
printf '%s/%s\n' "$remote" "$branch" > .deploy-state/current-git-source

if [ "$restart_after_update" = "true" ]; then
	systemctl --user daemon-reload || true
	systemctl --user restart "$restart_service"
fi

echo "deployed $remote/$branch at $after"
