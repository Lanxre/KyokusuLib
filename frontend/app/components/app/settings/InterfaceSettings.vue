<script lang="ts" setup>
import ThemeIcon from "@/assets/images/special/theme.png";
import { useInterfaceSettings } from "@/composables/api/settings/useInterfaceSettings";
import BaseToggle from "@/components/ui/BaseToggle/BaseToggle.vue";

const { isDarkTheme, isShowTag, syncSettingWithBackend } = useInterfaceSettings();

await useAsyncData('settings-sync', () => syncSettingWithBackend());

</script>

<template>
    <div class="space-y-8">
        <div class="flex flex-col items-center">
            <h2 class="text-3xl font-bold mb-1">Интерфейс</h2>
            <p class="text-sm text-zinc-500">Настройте внешний вид приложения.</p>
        </div>

        <div class="space-y-6 text-left">
            <div class="pb-6 border-b space-y-4 border-zinc-200 dark:border-zinc-800">
                <h3 class="text-base font-semibold">Оформление</h3>
                
                <div class="flex flex-col items-start gap-4 p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-800">
                    <div class="flex items-center gap-4">
                        <span class="text-sm font-medium">Тёмная тема</span>
                        <img :src="ThemeIcon" class="w-6 h-6 dark:invert"/>
                    </div>
                    <BaseToggle v-model="isDarkTheme" id="theme_toggle" />
                </div>
                
                <div class="flex flex-col items-start gap-4 p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-800">
                    <div class="flex items-center gap-4">
                        <span class="text-sm font-medium">Показывать тэг пользователя</span>
                        <img :src="ThemeIcon" class="w-6 h-6 dark:invert"/>
                    </div>
                    <BaseToggle v-model="isShowTag" id="show_tag_toggle" />
                </div>
            </div>
        </div>
    </div>
</template>