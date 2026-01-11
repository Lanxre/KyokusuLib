import { watch, computed } from "vue"; // Добавьте импорты
import { ThemeSetting } from "@/types/enums/theme-enum";
import { useAuthStore } from "@/stores/auth";
import type { UserInterfaceSettings, UserInterfaceSettingsPatch } from "@/types/backend/profile_settings";
import { useApi } from "../useApi";

export function useInterfaceSettings() {
  const colorMode = useColorMode();
  const authStore = useAuthStore();

  const isShowTag = useState<boolean>('settings-is-show-tag', () => false);
  const isInitialized = useState<boolean>('settings-is-initialized', () => false);

  const isDarkTheme = computed({
    get: () => colorMode.preference === ThemeSetting.DARK_THEME,
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
      console.error("Failed to save theme", e);
    }
  };

  const updateInterfaceSettings = async (settings: any) => {
    if (!authStore.isAuthenticated) return;
    try {
      await useApi("/api/profile/settings/interface", { method: "PATCH", body: settings });
    } catch (e) { console.error(e); }
  };


  const syncSettingWithBackend = async () => {
    if (isInitialized.value || !authStore.isAuthenticated) return;

    try {
      const { data } = await useApi<any>("/api/profile/settings");
      if (data.value) {
        if (data.value.theme) colorMode.preference = data.value.theme;
        if (data.value.is_show_tag !== undefined) isShowTag.value = data.value.is_show_tag;
        
        isInitialized.value = true; 
      }
    } catch (e) {
      console.error("Sync error", e);
    }
  };

  watch(isShowTag, (newVal) => {
    if (import.meta.client && isInitialized.value) {
      useApi("/api/profile/settings/interface", { 
        method: "PATCH", 
        body: { is_show_tag: newVal } 
      });
    }
  });

  return {
    isDarkTheme,
    isShowTag,
    syncSettingWithBackend,
    colorMode,
  };
}