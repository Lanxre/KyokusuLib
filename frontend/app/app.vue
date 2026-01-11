<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";
import { useActivityStore } from "@/stores/activity";
import { useInterfaceSettings } from "~/composables/api/settings/useInterfaceSettings";
import TheNotifications from "@/components/ui/TheNotifications/TheNotifications.vue";

const authStore = useAuthStore();
const activityStore = useActivityStore();

const { syncSettingWithBackend } = useInterfaceSettings();

if (import.meta.client) {
  activityStore.initActivityTracking();
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
    <TheNotifications />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>