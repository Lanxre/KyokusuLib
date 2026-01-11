import { useInterfaceSettings } from "~/composables/api/settings/useInterfaceSettings";

export default defineNuxtRouteMiddleware(async (to) => {
  const { syncSettingWithBackend } = useInterfaceSettings();
  const authStore = useAuthStore();

  if (import.meta.server && authStore.isAuthenticated) {
    await syncSettingWithBackend();
  }
});