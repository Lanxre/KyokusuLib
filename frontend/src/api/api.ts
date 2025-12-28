import { createFetch } from "@vueuse/core";

export const useApi = createFetch({
	baseUrl: import.meta.env.VITE_API_URL,
	options: {
		beforeFetch({ options }) {
			return { options };
		},
		onFetchError({ data, error }) {
			console.log("API Error:", error);
			return { data, error };
		},
	},
});
