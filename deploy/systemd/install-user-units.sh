#!/usr/bin/env bash
set -euo pipefail

# Migration note: if you previously ran `nemo-wiki-maintain.timer`, disable it once:
#   systemctl --user disable --now nemo-wiki-maintain.{service,timer}
#   rm -f "$unit_dir"/nemo-wiki-maintain.{service,timer}

deploy_dir="${NEMO_DEPLOY_DIR:-$(pwd)}"
web_addr="${NEMO_WEB_ADDR:-127.0.0.1:8787}"
maintain_mode="${NEMO_MAINTAIN_MODE:-auto}"
maintain_provider="${NEMO_MAINTAIN_PROVIDER:-deepseek}"
maintain_profile="${NEMO_MAINTAIN_PROFILE:-stable}"
maintain_interval="${NEMO_MAINTAIN_INTERVAL:-1h}"

unit_dir="${XDG_CONFIG_HOME:-$HOME/.config}/systemd/user"
mkdir -p "$unit_dir"

cat > "$unit_dir/nemo-web.service" <<EOF
[Unit]
Description=nemo-knows web console
After=network-online.target

[Service]
WorkingDirectory=$deploy_dir
ExecStart=$deploy_dir/.bin/nemo-web -addr $web_addr
Restart=on-failure
RestartSec=5

[Install]
WantedBy=default.target
EOF

cat > "$unit_dir/nemocli-once.service" <<EOF
[Unit]
Description=nemo-knows nemocli once trigger
After=network-online.target

[Service]
Type=oneshot
WorkingDirectory=$deploy_dir
Environment=NEMO_ONCE_MODE=$maintain_mode
ExecStart=/usr/bin/flock -n /tmp/nemo-ingest.lock $deploy_dir/.bin/nemocli --provider "$maintain_provider" --profile "$maintain_profile" once
EOF

cat > "$unit_dir/nemocli-once.timer" <<EOF
[Unit]
Description=Run nemocli once periodically

[Timer]
OnBootSec=5min
OnUnitActiveSec=$maintain_interval
Persistent=true

[Install]
WantedBy=timers.target
EOF

systemctl --user daemon-reload
systemctl --user enable --now nemo-web.service
systemctl --user enable --now nemocli-once.timer

echo "installed user units under $unit_dir"
echo "web:   systemctl --user status nemo-web.service"
echo "once:  systemctl --user list-timers nemocli-once.timer"
