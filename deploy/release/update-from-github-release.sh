#!/usr/bin/env bash
set -euo pipefail

deploy_dir="${NEMO_DEPLOY_DIR:-$(pwd)}"
repo="${NEMO_REPO:?set NEMO_REPO as owner/repo, for example karo/nemo-knows}"
release_tag="${NEMO_RELEASE_TAG:-main-latest}"
arch="${NEMO_ARCH:-}"
github_base="${NEMO_GITHUB_BASE_URL:-https://github.com}"
default_branch="${NEMO_DEFAULT_BRANCH:-main}"
update_default_branch="${NEMO_UPDATE_DEFAULT_BRANCH:-true}"
restart_service="${NEMO_RESTART_SERVICE:-nemo-web.service}"
restart_after_update="${NEMO_RESTART_AFTER_UPDATE:-true}"
lock_path="${NEMO_DEPLOY_LOCK:-/tmp/nemo-release-update.lock}"

if [ -z "$arch" ]; then
	case "$(uname -m)" in
		x86_64|amd64) arch="amd64" ;;
		aarch64|arm64) arch="arm64" ;;
		*) echo "unsupported architecture: $(uname -m); set NEMO_ARCH manually" >&2; exit 2 ;;
	esac
fi

package="nemo-knows-linux-${arch}.tar.gz"
state_dir="$deploy_dir/.deploy-state"
releases_dir="$deploy_dir/.releases"

mkdir -p "$state_dir" "$releases_dir" "$deploy_dir/.bin"
exec 9>"$lock_path"
if ! flock -n 9; then
	echo "another release update is already running"
	exit 0
fi

cd "$deploy_dir"

if [ "$update_default_branch" = "true" ] && [ -d .git ]; then
	echo "checking default branch update: origin/$default_branch"
	git fetch --quiet origin "$default_branch" || echo "warning: failed to fetch origin/$default_branch"
	if git diff --quiet && git diff --cached --quiet; then
		git merge --ff-only "origin/$default_branch" || echo "warning: default branch fast-forward skipped"
	else
		echo "tracked working tree has local changes; skipping default branch fast-forward"
	fi
fi

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

release_url="$github_base/$repo/releases/download/$release_tag"
echo "downloading $release_url/$package"
curl -fsSL -o "$tmp/$package" "$release_url/$package"
curl -fsSL -o "$tmp/checksums.txt" "$release_url/checksums.txt"

(
	cd "$tmp"
	grep "  $package\$" checksums.txt | sha256sum -c -
)

checksum="$(sha256sum "$tmp/$package" | awk '{print $1}')"
state_file="$state_dir/${release_tag}-${arch}.sha256"
if [ -f "$state_file" ] && [ "$(cat "$state_file")" = "$checksum" ]; then
	echo "release package unchanged: $release_tag/$package"
	exit 0
fi

extract_dir="$tmp/extract"
mkdir -p "$extract_dir"
tar -xzf "$tmp/$package" -C "$extract_dir"
release_dir="$(find "$extract_dir" -mindepth 1 -maxdepth 1 -type d | sort | head -n 1)"
if [ -z "$release_dir" ]; then
	echo "release archive did not contain a top-level directory" >&2
	exit 1
fi

install -m 0755 "$release_dir/nemo" "$deploy_dir/.bin/nemo"
install -m 0755 "$release_dir/nemo-web" "$deploy_dir/.bin/nemo-web"

printf '%s\n' "$checksum" > "$state_file"
printf '%s\n' "$release_tag" > "$state_dir/current-release-tag"
printf '%s\n' "$package" > "$state_dir/current-release-package"
cp "$tmp/$package" "$releases_dir/$package"

if [ "$restart_after_update" = "true" ]; then
	systemctl --user daemon-reload || true
	systemctl --user restart "$restart_service"
fi

echo "deployed $release_tag/$package"
