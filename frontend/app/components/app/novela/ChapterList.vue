<script setup lang="ts">
import { ref, computed } from "vue";
import type { NovelaVolume } from "@/types/backend/novela";

const props = defineProps<{
	volumes: NovelaVolume[];
}>();

const searchQuery = ref("");
const activeVolumeId = ref(props.volumes[0]?.id);

const activeVolume = computed(() =>
	props.volumes.find((v) => v.id === activeVolumeId.value),
);

const filteredChapters = computed(() => {
	if (!activeVolume.value) return [];
	if (!searchQuery.value) return activeVolume.value.chapters;
	return activeVolume.value.chapters.filter(
		(c) =>
			c.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
			c.number.toString().includes(searchQuery.value),
	);
});

const formatDate = (date: string) => {
	return new Date(date).toLocaleDateString("ru-RU", {
		day: "numeric",
		month: "short",
	});
};
</script>

<template>
    <div class="space-y-4">
        <div class="flex flex-col sm:flex-row gap-4 justify-between items-start sm:items-center">
            <div class="flex gap-2 overflow-x-auto no-scrollbar max-w-full">
                <button 
                    v-for="vol in volumes" 
                    :key="vol.id"
                    @click="activeVolumeId = vol.id"
                    class="px-4 py-2 rounded-lg text-sm font-medium whitespace-nowrap transition-colors border"
                    :class="activeVolumeId === vol.id 
                        ? 'bg-zinc-900 text-white border-zinc-900 dark:bg-zinc-100 dark:text-zinc-900 dark:border-zinc-100' 
                        : 'bg-white dark:bg-zinc-900 border-zinc-200 dark:border-zinc-700 text-zinc-600 dark:text-zinc-400 hover:border-zinc-400'"
                >
                    Том {{ vol.number }}
                </button>
            </div>
            
            <div class="relative w-full sm:w-64">
                <input 
                    v-model="searchQuery"
                    type="text" 
                    placeholder="Поиск главы..." 
                    class="w-full pl-9 pr-4 py-2 bg-zinc-100 dark:bg-zinc-800 border-none rounded-lg text-sm text-zinc-900 dark:text-zinc-200 placeholder-zinc-400 focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
                >
                <svg class="absolute left-3 top-2.5 w-4 h-4 text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
            </div>
        </div>

        <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl overflow-hidden">
            <div 
                v-for="chapter in filteredChapters" 
                :key="chapter.id"
                class="flex justify-between items-center p-4 border-b border-zinc-100 dark:border-zinc-800/50 last:border-0 hover:bg-zinc-50 dark:hover:bg-zinc-800/50 transition-colors cursor-pointer group"
            >
                <div class="flex items-center gap-3">
                    <span class="text-zinc-400 font-mono text-sm w-8">#{{ chapter.number }}</span>
                    <span class="text-zinc-900 dark:text-zinc-200 font-medium group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors">{{ chapter.title }}</span>
                </div>
                <span class="text-xs text-zinc-400">{{ formatDate(new Date().getDate().toString()) }}</span>
            </div>
            
            <div v-if="filteredChapters.length === 0" class="p-8 text-center text-zinc-500">
                Главы не найдены
            </div>
        </div>
    </div>
</template>

<style scoped>
.no-scrollbar::-webkit-scrollbar { display: none; }
.no-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }
</style>