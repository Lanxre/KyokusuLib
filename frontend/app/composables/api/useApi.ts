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

	return useFetch(url, defu(options, defaults));
}

export function $api<T>(
	url: string,
	options: any = {},
) {
	const defaults = {
		baseURL: "/api",
		headers: useRequestHeaders(["cookie"]),
	};

	return $fetch<T>(url, defu(options, defaults));
}