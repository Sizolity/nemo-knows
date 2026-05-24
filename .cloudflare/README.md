# nemo-knows Cloudflare 部署

通过 Cloudflare Tunnel 暴露本地 nemo-web，Worker 做边缘缓存代理。

## 架构

```
用户 → Cloudflare Worker (缓存) → Cloudflare Tunnel → nemo-web (localhost:8787)
```

## 1. 服务器端：启动 nemo-web

```bash
cd /path/to/nemo-knows
go build -o .bin/nemo-web ./cmd/nemo-web
.bin/nemo-web -addr 127.0.0.1:8787
```

## 2. 服务器端：安装并配置 cloudflared

```bash
# 安装（Debian/Ubuntu，二选一）

# 方法 a: 从 GitHub releases 下载 .deb
wget -q https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64.deb
sudo dpkg -i cloudflared-linux-amd64.deb

# 方法 b: 直接下载二进制
# curl -fsSL https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64 -o /usr/local/bin/cloudflared
# chmod +x /usr/local/bin/cloudflared

# 登录（浏览器授权）
cloudflared tunnel login

# 创建隧道
cloudflared tunnel create nemo

# 配置隧道
mkdir -p ~/.cloudflared
cat > ~/.cloudflared/config.yml << 'EOF'
tunnel: nemo
credentials-file: /home/<user>/.cloudflared/<tunnel-id>.json

ingress:
  - hostname: nemo.yourdomain.com
    service: http://localhost:8787
  - service: http_status:404
EOF

# 添加 DNS 记录（让 nemo.yourdomain.com 指向隧道）
cloudflared tunnel route dns nemo nemo.yourdomain.com

# 启动隧道
cloudflared tunnel run nemo
```

用 systemd 持久化：

```bash
sudo cloudflared service install
sudo systemctl enable --now cloudflared
```

## 3. 部署 Worker

```bash
cd .cloudflare/worker

# 安装依赖
npm install

# 修改 wrangler.toml 中的 NEMO_ORIGIN 为隧道域名
# NEMO_ORIGIN = "https://nemo.yourdomain.com"

# 本地测试
npm run dev

# 部署到 Cloudflare
npm run deploy
```

## 4. 服务器自动部署

仓库包含 GitHub Actions CD workflow（`.github/workflows/cd.yml`）：

- push 到 `main` 或手动 `workflow_dispatch` 时运行。
- 先执行 `go test ./...`。
- 为服务器构建 Linux 二进制产物：
  - `nemo-knows-linux-amd64.tar.gz`
  - `nemo-knows-linux-arm64.tar.gz`
- 每个产物包含 `nemo`、`nemo-web`、`README.md`、`AGENTS.md` 和 `deploy/`。
- 默认分支的成功构建会发布到固定 GitHub Release 通道 `main-latest`，
  并附带 `checksums.txt`。

公开仓库不应默认由 push 触发服务器远程执行。这里的安全边界是：
GitHub 只托管代码和构建结果；服务器由本机 timer 主动拉取和部署。
不要把服务器 SSH 私钥放进公开仓库的自动 push 部署链路里。

### 默认路线：Git over SSH 自动更新

```bash
# 服务器上的 origin 应该是 git@github.com:Sizolity/nemo-knows.git
git remote set-url origin git@github.com:Sizolity/nemo-knows.git
ssh -T git@github.com

NEMO_DEPLOY_DIR=/path/to/nemo-knows \
NEMO_DEPLOY_REMOTE=origin \
NEMO_DEPLOY_BRANCH=main \
NEMO_RUN_TESTS=true \
NEMO_GIT_UPDATE_INTERVAL=10min \
./deploy/systemd/install-git-updater.sh
```

`nemo-git-update.timer` 会定时执行：

1. `git fetch --prune origin main`。
2. 拒绝在 tracked working tree 有本地改动时自动部署。
3. 快进到 `origin/main`。
4. 运行 `go test ./...`。
5. 本地构建 `.bin/nemo` 和 `.bin/nemo-web`。
6. 重启 `nemo-web.service`。

这条路线适合当前服务器网络：SSH 可用，但直接 HTTPS 下载外网文件很慢或会
被防火墙阻挡。服务器需要能通过 SSH 读取 GitHub 仓库，并且需要安装 Go。
脚本会自动从 `PATH`、`/usr/local/go/bin/go`、`/usr/bin/go` 等位置查找 Go，
并按 `go.mod` 校验版本。若 Go 安装在其他位置，设置：

```bash
NEMO_GO=/absolute/path/to/go ./deploy/systemd/install-git-updater.sh
```

如果要立即运行一次：

```bash
systemctl --user start nemo-git-update.service
```

### 备用路线：Release 包更新（仅当服务器 HTTPS 可用）

如果服务器能稳定访问 GitHub Release asset，可以改用预构建包：

```bash
# 在服务器上的部署目录运行；NEMO_REPO 是 GitHub 的 owner/repo
NEMO_DEPLOY_DIR=/path/to/nemo-knows \
NEMO_REPO=yourname/nemo-knows \
NEMO_ARCH=amd64 \
NEMO_RELEASE_TAG=main-latest \
./deploy/systemd/install-release-updater.sh
```

`nemo-release-update.timer` 会下载 `main-latest` 的对应架构 tarball 和
`checksums.txt`，校验后替换二进制并重启服务。

仓库包含 systemd user unit 安装脚本：

```bash
# 在服务器上的部署目录运行
NEMO_DEPLOY_DIR=/path/to/nemo-knows \
NEMO_WEB_ADDR=127.0.0.1:8787 \
NEMO_MAINTAIN_MODE=auto \
./deploy/systemd/install-user-units.sh
```

它会安装并启用：

- `nemo-web.service`：常驻 `nemo-web`
- `nemo-wiki-maintain.timer`：定时运行 wiki 自动维护

分支、预发布分支、测试分支走手动路径。你可以先切换到对应 ref，再本地
构建部署：

```bash
git fetch origin test-branch
git switch --detach origin/test-branch
NEMO_RUN_TESTS=true ./deploy/release/build-local-current.sh
```

也可以让脚本显式拉一个 ref：

```bash
NEMO_DEPLOY_REF=test-branch ./deploy/release/build-local-current.sh
```

## 5. Worker 自动化

Worker 部署 workflow 位于 `.github/workflows/worker.yml`。由于仓库公开，
它只支持手动 `workflow_dispatch`，不会在 push 时自动部署。

需要在 GitHub Secrets 中配置：

- `CLOUDFLARE_API_TOKEN`
- `CLOUDFLARE_ACCOUNT_ID`

手动运行该 workflow 时会执行：

1. `npm ci`
2. `npm run check`（Wrangler dry-run）
3. `npm run deploy`

Worker 部署后绑定到你想要的公开域名（在 Cloudflare Dashboard 或 wrangler.toml
中配置 `routes`）。

## 缓存行为

| 路径 | 方法 | 缓存 |
|------|------|------|
| `/view*` | GET | 5 分钟 |
| `/graph*` | GET | 5 分钟 |
| `/static/*` | GET | 5 分钟 |
| `/run`, `/build` | POST | 不缓存，写入后清除相关缓存 |
| 其他 | GET | 不缓存 |

## 后续扩展

- 加认证：在 Worker 中检查 `Authorization` header 或接入 Cloudflare Access
- 加 JSON API：给 nemo-web 添加 `/api/` 端点，Worker 按路径分发
- 加 R2 存储：把 wiki 静态资源放到 R2，Worker 直接返回
