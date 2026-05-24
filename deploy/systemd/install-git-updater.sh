#!/usr/bin/env bash
set -euo pipefail

deploy_dir="${NEMO_DEPLOY_DIR:-$(pwd)}"
remote="${NEMO_DEPLOY_REMOTE:-origin}"
branch="${NEMO_DEPLOY_BRANCH:-main}"
interval="${NEMO_GIT_UPDATE_INTERVAL:-10min}"
run_tests="${NEMO_RUN_TESTS:-true}"
restart_service="${NEMO_RESTART_SERVICE:-nemo-web.service}"

unit_dir="${XDG_CONFIG_HOME:-$HOME/.config}/systemd/user"
mkdir -p "$unit_dir"

cat > "$unit_dir/nemo-git-update.service" <<EOF
[Unit]
Description=nemo-knows git updater
After=network-online.target

[Service]
Type=oneshot
WorkingDirectory=$deploy_dir
Environment=NEMO_DEPLOY_DIR=$deploy_dir
Environment=NEMO_DEPLOY_REMOTE=$remote
Environment=NEMO_DEPLOY_BRANCH=$branch
Environment=NEMO_RUN_TESTS=$run_tests
Environment=NEMO_RESTART_SERVICE=$restart_service
ExecStart=$deploy_dir/deploy/release/update-from-git.sh
EOF

cat > "$unit_dir/nemo-git-update.timer" <<EOF
[Unit]
Description=Update nemo-knows from git over SSH

[Timer]
OnBootSec=2min
OnUnitActiveSec=$interval
Persistent=true

[Install]
WantedBy=timers.target
EOF

systemctl --user daemon-reload
systemctl --user enable --now nemo-git-update.timer

echo "installed git updater under $unit_dir"
echo "updater: systemctl --user list-timers nemo-git-update.timer"
