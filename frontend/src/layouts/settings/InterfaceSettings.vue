<script lang="ts" setup>
import ThemeIcon from "@/assets/images/special/theme.png"
import { useInterfaceSettings } from "@/composables/api/settings/useInterfaceSettings";
import BaseToggle from "@/components/ui/BaseToggle/BaseToggle.vue";

const { isDarkTheme, isShowTag } = useInterfaceSettings();
</script>

<template>
    <div class="space-y-8 cursor-default transition-colors duration-300">
        <div>
            <h2 class="text-xl font-medium mb-1 text-zinc-900 dark:text-white">Интерфейс</h2>
            <p class="text-sm text-zinc-500 dark:text-stone-400">Настройте внешний вид приложения под свои предпочтения.</p>
        </div>

        <div class="space-y-6 text-left">
            <div class="pb-6 border-b space-y-4 border-zinc-200 dark:border-zinc-800">
                <h3 class="text-base leading-xl font-semibold text-zinc-800 dark:text-zinc-200">Оформление</h3>
                
                <div class="flex flex-col items-start gap-4 justify-between p-4 rounded-lg border transition-colors duration-300 bg-white border-zinc-200 dark:bg-zinc-900/50 dark:border-zinc-800">
                    <div class="space-y-1">
                        <div class="flex col items-center gap-4">
                            <span class="text-sm font-medium block text-zinc-700 dark:text-zinc-200">Тёмная тема</span>
                            <img :src="ThemeIcon" class="w-6 h-6 invert-0 dark:invert"/>
                        </div>
                        <p class="text-xs text-zinc-500 dark:text-stone-500">
                            Переключите для использования светлой или темной цветовой схемы.
                        </p>
                    </div>
                    <BaseToggle 
                        v-model="isDarkTheme" 
                        id="theme_toggle"
                    />
                </div>
                
                <div class="flex flex-col items-start gap-4 justify-between p-4 rounded-lg border transition-colors duration-300 bg-white border-zinc-200 dark:bg-zinc-900/50 dark:border-zinc-800">
                    <div class="space-y-1">
                        <div class="flex col items-center gap-4">
                            <span class="text-sm font-medium block text-zinc-700 dark:text-zinc-200">Показывать тэг пользователя</span>
                            <img :src="ThemeIcon" class="w-6 h-6 invert-0 dark:invert"/>
                        </div>
                        <p class="text-xs text-zinc-500 dark:text-stone-500">
                            Переключите для показа или скрытия тэга пользователя.
                        </p>
                    </div>
                    <BaseToggle 
                        v-model="isShowTag" 
                        id="show_tag_toggle"
                    />
                </div>
            </div>
        </div>
    </div>
</template>