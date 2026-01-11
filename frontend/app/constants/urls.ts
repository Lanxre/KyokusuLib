export function getApiBase() {
	try {
		const config = useRuntimeConfig();
		return config.public.apiBase;
	} catch (e) {
		return "http://localhost:8080";
	}
}
