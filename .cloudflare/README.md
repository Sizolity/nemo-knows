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
