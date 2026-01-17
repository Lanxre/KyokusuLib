<script setup lang="ts">
import Carousel from "@/components/features/Carousel/Carousel.vue";
import NovelaShelf from "./NovelaShelf.vue";
import NovelaColumns from "./NovelaColumns.vue";
import { useNovela } from "~/composables/api/novela/useNovela";

const { fetchNovels } = useNovela();
const { data: novels } = await useAsyncData('novelas-home', () => 
    fetchNovels({ limit: 10 })
);

</script>

<template>
    <div class="min-h-screen flex flex-col bg-zinc-50 dark:bg-radial-[at_center] dark:from-zinc-800 dark:to-zinc-950 dark:to-90% transition-colors duration-300">
    
    <Carousel v-if="!!novels" :items="novels">
		<template #card="{ item }">
			<div class="flex flex-col gap-2 md:gap-3 w-[150px] md:w-[220px] group/card transition-opacity duration-300">
			<div class="
				relative h-[210px] md:h-[320px] w-full rounded-xl overflow-hidden 
				shadow-lg dark:shadow-2xl border border-zinc-200 dark:border-zinc-800 
				transition-transform duration-300 group-hover/card:scale-[1.02] bg-zinc-200 dark:bg-zinc-900 cursor-pointer
			"
			@click="() => navigateTo(`/novela/${item.id}`)"
			>
				<img 
				:src="staticImage(item.poster_url)" 
				class="h-full w-full object-cover pointer-events-none" 
				alt="cover" 
				/>
				<div class="absolute top-2 left-2 bg-white/90 text-black font-bold text-[10px] md:text-xs px-2 py-1 rounded backdrop-blur-sm shadow-sm">
				★ {{ roundTo(item.rating, 1) }}
				</div>
			</div>

			<div class="flex flex-col items-center px-1">
				<h3 class="cs-text text-center leading-md line-clamp-2 text-foreground text-[15px] leading-[1.15] font-semibold text-balance  dark:font-semibold group-hover/card:text-yellow-600 dark:group-hover/card:text-yellow-400 transition-colors">
					{{ item.title }}
				</h3>
				<span class="text-[10px] md:text-sm text-zinc-500 dark:text-zinc-400">
					{{ item.type }}
				</span>
				
				<button 
					class="mt-2 w-full py-1.5 md:py-2 cursor-pointer rounded-lg bg-zinc-100 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 text-zinc-900 dark:text-white text-xs md:text-sm font-medium transition-colors border border-zinc-200 dark:border-zinc-700"
				>
					Читать
				</button>
			</div>
			</div>
		</template>
	</Carousel>

	<div class="container w-[1200px] mx-auto py-8">
		<NovelaColumns 
            :new-items="Array(7).fill(novels[0]) || []" 
            :popular-items="Array(7).fill(novels[0]) || []" 
        />
	</div>

	<div class="container w-[1200px] mx-auto space-y-8">
		<NovelaShelf 
            title="Популярное сейчас" 
            :items="Array(6).fill(novels[0]) || []"
        />

		<NovelaShelf 
            title="Свежие главы" 
            :items="Array(6).fill(novels[0]) || []"
        />
	</div>

  </div>
</template>