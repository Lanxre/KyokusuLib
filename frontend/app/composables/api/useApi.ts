import type { UseFetchOptions } from "nuxt/app";
import { defu } from "defu";

export function useApi<T>(
	url: string | (() => string),
	options: UseFetchOptions<T> = {},
) {
	const defaults: UseFetchOptions<T> = {
		baseURL: "/api",
		headers: useRequestHeaders(["cookie"]),
	};

	const params = defu(options, defaults);

	return useFetch(url, params);
}
