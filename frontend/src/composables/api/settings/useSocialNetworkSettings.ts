import { useApi } from "@/api/api";
import { useNotificationStore } from "@/stores/notification";
import { MessageNetworkSettings, SocialNewtworkSettings } from "@/types/frontend/social-settings";
import { ref } from "vue";

export function useSocialNetworkSettings() {
    const { notify } = useNotificationStore();

    const socialNetworkData = ref<SocialNewtworkSettings>({
        is_discord_connected: false,
        is_google_connected: false
    });

    const syncSocialNetworkConnects = async () => {
        try {
            const { data } = await useApi("/api/socials", { credentials: "include" })
                .json<SocialNewtworkSettings>();

            if (data.value !== undefined && data.value !== null) {
                socialNetworkData.value = data.value
            }

        } catch (e) {
            console.log(e)
        }
    }

    const handleDiscordLogout = async () => {
        try {
            const { data } = await useApi("/api/socials/discord/logout", { credentials: "include" })
                .json<MessageNetworkSettings>();

            if (data.value !== undefined && data.value !== null) {
                notify({
                    title: "Успех",
                    content: "Произошел выход из discord аккаунта",
                    type: "info"
                });
                socialNetworkData.value.is_discord_connected = false;
            }

        } catch (e) {
            console.log(e)
        }
    }

    const handleGoogleLogout = async () => {
        try {
            const { data } = await useApi("/api/socials/google/logout", { credentials: "include" })
                .json<MessageNetworkSettings>();

            if (data.value !== undefined && data.value !== null) {
                notify({
                    title: "Успех",
                    content: "Произошел выход из google аккаунта",
                    type: "info"
                });
                socialNetworkData.value.is_google_connected = false;
            }

        } catch (e) {
            console.log(e)
        }
    }

    const handleGoogleLogin = () => {
		window.location.href = `${import.meta.env.VITE_API_URL}/api/socials/google/login`;
	};

	const handleDiscordLogin = () => {
		window.location.href = `${import.meta.env.VITE_API_URL}/api/socials/discord/login`;
	};

    const handleDiscord = async () => {
        if (socialNetworkData.value.is_discord_connected) {
            await handleDiscordLogout();
        } else {
            handleDiscordLogin();
        }
    };

    const handleGoogle = async () => {
        if (socialNetworkData.value.is_google_connected) {
            await handleGoogleLogout();
        } else {
            handleGoogleLogin();
        }
    };

    syncSocialNetworkConnects();

    return {
        socialNetworkData,
        handleDiscordLogin,
        handleGoogleLogin,
        handleDiscordLogout,
        handleGoogleLogout,
        handleDiscord,
        handleGoogle,

    }
}