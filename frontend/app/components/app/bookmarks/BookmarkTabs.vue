<script setup lang="ts">
import type { BookmarkCategory } from "@/types/frontend/bookmarks";

defineProps<{
    categories: { id: string; label: string; icon: string }[];
    activeCategory: BookmarkCategory;
}>();

defineEmits<{
    (e: 'update:activeCategory', val: BookmarkCategory): void;
}>();
</script>

<template>
    <div class="flex gap-2 overflow-x-auto pb-4 no-scrollbar snap-x">
        <button
            v-for="cat in categories"
            :key="cat.id"
            @click="$emit('update:activeCategory', cat.id as BookmarkCategory)"
            class="px-5 py-1.5 rounded-2xl text-xs font-black uppercase tracking-widest transition-all whitespace-nowrap border cursor-pointer snap-start"
            :class="[
                activeCategory === cat.id 
                    ? 'bg-zinc-900 dark:bg-white text-white dark:text-zinc-900 border-transparent shadow-xl shadow-zinc-900/10 dark:shadow-white/10'
                    : 'bg-white dark:bg-zinc-900 border-zinc-200 dark:border-zinc-800 text-zinc-500 hover:border-zinc-400'
            ]"
        >
            <div class="flex justify-center items-center gap-3">
                <Icon :name="cat.icon" size="18" />
                <span>{{ cat.label }}</span>
            </div>
        </button>
    </div>
</template>