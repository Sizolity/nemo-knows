# Cloudflare Worker Deployment

This repository keeps Cloudflare Worker deployment code outside the main Go
application tree. The Worker is a small, independent project under
`.cloudflare/worker/`, while GitHub Actions provides the remote auto-deploy
trigger.

## Layout

```text
.cloudflare/
  worker/
    package.json
    tsconfig.json
    wrangler.toml
    src/
      index.ts
.github/
  workflows/
    deploy-cloudflare-worker.yml
```

The Worker currently exposes a minimal edge entry point and a health check at
`/healthz`. It can be expanded later without coupling the deploy target to the
local `cmd/nemo` or `internal/web` packages.

## Required Manual Secrets

Do not commit Cloudflare credentials to the repository. Upload these in GitHub:

1. Open the GitHub repository.
2. Go to `Settings` -> `Secrets and variables` -> `Actions`.
3. Add repository secrets:
   - `CLOUDFLARE_API_TOKEN`
   - `CLOUDFLARE_ACCOUNT_ID`

The API token should have permission to deploy Workers for the target account.
For a narrowly scoped token, start with Cloudflare permissions equivalent to
`Workers Scripts:Edit` and account read access for the account that owns the
Worker.

## Automatic Deploys

The workflow in `.github/workflows/deploy-cloudflare-worker.yml` deploys when a
push to `main` changes:

- `.cloudflare/worker/**`
- `.github/workflows/deploy-cloudflare-worker.yml`

It can also be launched manually from the GitHub Actions UI through
`workflow_dispatch`.

This keeps normal wiki/source changes from redeploying the Worker. If every push
to `main` should deploy, remove the `paths` filter from the workflow.

## Local Commands

From the Worker directory:

```sh
cd .cloudflare/worker
npm install
npm run typecheck
npm run deploy
```

Local deploys use the same Wrangler project configuration as CI. Authenticate
with Wrangler locally or provide Cloudflare environment variables in the shell;
do not write secrets into `wrangler.toml`.

## Configuration Notes

- `wrangler.toml` sets `workers_dev = true` so the first deployment can use the
  default workers.dev route.
- `NEMO_WORKER_ENV` is a non-secret Worker variable for environment labeling.
- Secrets that the Worker itself needs at runtime should be uploaded with
  `wrangler secret put <NAME>` or configured in the Cloudflare dashboard, not
  committed to this repository.
