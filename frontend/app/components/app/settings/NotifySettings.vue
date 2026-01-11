<script lang="ts" setup>
import EmailIcon from "@/assets/images/special/email.png";
import ChatBubbleIcon from "@/assets/images/special/chat-bubble.png";
import BaseToggle from "@/components/ui/BaseToggle/BaseToggle.vue";
import { useNotifySettings } from "@/composables/api/settings/useNotifySettings";

const { profileSettings } = useNotifySettings();
</script>

<template>
    <div class="space-y-8 cursor-default transition-colors duration-300">
        <div class="flex flex-col items-center">
            <h2 class="text-3xl font-bold mb-1 text-zinc-900 dark:text-white">Уведомления</h2>
            <p class="text-sm text-zinc-500 dark:text-stone-400">Настройте уведомления приложения.</p>
        </div>

        <div class="space-y-6 text-left">
            <div class="pb-6 border-b space-y-4 border-zinc-200 dark:border-zinc-800">
                <h3 class="text-base leading-xl font-semibold text-zinc-800 dark:text-zinc-200">На сайте</h3>
                
                <div class="flex items-center justify-between p-4 rounded-lg border transition-colors duration-300 bg-white border-zinc-200 dark:bg-zinc-900/50 dark:border-zinc-800">
                    <div class="space-y-1">
                        <div class="flex items-center gap-4">
                            <span class="text-sm font-medium block text-zinc-700 dark:text-zinc-200">Оповещения</span>
                            <img :src="EmailIcon" class="w-6 h-6 invert-0 dark:invert"/>
                        </div>
                        <p class="text-xs text-zinc-500 dark:text-stone-500">
                            Переключите для отключения оповещений на сайте.
                        </p>
                    </div>
                    <BaseToggle 
                        v-if="profileSettings"
                        v-model="profileSettings.is_app_notify" 
                        id="app_notify_toggle"
                    />
                </div>

                <div class="flex items-center justify-between p-4 rounded-lg border transition-colors duration-300 bg-white border-zinc-200 dark:bg-zinc-900/50 dark:border-zinc-800">
                    <div class="space-y-1">
                        <div class="flex items-center gap-4">
                            <span class="text-sm font-medium block text-zinc-700 dark:text-zinc-200">Выход новых работ</span>
                            <img :src="ChatBubbleIcon" class="w-6 h-6 invert-0 dark:invert"/>
                        </div>
                        <p class="text-xs text-zinc-500 dark:text-stone-500">
                            Переключите для отключения оповещений на сайте.
                        </p>
                    </div>
                    <BaseToggle 
                        v-if="profileSettings"
                        v-model="profileSettings.is_new_published_notify" 
                        id="new_published_toggle"
                    />
                </div>
            </div>
        </div>
    </div>
</template>