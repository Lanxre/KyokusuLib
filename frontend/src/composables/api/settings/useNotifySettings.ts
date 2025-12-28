import { ref, watch, nextTick } from "vue";
import { useApi } from "@/api/api";
import type { UserNotifySettings } from "@/types/backend/profile_settings";

export function useNotifySettings() {
    const profileSettings = ref<UserNotifySettings>({
        is_app_notify: true,
        is_new_published_notify: true
    });
    
    const isSyncing = ref(false);

    const updateBackendNotify = async (settings: Partial<UserNotifySettings>) => {
        try {
            await useApi("/api/profile/settings/notify", { credentials: "include" })
                .patch(settings)
                .json();
        } catch (e) {
            console.error("Failed to save notify preference", e);
        }
    };

    const syncWithBackend = async () => {
        isSyncing.value = true;
        try {
            const { data } = await useApi("/api/profile/settings/notify", { credentials: "include" })
                .json<UserNotifySettings>();

            if (data.value) {
                profileSettings.value = data.value;
            }
        } catch (e) {
            console.error(e);
        } finally {
            await nextTick();
            isSyncing.value = false;
        }
    };

    watch(() => profileSettings.value.is_app_notify, (newVal) => {
        if (!isSyncing.value) {
            updateBackendNotify({ is_app_notify: newVal });
        }
    });

    watch(() => profileSettings.value.is_new_published_notify, (newVal) => {
        if (!isSyncing.value) {
            updateBackendNotify({ is_new_published_notify: newVal });
        }
    });

    syncWithBackend();

    return {
        profileSettings,
    };
}