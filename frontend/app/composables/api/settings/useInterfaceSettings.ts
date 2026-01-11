import { watch } from "vue";
import { ThemeSetting } from "@/types/enums/theme-enum";
import { useAuthStore } from "@/stores/auth";
import { useApi } from "../useApi";
import type { UserInterfaceSettings, UserInterfaceSettingsPatch } from "@/types/backend/profile_settings";

export function useInterfaceSettings() {
  const colorMode = useColorMode();
  const authStore = useAuthStore();
  const isShowTag = useState('isShowTag', () => false);
  
  const isDarkTheme = computed({
    get: () => colorMode.value === ThemeSetting.DARK_THEME,
    set: (val: boolean) => {
      colorMode.preference = val ? ThemeSetting.DARK_THEME : ThemeSetting.LIGHT_THEME;
    }
  });

  const updateBackendTheme = async (theme: string) => {
    if (!authStore.isAuthenticated) return;
    try {
      await useApi("/api/profile/settings", {
        method: "PUT",
        body: { theme },
      });
    } catch (e) {
      console.error("Failed to save theme preference", e);
    }
  };

  const updateInterfaceSettings = async (settings: UserInterfaceSettingsPatch) => {
    if (!authStore.isAuthenticated) return;
    try {
      await useApi("/api/profile/settings/interface", {
        method: "PATCH",
        body: settings,
      });
    } catch (e) {
      console.error("Failed to save interface settings", e);
    }
  };

  const syncSettingWithBackend = async () => {
    if (!authStore.isAuthenticated) return;
    try {
      const { data } = await useApi<UserInterfaceSettings>("/api/profile/settings");

      if (data.value?.theme) {
        colorMode.preference = data.value.theme;
      }

      if (data.value?.is_show_tag !== undefined) {
        isShowTag.value = data.value.is_show_tag;
      }
    } catch (e) {
      console.error("Sync error", e);
    }
  };

  watch(() => colorMode.preference, (newTheme) => {
    if (import.meta.client) {
      updateBackendTheme(newTheme);
    }
  });

  watch(isShowTag, (val) => {
    updateInterfaceSettings({ is_show_tag: val });
  });

  return {
    isDarkTheme,
    isShowTag,
    syncSettingWithBackend,
    colorMode 
  };
}