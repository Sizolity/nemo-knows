export interface Env {
	NEMO_WORKER_ENV?: string;
}

const textHeaders = {
	"content-type": "text/plain; charset=utf-8",
};

export default {
	async fetch(request: Request, env: Env): Promise<Response> {
		const url = new URL(request.url);

		if (url.pathname === "/healthz") {
			return Response.json({
				ok: true,
				service: "nemo-knows-worker",
				env: env.NEMO_WORKER_ENV ?? "unset",
			});
		}

		return new Response(
			[
				"nemo-knows Cloudflare Worker",
				"",
				"This is the deployment entry point for edge automation.",
				"Health check: /healthz",
			].join("\n"),
			{ headers: textHeaders },
		);
	},
};
