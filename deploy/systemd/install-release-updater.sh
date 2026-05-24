#!/usr/bin/env bash
set -euo pipefail

deploy_dir="${NEMO_DEPLOY_DIR:-$(pwd)}"
repo="${NEMO_REPO:?set NEMO_REPO as owner/repo, for example karo/nemo-knows}"
release_tag="${NEMO_RELEASE_TAG:-main-latest}"
arch="${NEMO_ARCH:-}"
default_branch="${NEMO_DEFAULT_BRANCH:-main}"
interval="${NEMO_RELEASE_UPDATE_INTERVAL:-10min}"
restart_service="${NEMO_RESTART_SERVICE:-nemo-web.service}"

unit_dir="${XDG_CONFIG_HOME:-$HOME/.config}/systemd/user"
mkdir -p "$unit_dir"

cat > "$unit_dir/nemo-release-update.service" <<EOF
[Unit]
Description=nemo-knows release updater
After=network-online.target

[Service]
Type=oneshot
WorkingDirectory=$deploy_dir
Environment=NEMO_DEPLOY_DIR=$deploy_dir
Environment=NEMO_REPO=$repo
Environment=NEMO_RELEASE_TAG=$release_tag
Environment=NEMO_ARCH=$arch
Environment=NEMO_DEFAULT_BRANCH=$default_branch
Environment=NEMO_RESTART_SERVICE=$restart_service
ExecStart=$deploy_dir/deploy/release/update-from-github-release.sh
EOF

cat > "$unit_dir/nemo-release-update.timer" <<EOF
[Unit]
Description=Update nemo-knows from the GitHub release channel

[Timer]
OnBootSec=2min
OnUnitActiveSec=$interval
Persistent=true

[Install]
WantedBy=timers.target
EOF

systemctl --user daemon-reload
systemctl --user enable --now nemo-release-update.timer

echo "installed release updater under $unit_dir"
echo "updater: systemctl --user list-timers nemo-release-update.timer"
