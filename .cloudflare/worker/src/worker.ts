interface Env {
	NEMO_ORIGIN: string;
}

// Paths whose GET responses are cached at the edge.
const CACHEABLE_PREFIXES = ["/view", "/graph", "/static/"];
const CACHE_TTL = 300; // seconds

export default {
	async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
		const url = new URL(request.url);
		const origin = env.NEMO_ORIGIN.replace(/\/+$/, "");
		const target = origin + url.pathname + url.search;

		if (request.method === "GET" && isCacheable(url.pathname)) {
			const cache = caches.default;
			const cacheKey = new Request(target, request);
			const hit = await cache.match(cacheKey);
			if (hit) return hit;

			const res = await forward(request, target);
			if (res.ok) {
				const cached = new Response(res.body, res);
				cached.headers.set("Cache-Control", `public, max-age=${CACHE_TTL}`);
				ctx.waitUntil(cache.put(cacheKey, cached.clone()));
				return cached;
			}
			return res;
		}

		// POST /run, /build etc. — pass through, bust cache for related pages.
		const res = await forward(request, target);
		if (request.method === "POST" && res.ok) {
			ctx.waitUntil(purgeViewCache(env));
		}
		return res;
	},
} satisfies ExportedHandler<Env>;

function isCacheable(pathname: string): boolean {
	return CACHEABLE_PREFIXES.some((p) => pathname.startsWith(p));
}

async function forward(original: Request, target: string): Promise<Response> {
	const headers = new Headers(original.headers);
	headers.delete("host");

	const init: RequestInit = {
		method: original.method,
		headers,
		redirect: "follow",
	};
	if (original.method !== "GET" && original.method !== "HEAD") {
		init.body = original.body;
	}
	return fetch(target, init);
}

// Best-effort cache purge after mutations so the next read sees fresh data.
async function purgeViewCache(env: Env) {
	const cache = caches.default;
	const origin = env.NEMO_ORIGIN.replace(/\/+$/, "");
	for (const prefix of ["/view", "/graph", "/"]) {
		await cache.delete(new Request(origin + prefix));
	}
}
