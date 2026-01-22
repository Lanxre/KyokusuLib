<script setup lang="ts">
import { ref, watch } from "vue";
import { getBookmarkCategoryLabel } from "@/types/frontend/bookmarks";
import type { BookmarkCategory } from "@/types/frontend/bookmarks";
import NovelaCard from "@/components/app/novela/NovelaCard.vue";
import { useBookmark } from "~/composables/api/novela/useBookmark";

const props = defineProps<{ userId: number }>();

const activeCategory = ref<BookmarkCategory>('reading');

const { fetchUserBookmarkNovels, userBookmarkNovels, bookmarkCategories, loading} = useBookmark();

const load = () => fetchUserBookmarkNovels(props.userId, activeCategory.value);

watch(activeCategory, load, { immediate: true });
</script>

<template>
    <div class="space-y-6">
        <div class="flex gap-2 overflow-x-auto pb-2 no-scrollbar">
            <button
                v-for="cat in bookmarkCategories"
                :key="cat.id"
                @click="activeCategory = cat.id as BookmarkCategory"
                class="px-4 py-2 rounded-xl text-xs font-black uppercase tracking-widest transition-all whitespace-nowrap border cursor-pointer"
                :class="[
                    activeCategory === cat.id 
                        ? 'bg-zinc-900 dark:bg-white text-white dark:text-zinc-900 border-transparent shadow-lg'
                        : 'bg-white dark:bg-zinc-900 border-zinc-200 dark:border-zinc-800 text-zinc-500 hover:border-zinc-400'
                ]"
            >
                <div class="flex justify-center items-center gap-2">
                    <Icon :name="cat.icon" size="16" />
                    <span>{{ getBookmarkCategoryLabel(cat.id as BookmarkCategory) }}</span>
                </div>
            </button>
        </div>

        <div class="min-h-[400px] relative">
            <div v-if="loading" class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
                <div v-for="i in 5" :key="i" class="aspect-[2/3] rounded-2xl bg-zinc-200 dark:bg-zinc-800 animate-pulse"></div>
            </div>

            <div v-else-if="!!userBookmarkNovels.length" class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-6">
                <NovelaCard
                    v-for="novela in userBookmarkNovels"
                    :key="novela.id"
                    v-bind="novela"
                />
            </div>

            <div v-else class="flex flex-col items-center justify-center py-20 bg-zinc-50 dark:bg-zinc-900/30 border-2 border-dashed border-zinc-200 dark:border-zinc-800 rounded-3xl">
                <Icon name="ph:bookmarks-simple-light" size="48" class="text-zinc-300 dark:text-zinc-700 mb-4" />
                <p class="text-zinc-500 font-medium">В этой категории пока пусто</p>
            </div>
        </div>
    </div>
</template>