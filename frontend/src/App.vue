<script lang="ts" setup>
import { onMounted } from "vue";
import { RouterView } from "vue-router";
import TheNotifications from "@/components/ui/TheNotifications/TheNotifications.vue";
import { useAuthStore } from "@/stores/auth";
import { useInterfaceSettings } from "@/composables/api/settings/useInterfaceSettings";

const authStore = useAuthStore();
useInterfaceSettings();

onMounted(async () => {
  
  setTimeout(() => {
    document.getElementById('app-loading-style')?.remove();
  }, 100);

  await authStore.initAuth();
});
</script>

<template>

  <div class="min-h-screen flex flex-col">
    <TheNotifications/>
    <div v-if="authStore.isAuthChecking" class="h-screen w-full flex items-center justify-center bg-[#0f0f0f] text-white">
      Загрузка...
    </div>
    <main v-else class="flex-grow">
      <RouterView />
    </main>

  </div>
  
</template>