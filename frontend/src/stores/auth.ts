import { createFetch } from "@vueuse/core";
import { defineStore } from "pinia";
import { computed, ref } from "vue";
import type { User } from "@/types/backend/user";
import { BACKEND_URL } from "../const";
import { KyokusuAppRole } from "@/types/enums/role-enum";

const useApi = createFetch({
	baseUrl: BACKEND_URL,
	options: {
		immediate: false,
	},
	fetchOptions: {
		mode: "cors",
		credentials: "include",
	},
});

export const useAuthStore = defineStore("auth", () => {
	const user = ref<User | null>(null);
	const isAuthChecking = ref(true);

	const isAuthenticated = computed(() => !!user.value);
	const isAdmin = computed(() => user.value?.role === KyokusuAppRole.ADMIN);

	const profileFetcher = useApi(`/api/auth/me`, {
		onFetchError(ctx) {
			if (ctx.response?.status === 401) {
				ctx.data = null;
			}
			return ctx;
		},
	}).json<User>();

	async function fetchUser() {
		await profileFetcher.execute();
		if (profileFetcher.data.value) {
			user.value = profileFetcher.data.value;
		} else {
			user.value = null;
		}
	}

	async function initAuth() {
		isAuthChecking.value = true;
		try {
			await fetchUser();
		} catch (e) {
			user.value = null;
		} finally {
			isAuthChecking.value = false;
		}
	}

	async function logout() {
		try {
			await useApi(`/api/auth/logout`).post().execute();
		} catch (e) {
			console.error("Logout error:", e);
		} finally {
			user.value = null;
		}
	}

	return {
		user,
		isAuthChecking,
		isAuthenticated,
		isAdmin,

		isLoading: profileFetcher.isFetching,
		error: profileFetcher.error,
		initAuth,
		fetchUser,
		logout,
	};
});
