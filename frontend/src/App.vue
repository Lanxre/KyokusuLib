<script lang="ts" setup>
import { onMounted } from "vue";
import { RouterView } from "vue-router";
import TheNotifications from "@/components/ui/TheNotifications/TheNotifications.vue";
import { useAuthStore } from "@/stores/auth";
import { useInterfaceSettings } from "@/composables/api/settings/useInterfaceSettings";
import { useActivityStore } from "./stores/activity";

const authStore = useAuthStore();
const activityStore = useActivityStore();

useInterfaceSettings();

onMounted(async () => {
  activityStore.initActivityTracking();
  
  setTimeout(() => {
    document.getElementById('app-loading-style')?.remove();
  }, 100);

  await authStore.initAuth();
});
</script>

<template>

  <div class="min-h-screen flex flex-col">
    <TheNotifications/>
    <div v-if="authStore.isAuthChecking" class="h-screen w-full flex items-center justify-center bg-zinc-50 dark:bg-[#0f0f0f] text-zinc-900 dark:text-white">
        <svg class="animate-spin h-8 w-8 text-zinc-900 dark:text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
    </div>
    <main v-else class="grow">
      <RouterView />
    </main>

  </div>
  
</template>