<script setup lang="ts">
import { staticImage } from "@/utils/str";
import type { SearchCategory } from "@/composables/api/search/useSearch";

const props = defineProps<{
    results: any[];
    activeCategory: SearchCategory;
}>();

const emit = defineEmits(["select"]);

const getIcon = () => {
    switch (props.activeCategory) {
        case "ranobe": return "ph:book-open-bold";
        case "users": return "ph:user-bold";
        case "teams": return "ph:users-three-bold";
        case "authors": return "ph:pen-nib-bold";
        default: return "ph:image-bold";
    }
};

const getTitle = (item: any) => {
    return item.title || item.name || item.label || 'Без названия';
};

const getSubtitle = (item: any) => {
    return item.type || item.role || item.description || '';
};
</script>

<template>
    <div class="space-y-2">
        <button 
            v-for="item in results" 
            :key="item.id"
            @click="emit('select', item)"
            class="w-full flex items-center gap-4 p-3 rounded-xl hover:bg-zinc-50 dark:hover:bg-zinc-800/80 transition-colors text-left group cursor-pointer"
        >
            <div class="w-12 h-16 shrink-0 rounded-lg bg-zinc-200 dark:bg-zinc-800 overflow-hidden border border-zinc-200 dark:border-zinc-700">
                <img 
                    v-if="item.poster_url || item.picture || item.avatar || item.banner" 
                    :src="staticImage(item.poster_url || item.picture || item.avatar || item.banner)" 
                    class="w-full h-full object-cover" 
                />
                <div v-else class="w-full h-full flex items-center justify-center bg-zinc-100 dark:bg-zinc-800 text-zinc-400">
                    <Icon :name="getIcon()" size="20" />
                </div>
            </div>
            <div class="flex flex-col flex-1 min-w-0">
                <span class="text-sm font-bold text-zinc-900 dark:text-zinc-100 truncate group-hover:text-yellow-600 dark:group-hover:text-yellow-500 transition-colors">
                    {{ getTitle(item) }}
                </span>
                <span v-if="getSubtitle(item)" class="text-xs text-zinc-500 dark:text-zinc-400 truncate mt-0.5">
                    {{ getSubtitle(item) }}
                </span>
            </div>
            <Icon name="ph:caret-right-bold" size="16" class="text-zinc-300 dark:text-zinc-600 opacity-0 group-hover:opacity-100 transition-opacity" />
        </button>
    </div>
</template>
