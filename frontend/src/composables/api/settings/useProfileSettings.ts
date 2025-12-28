import { ref } from "vue";
import { useApi } from "@/api/api";
import { parseDateToISO } from "@/api/utils/date";
import PersonIcon from "@/assets/images/special/user.png";
import { useAuthStore } from "@/stores/auth";
import { useNotificationStore } from "@/stores/notification";
import type { User } from "@/types/backend/user";

export function useProfileSettings() {
	const authStore = useAuthStore();

	const { user } = authStore;
	const { notify } = useNotificationStore();

	const errors = ref<Record<string, string>>({});
	const isLoading = ref(false);
	const isSuccess = ref(false);
	const isBirthDateDisable = ref(user?.birthday ? true : false);

	const fileInput = ref<HTMLInputElement | null>(null);
	const selectedFile = ref<File | null>(null);

	const bannerInput = ref<HTMLInputElement | null>(null);
	const selectedBannerFile = ref<File | null>(null);

	const profile = ref<User>({ ...user! });

	const validate = (): boolean => {
		errors.value = {};
		let isValid = true;

		const nick = profile.value.name?.trim() || "";
		const allowedCharsRegex = /^[A-Za-z\s]{6,50}$/;
		const hasLowercase = /[a-z]/;
		const hasUppercase = /[A-Z]/;

		if (!nick) {
			errors.value.name = "Никнейм обязателен.";
			isValid = false;
		} else if (!allowedCharsRegex.test(nick)) {
			if (nick.length < 6 || nick.length > 50) {
				errors.value.name = "Длина никнейма от 6 до 50 символов.";
			} else {
				errors.value.name =
					"Никнейм может содержать только латинские буквы (A–Z, a–z) и пробелы.";
			}
			isValid = false;
		} else if (!hasLowercase.test(nick) || !hasUppercase.test(nick)) {
			errors.value.name =
				"Никнейм должен содержать минимум 1 строчную (a–z) и 1 заглавную (A–Z) букву.";
			isValid = false;
		}

		const aboutText = profile.value.about || "";
		if (aboutText.length > 500) {
			errors.value.about = `Превышен лимит символов (${aboutText.length}/500).`;
			isValid = false;
		}

		const birthdayText = profile.value.birthday;
		if (birthdayText === "01.01.1") {
			profile.value.birthday = "";
		}

		return isValid;
	};

	const handleAvatarClick = () => {
		fileInput.value?.click();
	};

	const handleFileChange = (event: Event) => {
		const target = event.target as HTMLInputElement;
		if (target.files && target.files[0]) {
			const file = target.files[0];
			selectedFile.value = file;

			const reader = new FileReader();
			reader.onload = (e) => {
				if (e.target?.result) {
					profile.value.picture = e.target.result as string;
				}
			};
			reader.readAsDataURL(file);
		}
	};

	const removeAvatar = () => {
		profile.value.picture = PersonIcon;
		selectedFile.value = null;
		if (fileInput.value) fileInput.value.value = "";
	};

	const handleBannerClick = () => {
		bannerInput.value?.click();
	};

	const handleBannerChange = (event: Event) => {
		const target = event.target as HTMLInputElement;
		if (target.files && target.files[0]) {
			const file = target.files[0];
			selectedBannerFile.value = file;

			const reader = new FileReader();
			reader.onload = (e) => {
				if (e.target?.result) {
					profile.value.banner = e.target.result as string;
				}
			};
			reader.readAsDataURL(file);
		}
	};

	const removeBanner = () => {
		profile.value.banner = "";
		selectedBannerFile.value = null;
		if (bannerInput.value) bannerInput.value.value = "";
	};

	const saveChanges = async () => {
		if (!validate()) {
			return;
		}

		isLoading.value = true;
		isSuccess.value = false;

		try {
			const updatePayload = {
				name: profile.value.name,
				about: profile.value.about,
				gender: profile.value.gender,
				birthday: parseDateToISO(profile.value.birthday),
				is_public: profile.value.is_public,
			};
			
			const { error: profileError } = await useApi("/api/profile/update", {
				credentials: "include",
			})
				.put(updatePayload)
				.json();

			if (profileError.value) {
				notify({
					title: "Неудача",
					content: "Что-то пошло не так :(",
					type: "error",
				});
				throw new Error(profileError.value);
			}

			if (selectedFile.value) {
				const formData = new FormData();
				formData.append("avatar", selectedFile.value);

				const { data: avatarData, error: avatarError } = await useApi(
					"/api/profile/avatar",
					{ credentials: "include" },
				)
					.post(formData)
					.json();

				if (avatarError.value) {
					notify({
						title: "Неудача",
						content: "Неудалось сохранить аватар",
						type: "error",
					});
					throw new Error("Failed to upload avatar");
				}

				if (avatarData.value && avatarData.value.picture) {
					profile.value.picture = avatarData.value.picture;
				}
			} else if (
				selectedFile.value === null &&
				profile.value.picture === PersonIcon
			) {
				const { error: avatarError } = await useApi("/api/profile/avatar", {
					credentials: "include",
				})
					.delete()
					.json();
				if (avatarError.value) {
					notify({
						title: "Неудача",
						content: "Неудалось удалить аватар",
						type: "error",
					});
					throw new Error("Failed to remove avatar");
				}
			}

			if (selectedBannerFile.value) {
				const formData = new FormData();
				formData.append("banner", selectedBannerFile.value);

				const { data: bannerData, error: bannerError } = await useApi(
					"/api/profile/banner",
					{ credentials: "include" },
				)
					.post(formData)
					.json();

				if (bannerError.value) {
					notify({
						title: "Неудача",
						content: "Неудалось сохранить баннер",
						type: "error",
					});
					throw new Error("Failed to upload banner");
				}

				if (bannerData.value && bannerData.value.banner) {
					profile.value.banner = bannerData.value.banner;
				}
			} else if (
				selectedBannerFile.value === null &&
				!profile.value.banner
			) {
				const { error: bannerError } = await useApi("/api/profile/banner", {
					credentials: "include",
				})
					.delete()
					.json();

				if (bannerError.value) {
					notify({
						title: "Неудача",
						content: "Неудалось удалить баннер",
						type: "error",
					});
					throw new Error("Failed to remove banner");
				}
			}

			authStore.user = { ...authStore.user, ...profile.value };
			isSuccess.value = true;
			notify({
				title: "Успешно",
				content: "Настройки профиля сохранены",
				type: "success",
			});
		} catch (e: any) {
			errors.value.global = "Ошибка при сохранении";
			notify({
				title: "Неудача",
				content: "Настройки профиля не сохранены",
				type: "error",
			});
		} finally {
			isLoading.value = false;
		}
	};

	return {
		profile,
		fileInput,
		bannerInput,
		isBirthDateDisable,
		errors,
		isLoading,
		isSuccess,
		handleAvatarClick,
		handleFileChange,
		removeAvatar,
		handleBannerClick,
		handleBannerChange,
		removeBanner,
		saveChanges,
	};
}