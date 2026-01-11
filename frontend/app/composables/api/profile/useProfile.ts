import { ref, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useUserApi } from "~/composables/api/user/userApi";
import { useAuthStore } from "@/stores/auth";
import { GenderSetting } from "@/types/enums/gender-enum";
import { KyokusuAppRole } from "@/types/enums/role-enum";
import type { GetUserDto } from "@/types/backend/user";

export function useProfile() {
	const route = useRoute();
	const router = useRouter();
	const userApi = useUserApi();
	const authStore = useAuthStore();

	const isLoading = ref(false);
	const fetchedUser = ref<GetUserDto | null>(null);

	const profileTabs = [
		{ id: "overview", label: "Обзор" },
		{ id: "bookmarks", label: "Закладки" },
		{ id: "comments", label: "Комментарии" },
	];

	const isSelfProfile = computed(() => {
		const routeId = Number(route.params.id);
		return authStore.isAuthenticated && authStore.user?.id === routeId;
	});

	const profileData = computed<GetUserDto | null>(() => {
		if (isSelfProfile.value) {
			return authStore.user;
		}
		return fetchedUser.value;
	});

	const isPublicAccount = computed(() => !!profileData.value?.is_public);

	const getRoleColor = (role?: string) => {
		switch (role) {
			case KyokusuAppRole.ADMIN:
				return "text-red-400 bg-red-400/10 border-red-400/20";
			case KyokusuAppRole.MODERATOR:
				return "text-blue-400 bg-blue-400/10 border-blue-400/20";
			default:
				return "text-zinc-400 bg-zinc-400/10 border-zinc-400/20";
		}
	};

	const getGenderText = (gender?: string) => {
		switch (gender) {
			case GenderSetting.MALE:
				return "Мужской";
			case GenderSetting.FEMALE:
				return "Женский";
			case GenderSetting.HIDDEN:
			default:
				return "Не указан";
		}
	};

	const formatLastLogin = (last_login: string | undefined | null) => {
		if (!last_login) return "Неизвестно";

		const date = new Date(last_login);
		const today = new Date();
		date.setHours(date.getHours() - 3);

		const isToday =
			date.getFullYear() === today.getFullYear() &&
			date.getMonth() === today.getMonth() &&
			date.getDate() === today.getDate();

		if (isToday) {
			return date.toLocaleTimeString("ru-RU", {
				hour: "2-digit",
				minute: "2-digit",
			});
		} else {
			return date.toLocaleDateString("ru-RU", {
				day: "2-digit",
				month: "2-digit",
				year: "numeric",
			});
		}
	};

	const checkIsLogin = (last_login: string | undefined | null) => {
		if (!last_login) return false;

		const date = new Date(last_login);
		const today = new Date();
		date.setHours(date.getHours() - 3);

		return (
			date.getFullYear() === today.getFullYear() &&
			date.getMonth() === today.getMonth() &&
			date.getDate() === today.getDate()
		);
	};

	const accountCreated = computed(() => {
		if (isSelfProfile.value && profileData.value) {
			return new Date(profileData.value.create_at).toLocaleDateString("ru-RU");
		}
		return "Неизвестно";
	});

	const lastLogin = computed(() => {
		if (
			!profileData.value ||
			(!profileData.value.is_public && !isSelfProfile.value)
		) {
			return "Неизвестно";
		}
		return formatLastLogin(profileData.value.last_login);
	});

	const isLogin = computed(() => {
		if (
			!profileData.value ||
			(!profileData.value.is_public && !isSelfProfile.value)
		) {
			return false;
		}
		return checkIsLogin(profileData.value.last_login);
	});

	const userRoleColor = computed(() => getRoleColor(profileData.value?.role));
	const userGender = computed(() => getGenderText(profileData.value?.gender));

	const loadProfileData = async (id: number) => {
		if (isSelfProfile.value) {
			fetchedUser.value = null;
			return;
		}

		isLoading.value = true;
		try {
			const data = await userApi.getUser(id);
			if (!data) {
				router.replace({ name: "notFound" });
				return;
			}
			fetchedUser.value = data;
		} catch (error) {
			console.error(error);
			fetchedUser.value = null;
		} finally {
			isLoading.value = false;
		}
	};

	const init = async () => {
		const id = Number(route.params.id);
		if (!isNaN(id)) {
			await loadProfileData(id);
		}
	};

	watch(() => route.params.id, init, { immediate: true });

	return {
		profileData,
		isSelfProfile,
		isPublicAccount,
		isLoading,
		profileTabs,

		accountCreated,
		lastLogin,
		isLogin,
		userRoleColor,
		userGender,

		getRoleColor,
		getGenderText,
		getLastLogin: formatLastLogin,
		getIsLogin: checkIsLogin,
		loadProfileData,
	};
}
