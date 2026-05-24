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

## 4. GitHub CD

仓库包含 GitHub Actions CD workflow（`.github/workflows/cd.yml`）：

- push 到 `main` 或手动 `workflow_dispatch` 时运行。
- 先执行 `go test ./...`。
- 为服务器构建 Linux 二进制产物：
  - `nemo-knows-linux-amd64.tar.gz`
  - `nemo-knows-linux-arm64.tar.gz`
- 每个产物包含 `nemo`、`nemo-web`、`README.md`、`AGENTS.md`。
- 产物通过 GitHub Actions artifacts 保存，服务器可以下载后替换本地
  `.bin/nemo` 和 `.bin/nemo-web`。

Cloudflare Worker 仍走 Cloudflare 自己的拉取/部署流程。GitHub CD 不运行
`wrangler deploy`，也不需要 Cloudflare API token。

服务器更新示例：

```bash
# 从 GitHub Actions 下载对应架构的 nemo-knows-linux-*.tar.gz 后：
tar -xzf nemo-knows-linux-amd64.tar.gz
install -m 0755 nemo-knows-linux-amd64/nemo .bin/nemo
install -m 0755 nemo-knows-linux-amd64/nemo-web .bin/nemo-web
```

后续如果要让 GitHub Actions 直接部署到服务器，可以再加 SSH deploy job
（需要配置服务器 host、user、private key、部署目录和 systemd 服务名）。

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
