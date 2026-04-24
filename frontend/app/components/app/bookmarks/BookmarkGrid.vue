<script setup lang="ts">
import { computed } from "vue";
import type { BookmarkCategory } from "@/types/frontend/bookmarks";
import NovelaCard from "@/components/app/novela/NovelaCard.vue";

const props = defineProps<{
    loading: boolean;
    novels: any[];
    categories: { id: string; label: string }[];
}>();

const emit = defineEmits<{
    (e: 'change-status', id: number, category: BookmarkCategory | 'remove'): void;
}>();

const categoryOptions = computed(() => {
    return [
        ...props.categories.map(c => ({ value: c.id, label: c.label })),
        { value: 'remove', label: 'Удалить из закладок' }
    ];
});
</script>

<template>
    <div class="min-h-100 relative">
        <div v-if="loading" class="flex sm:grid overflow-x-auto sm:overflow-visible sm:grid-cols-4 md:grid-cols-5 lg:grid-cols-6 gap-4 sm:gap-6 pb-6 no-scrollbar">
            <div v-for="i in 6" :key="i" class="w-28 sm:w-32 shrink-0 aspect-2/3 rounded-3xl bg-zinc-200 dark:bg-zinc-800 animate-pulse"></div>
        </div>

        <div v-else-if="novels.length" class="flex sm:grid overflow-x-auto sm:overflow-visible sm:grid-cols-4 md:grid-cols-5 lg:grid-cols-6 gap-4 sm:gap-6 pb-6 no-scrollbar snap-x">
            <div 
                v-for="novela in novels" 
                :key="novela.id" 
                class="shrink-0 snap-start relative group"
            >
                <div class="w-40">
                    <NovelaCard v-bind="novela" />
                </div>
                
                <div class="absolute top-2 right-2 opacity-100 sm:opacity-0 group-hover:opacity-100 transition-opacity z-10" @click.prevent>
                    
                    <div class="relative">
                        <div>sadsad</div>
                        <select 
                            class="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
                            @change="(e) => emit('change-status', novela.id, (e.target as HTMLSelectElement).value as any)"
                        >
                            <option disabled selected>Переместить</option>
                            <option v-for="opt in categoryOptions" :key="opt.value" :value="opt.value">
                                {{ opt.label }}
                            </option>
                        </select>
                        <div class="w-7 h-7 rounded-xl bg-zinc-900/80 backdrop-blur border border-white/10 flex items-center justify-center text-white shadow-lg pointer-events-none transition-transform group-hover:scale-110">
                            <Icon name="ph:dots-three-bold" size="18" />
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div v-else class="flex flex-col items-center justify-center p-12 bg-white dark:bg-zinc-900/30 border-2 border-dashed border-zinc-200 dark:border-zinc-800 rounded-[3rem]">
            <Icon name="ph:bookmarks-simple-light" size="64" class="text-zinc-300 dark:text-zinc-700 mb-6" />
            <p class="text-zinc-400 font-bold uppercase tracking-[0.2em] text-xs text-center px-4">В этой категории пока пусто</p>
        </div>
    </div>
</template>