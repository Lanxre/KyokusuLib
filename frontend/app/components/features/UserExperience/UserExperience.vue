<script setup lang="ts">
import { computed } from "vue";

interface Props {
	currentExp: number;
	expToNextLevel: number;
	level: number;
}

const props = defineProps<Props>();

const progress = computed(() => {
	if (props.expToNextLevel === 0) return 0;
	return Math.min(100, (props.currentExp / props.expToNextLevel) * 100);
});
</script>

<template>
    <div class="flex flex-row items-center gap-3 p-1.5 pr-3 mt-2 rounded-full border-2 border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 shadow-sm transition-colors duration-300 w-full sm:w-auto min-w-[300px] h-8">
        <div class="flex items-center justify-center bg-zinc-100 dark:bg-zinc-900 rounded-full w-4 h-4 shrink-0 border border-zinc-200 dark:border-zinc-700">
            <span class="text-[10px] font-bold text-zinc-900 dark:text-white">{{ level }}</span>
        </div>
        
        <span class="text-[10px] font-bold text-zinc-900 dark:text-white">
            Ур.
        </span>
        
        <div class="flex-1 h-2.5 bg-zinc-100 dark:bg-zinc-900 rounded-full overflow-hidden relative">
            <div 
                class="absolute inset-y-0 left-0 bg-linear-to-r from-yellow-200 to-yellow-600 rounded-full transition-all duration-500 ease-out"
                :style="{ width: `${progress}%` }"
            ></div>
        </div>

        <span class="text-xs font-medium text-zinc-500 dark:text-zinc-400 whitespace-nowrap shrink-0">
            {{ currentExp }} / {{ expToNextLevel }} очков
        </span>
    </div>
</template>