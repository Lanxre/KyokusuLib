<script setup lang="ts">
import NovelaRowItem from "./NovelaRowItem.vue";
import type { NovelaDetails } from "~/types/backend/novela";

interface Props {
	newItems: NovelaDetails[];
	popularItems: NovelaDetails[];
}
defineProps<Props>();
</script>

<template>
    <div class="flex flex-col lg:flex-row items-start justify-around gap-8 px-4 py-6 bg-white dark:bg-[#0f0f0f] border border-zinc-200 dark:border-zinc-800 rounded-[2.5rem] shadow-sm">
        
        <div class="flex flex-col w-full lg:w-1/2 gap-6">
            <div class="flex items-center justify-between px-2">
                <h3 class="text-lg font-black uppercase tracking-widest italic border-l-4 border-yellow-500 pl-3 text-zinc-900 dark:text-white">
                    Горячие новинки
                </h3>
                <NuxtLink to="/catalog?sort=new" class="text-[10px] font-bold uppercase text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-200 transition-colors">
                    Смотреть все
                </NuxtLink>
            </div>
            
            <div class="flex flex-col gap-1">
                <NovelaRowItem 
                    v-for="(item, index) in newItems" 
                    :key="item.id" 
                    :novela="item" 
                    :index="index"
                />
            </div>
        </div>

        <div class="hidden lg:block w-px self-stretch bg-linear-to-b from-transparent via-zinc-200 dark:via-zinc-800 to-transparent"></div>

        <div class="flex flex-col w-full lg:w-1/2 gap-6">
            <div class="flex items-center justify-between px-2">
                <h3 class="text-lg font-black uppercase tracking-widest italic border-l-4 border-yellow-500 pl-3 text-zinc-900 dark:text-white">
                    Популярно сегодня
                </h3>
                <NuxtLink to="/catalog?sort=popular" class="text-[10px] font-bold uppercase text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-200 transition-colors">
                    Смотреть все
                </NuxtLink>
            </div>
            
            <div class="flex flex-col gap-1">
                <NovelaRowItem 
                    v-for="(item, index) in popularItems" 
                    :key="item.id" 
                    :novela="item" 
                    :index="index"
                />
            </div>
        </div>

    </div>
</template>