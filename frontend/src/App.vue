<script lang="ts" setup>
import { onMounted } from "vue";
import { RouterView } from "vue-router";
import TheNotifications from "@/components/ui/TheNotifications/TheNotifications.vue";
import { useAuthStore } from "@/stores/auth";
import { useInterfaceSettings } from "@/composables/api/settings/useInterfaceSettings";
import { useActivityStore } from "./stores/activity";

const authStore = useAuthStore();
const activityStore = useActivityStore();

onMounted(async () => {
  if (authStore.isAuthenticated) {
    useInterfaceSettings()
    activityStore.initActivityTracking();
  }
  
  setTimeout(() => {
    document.getElementById('app-loading-style')?.remove();
  }, 100);
});
</script>

<template>

  <div class="min-h-screen flex flex-col">
    <TheNotifications/>
    <main class="grow">
      <RouterView />
    </main>

  </div>
  
</template>