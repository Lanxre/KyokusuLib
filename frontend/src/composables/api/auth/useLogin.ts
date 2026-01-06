import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { useApi } from "@/api/api";
import { useAuthStore } from "@/stores/auth";
import { useNotificationStore } from "@/stores/notification";

export function useLogin() {
	const authStore = useAuthStore();
	const { notify } = useNotificationStore();
	const router = useRouter();
	const isLoading = ref(false);
	const errorMessage = ref("");
	const successMessage = ref("");

	const isForgotPasswordMode = ref(false);
	const isResetPasswordMode = ref(false);
	const resetToken = ref("");

	const form = reactive({
		email: "",
		password: "",
	});

	const toggleForgotPasswordMode = () => {
		isForgotPasswordMode.value = !isForgotPasswordMode.value;
		isResetPasswordMode.value = false;
		errorMessage.value = "";
		successMessage.value = "";
		form.password = "";
	};

	const handleManualLogin = async () => {
		errorMessage.value = "";
		isLoading.value = true;

		try {
			const { data, error, response } = await useApi("/api/auth/login", {
				credentials: "include",
			})
				.post(form)
				.json();

			if (error.value || response.value?.status !== 200) {
				console.log(error.value);
				throw new Error("Неверный email или пароль");
			}

			if (data.value) {
				authStore.user = data.value;
				router.push(`/profile/${data.value.id}`);
			}

			notify({
				title: "Успех",
				content: "Успешных вход в аккаунт",
				type: "success",
			});
		} catch (err: any) {
			errorMessage.value = err.message || "Произошла ошибка при входе";
			notify({
				title: "Неудача",
				content: "Произошла ошибка при входе",
				type: "error",
			});
		} finally {
			isLoading.value = false;
		}
	};

	const handleSendResetLink = async () => {
		errorMessage.value = "";
		successMessage.value = "";
		isLoading.value = true;

		try {
			const { error, response } = await useApi("/api/auth/forgot-password")
				.post({ email: form.email })
				.json();

			if (error.value || response.value?.status !== 200) {
				throw new Error("Не удалось отправить письмо. Проверьте email.");
			}

			successMessage.value =
				"Ссылка для сброса пароля отправлена на вашу почту.";
		} catch (err: any) {
			errorMessage.value = err.message || "Ошибка отправки запроса";
		} finally {
			isLoading.value = false;
		}
	};

	const handleResetPassword = async () => {
		errorMessage.value = "";
		successMessage.value = "";
		isLoading.value = true;

		try {
			const { error, response } = await useApi("/api/auth/reset-password")
				.post({
					token: resetToken.value,
					password: form.password,
				})
				.json();

			if (error.value || response.value?.status !== 200) {
				throw new Error("Ошибка сброса пароля. Возможно, ссылка устарела.");
			}

			successMessage.value = "Пароль успешно изменен! Теперь вы можете войти.";

			setTimeout(() => {
				isResetPasswordMode.value = false;
				resetToken.value = "";
				form.password = "";
				successMessage.value = "";
			}, 3000);
		} catch (err: any) {
			errorMessage.value = err.message || "Ошибка сброса пароля";
		} finally {
			isLoading.value = false;
		}
	};

	const handleGoogleLogin = () => {
		window.location.href = `${import.meta.env.VITE_API_URL}/api/auth/google/login`;
	};

	const handleDiscordLogin = () => {
		window.location.href = `${import.meta.env.VITE_API_URL}/api/auth/discord/login`;
	};

	return {
		form,
		isLoading,
		errorMessage,
		successMessage,
		isForgotPasswordMode,
		isResetPasswordMode,
		resetToken,
		toggleForgotPasswordMode,
		handleManualLogin,
		handleSendResetLink,
		handleResetPassword,
		handleGoogleLogin,
		handleDiscordLogin,
	};
}
