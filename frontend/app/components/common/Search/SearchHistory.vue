<script setup lang="ts">
import type { MostSearched } from '@/types/frontend/search/searches';

defineProps<{
    recentSearches: string[];
    popularSearches: MostSearched[];
}>();

const emit = defineEmits(["select", "remove", "clear"]);

</script>

<template>
    <div class="space-y-8">
        <!-- Recent Searches -->
        <div v-if="recentSearches.length > 0">
            <div class="flex items-center justify-between mb-3">
                <h3 class="text-xs font-bold text-zinc-500 uppercase tracking-widest">Недавние поиски</h3>
                <button @click="emit('clear')" class="text-[10px] font-bold text-zinc-400 dark:hover:text-zinc-100 transition-colors uppercase tracking-widest bg-zinc-100 hover:bg-red-50 dark:bg-zinc-800/50 dark:hover:bg-zinc-500/10 px-2 py-1 rounded-md">
                    Очистить
                </button>
            </div>
            <div class="flex flex-wrap gap-2">
                <div 
                    v-for="(recent, i) in recentSearches" 
                    :key="i"
                    class="h-8 group flex items-center bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-700/80 hover:border-yellow-500/50 dark:hover:border-yellow-500/50 rounded-full pl-4 pr-2 py-1.5 shadow-sm transition-all duration-200 cursor-pointer"
                >
                    <Icon name="ph:clock-counter-clockwise-bold" size="14" class="text-zinc-400 group-hover:text-yellow-500 transition-colors" />
                    <button 
                        @click="emit('select', recent)" 
                        class="text-sm cursor-pointer font-semibold text-zinc-700 dark:text-zinc-300 group-hover:text-zinc-900 dark:group-hover:text-white mx-2 focus:outline-none"
                    >
                        {{ recent }}
                    </button>
                    <button @click.stop="emit('remove', recent)" class="px-0.5 py-0.5 mt-0.5 cursor-pointer rounded-full text-zinc-400 dark:hover:text-zinc-100 hover:bg-red-50 dark:hover:bg-zinc-500/10 transition-colors focus:outline-none">
                        <Icon name="ph:x-bold" size="14" />
                    </button>
                </div>
            </div>
        </div>

        <!-- Popular Searches -->
        <div>
            <h3 class="text-xs font-bold text-zinc-500 uppercase tracking-widest mb-3">Часто ищут</h3>
            <div class="flex flex-wrap gap-2">
                <button 
                    v-for="(popular, i) in popularSearches" 
                    :key="i"
                    @click="emit('select', popular)"
                    class="group flex items-center gap-1.5 px-4 py-1.5 bg-yellow-50 dark:bg-yellow-500/10 text-yellow-600 dark:text-yellow-500 border border-yellow-200/50 dark:border-yellow-500/20 hover:bg-yellow-100 hover:border-yellow-300 dark:hover:bg-yellow-500/20 dark:hover:border-yellow-500/40 rounded-full text-sm font-semibold transition-all duration-200 cursor-pointer shadow-sm active:scale-95"
                >
                    <Icon name="ph:trend-up-bold" size="14" class="opacity-70 group-hover:opacity-100 group-hover:scale-110 transition-all" />
                    {{ popular.label }}
                </button>
            </div>
        </div>
    </div>
</template>
