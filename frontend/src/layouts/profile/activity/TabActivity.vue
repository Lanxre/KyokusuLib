<script setup lang="ts">
import { formatDateUserActivity } from '@/api/utils/str';
import ActivityIcon from './ActivityIcon.vue';
import type { UserActivity } from '@/types/backend/user_activity';

import { useActivityConfig } from '@/composables/api/profile/useActivityConfig';

interface Props {
    activities: UserActivity[];
    isLoading?: boolean;
}

defineProps<Props>();

const { getActivityConfig } = useActivityConfig();
</script>

<template>
    <div class="space-y-4 bg-zinc-50 p-4 border border-zinc-200 dark:border-zinc-700 rounded-2xl dark:bg-radial-[at_center] dark:from-zinc-900 cursor-default dark:to-zinc-950 dark:to-90% text-zinc-900 dark:text-zinc-200 transition-colors duration-300 font-sans">
        
        <div v-if="isLoading" class="flex justify-center py-12">
            <svg class="animate-spin h-8 w-8 text-zinc-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
        </div>

        <div v-else-if="activities.length === 0" class="flex flex-col items-center justify-center py-12 text-center border-2 border-dashed border-zinc-300 dark:border-zinc-800 rounded-xl">
            <div class="w-12 h-12 rounded-full bg-zinc-200 dark:bg-zinc-800 flex items-center justify-center mb-3">
                <ActivityIcon name="activity" class="text-zinc-400"/>
            </div>
            <p class="text-zinc-500 dark:text-zinc-400 font-medium">История активности пуста</p>
        </div>

        <div 
            v-else
            v-for="item in activities" 
            :key="item.id"
            class="group relative flex justify-start items-center gap-4 p-4 bg-white dark:bg-zinc-900/30 border border-zinc-200 dark:border-zinc-800 rounded-xl hover:border-zinc-300 dark:hover:border-zinc-700 transition-all duration-200"
        >
            <div 
                class="shrink-0 w-10 h-10 rounded-full flex items-center justify-center"
                :class="getActivityConfig(item).color"
            >
                <ActivityIcon :name="getActivityConfig(item).icon" />
            </div>

            <div class="flex flex-col flex-1 min-w-0">
                <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-1 mb-1">
                    <h4 class="text-sm font-semibold text-zinc-900 dark:text-zinc-100 truncate">
                        {{ getActivityConfig(item).title }}
                    </h4>
                    <span class="text-xs text-zinc-400 whitespace-nowrap">
                        {{ formatDateUserActivity(item.timestamp) }}
                    </span>
                </div>
                
                <p class="text-sm text-zinc-600 dark:text-zinc-400 leading-relaxed wrap-break-words line-clamp-2">
                    {{ getActivityConfig(item).description }}
                </p>
            </div>
        </div>
    </div>
</template>