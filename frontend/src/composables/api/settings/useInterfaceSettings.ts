import { ref, watch } from "vue";
import { useApi } from "@/api/api";
import type { UserInterfaceSettings } from "@/types/backend/profile_settings";
import { ThemeSetting } from "@/types/enums/theme-enum";

export function useInterfaceSettings() {
    const isDarkTheme = ref(document.documentElement.classList.contains(ThemeSetting.DARK_THEME));

    const updateBackendTheme = async (theme: string) => {
        try {
            await useApi("/api/profile/settings", { credentials: "include" })
                .put({ theme })
                .json();
        } catch (e) {
            console.error("Failed to save theme preference", e);
        }
    };

    const applyTheme = (isDark: boolean) => {
        const root = document.documentElement;
        const newTheme = isDark ? ThemeSetting.DARK_THEME : ThemeSetting.LIGHT_THEME;

        if (isDark) {
            root.classList.add(ThemeSetting.DARK_THEME);
        } else {
            root.classList.remove(ThemeSetting.DARK_THEME);
        }
        
        localStorage.setItem(ThemeSetting.LOCAL_STORAGE_KEY, newTheme);
        
        updateBackendTheme(newTheme);
    };

    const syncWithBackend = async () => {
        try {
            const { data } = await useApi("/api/profile/settings", { credentials: "include" })
                .json<UserInterfaceSettings>();

            if (data.value && data.value.theme) {
                const backendIsDark = data.value.theme === ThemeSetting.DARK_THEME;
                if (isDarkTheme.value !== backendIsDark) {
                    isDarkTheme.value = backendIsDark;
                }
            }
        } catch (e) {
            console.log(e)
        }
    };

    watch(isDarkTheme, (val) => {
        applyTheme(val);
    });

    syncWithBackend();

    return {
        isDarkTheme,
    };
}