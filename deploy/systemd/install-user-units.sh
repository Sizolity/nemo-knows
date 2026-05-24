#!/usr/bin/env bash
set -euo pipefail

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

cat > "$unit_dir/nemo-wiki-maintain.service" <<EOF
[Unit]
Description=nemo-knows wiki maintainer
After=network-online.target

[Service]
Type=oneshot
WorkingDirectory=$deploy_dir
ExecStart=/usr/bin/flock -n /tmp/nemo-wiki-maintain.lock /usr/bin/env bash -lc 'cd "$deploy_dir"; run_id=\$(date +%%Y%%m%%d-%%H%%M%%S); .bin/nemo -provider "$maintain_provider" -profile "$maintain_profile" -maintain-wiki -mode "$maintain_mode" -out-dir ".wiki-maintain/\$run_id"'
EOF

cat > "$unit_dir/nemo-wiki-maintain.timer" <<EOF
[Unit]
Description=Run nemo-knows wiki maintainer periodically

[Timer]
OnBootSec=5min
OnUnitActiveSec=$maintain_interval
Persistent=true

[Install]
WantedBy=timers.target
EOF

systemctl --user daemon-reload
systemctl --user enable --now nemo-web.service
systemctl --user enable --now nemo-wiki-maintain.timer

echo "installed user units under $unit_dir"
echo "web:       systemctl --user status nemo-web.service"
echo "maintain:  systemctl --user list-timers nemo-wiki-maintain.timer"
