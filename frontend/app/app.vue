<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";
import { useActivityStore } from "@/stores/activity";
import { useNotificationStore } from "@/stores/notification";
import { useInterfaceSettings } from "~/composables/api/settings/useInterfaceSettings";
import { useNotifications } from "~/composables/api/notifications/useNotifications";
import { NotificationContainer as TheNotifications } from "@kyokusu-ui/vue";

const authStore = useAuthStore();
const activityStore = useActivityStore();
const notificationStore = useNotificationStore();

const { syncSettingWithBackend } = useInterfaceSettings();
const { connect: connectNotifications } = useNotifications();

if (import.meta.client) {
	activityStore.initActivityTracking();
	connectNotifications();
}

onMounted(async () => {
	if (authStore.isAuthenticated) {
		await syncSettingWithBackend();
	}

	setTimeout(() => {
		const loader = document.getElementById("app-loading-style");
		if (loader) loader.remove();
	}, 100);
});
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <ClientOnly>
      <TheNotifications 
        :notifications="notificationStore.notifications" 
        @remove="notificationStore.remove" 
      />
    </ClientOnly>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>