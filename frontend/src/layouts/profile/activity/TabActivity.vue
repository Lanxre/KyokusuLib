<script setup lang="ts">
import { computed } from 'vue';
import { formatDateUserActivity } from '@/api/utils/str';
import ActivityIcon from './ActivityIcon.vue';
import type { UserActivity } from '@/types/backend/user_activity';
import { useActivityConfig } from '@/composables/api/profile/useActivityConfig';

interface Props {
    activities: UserActivity[];
    isLoading?: boolean;
}

const props = defineProps<Props>();
const { getActivityConfig } = useActivityConfig();

const processedActivities = computed(() => {
    return props.activities.map(activity => ({
        ...activity,
        ui: getActivityConfig(activity)
    }));
});
</script>

<template>
    <div class="space-y-4 bg-zinc-50 p-4 border border-zinc-200 dark:border-zinc-700 rounded-2xl dark:bg-radial-[at_center] dark:from-zinc-900 cursor-default dark:to-zinc-950 dark:to-90% text-zinc-900 dark:text-zinc-200 transition-colors duration-300 font-sans min-h-[300px]">

        <div v-if="isLoading" class="space-y-4">
            <div v-for="i in 3" :key="i" class="flex items-start gap-4 p-4 border border-zinc-200 dark:border-zinc-800 rounded-xl animate-pulse">
                <div class="shrink-0 w-10 h-10 rounded-full bg-zinc-200 dark:bg-zinc-800"></div>
                <div class="flex-1 space-y-2">
                    <div class="flex justify-between">
                        <div class="h-4 bg-zinc-200 dark:bg-zinc-800 rounded w-1/3"></div>
                        <div class="h-3 bg-zinc-200 dark:bg-zinc-800 rounded w-16"></div>
                    </div>
                    <div class="h-3 bg-zinc-200 dark:bg-zinc-800 rounded w-3/4"></div>
                </div>
            </div>
        </div>

        <div v-else-if="activities.length === 0" class="flex flex-col items-center justify-center py-12 text-center border-2 border-dashed border-zinc-300 dark:border-zinc-800 rounded-xl h-full">
            <div class="w-12 h-12 rounded-full bg-zinc-200 dark:bg-zinc-800 flex items-center justify-center mb-3">
                <ActivityIcon name="activity" class="text-zinc-400"/>
            </div>
            <p class="text-zinc-500 dark:text-zinc-400 font-medium">История активности пуста</p>
        </div>

        <TransitionGroup 
            v-else 
            name="list" 
            tag="div" 
            class="space-y-4"
        >
            <div 
                v-for="item in processedActivities" 
                :key="item.id"
                class="group relative flex justify-start items-center gap-4 p-4 bg-white dark:bg-zinc-900/30 border border-zinc-200 dark:border-zinc-800 rounded-xl hover:border-zinc-300 dark:hover:border-zinc-700 transition-all duration-200 hover:shadow-sm dark:hover:shadow-none"
            >
                <div 
                    class="shrink-0 w-10 h-10 rounded-full flex items-center justify-center transition-transform group-hover:scale-110 duration-300"
                    :class="item.ui.color"
                >
                    <ActivityIcon :name="item.ui.icon" />
                </div>

                <div class="flex flex-col flex-1 min-w-0 cursor-pointer">
                    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-center gap-4 mb-1">
                        <h4 class="text-sm font-semibold text-zinc-900 dark:text-zinc-100 truncate group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors">
                            {{ item.ui.title }}
                        </h4>
                        <span class="text-xs text-zinc-400 whitespace-nowrap">
                            {{ formatDateUserActivity(item.timestamp) }}
                        </span>
                    </div>
                    
                    <p class="text-sm text-zinc-600 dark:text-zinc-400 leading-relaxed wrap-break-words line-clamp-2">
                        {{ item.ui.description }}
                    </p>
                </div>
            </div>
        </TransitionGroup>
    </div>
</template>

<style scoped>
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(-10px);
}
</style>