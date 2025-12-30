import { createFetch } from "@vueuse/core";

export const useApi = createFetch({
	baseUrl: import.meta.env.VITE_API_URL,
	options: {
		beforeFetch({ options }) {
			return { options };
		},
		async onFetchError(ctx) {
            if (ctx.response && ctx.response.headers.get('content-type')?.includes('application/json')) {
                try {
                    const errorData = await ctx.response.clone().json();
                    ctx.data = errorData;
                    if (errorData && errorData.error) {
                        ctx.error = errorData.error;
                    }
                } catch (e) {
                    console.error("Error parsing failed request:", e);
                }
            }
            return ctx;
        },
	},
});
