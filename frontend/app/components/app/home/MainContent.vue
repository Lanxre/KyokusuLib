<script setup lang="ts">
import Carousel from "@/components/features/Carousel/Carousel.vue";

const mangaList = [
	{
		id: 1,
		title: "Берсерк",
		rating: 9.8,
		img: "https://placehold.co/300x450/18181b/FFF?text=Berserk",
	},
	{
		id: 2,
		title: "Ван-Пис",
		rating: 9.5,
		img: "https://placehold.co/300x450/222/FFF?text=OnePiece",
	},
	{
		id: 3,
		title: "Наруто",
		rating: 8.9,
		img: "https://placehold.co/300x450/333/FFF?text=Naruto",
	},
	{
		id: 4,
		title: "Блич",
		rating: 8.5,
		img: "https://placehold.co/300x450/444/FFF?text=Bleach",
	},
	{
		id: 5,
		title: "Акира",
		rating: 9.0,
		img: "https://placehold.co/300x450/555/FFF?text=Akira",
	},
	{
		id: 1,
		title: "Берсерк",
		rating: 9.8,
		img: "https://placehold.co/300x450/18181b/FFF?text=Berserk",
	},
	{
		id: 2,
		title: "Ван-Пис",
		rating: 9.5,
		img: "https://placehold.co/300x450/222/FFF?text=OnePiece",
	},
	{
		id: 3,
		title: "Наруто",
		rating: 8.9,
		img: "https://placehold.co/300x450/333/FFF?text=Naruto",
	},
	{
		id: 4,
		title: "Блич",
		rating: 8.5,
		img: "https://placehold.co/300x450/444/FFF?text=Bleach",
	},
	{
		id: 5,
		title: "Акира",
		rating: 9.0,
		img: "https://placehold.co/300x450/555/FFF?text=Akira",
	},
];
</script>

<template>
    <div class="min-h-screen flex flex-col bg-zinc-50 dark:bg-radial-[at_center] dark:from-zinc-800 dark:to-zinc-950 dark:to-90% transition-colors duration-300">
    
    <Carousel :items="mangaList">
		<template #card="{ item }">
			<div class="flex flex-col gap-2 md:gap-3 w-[150px] md:w-[220px] group/card transition-opacity duration-300">
			<div class="
				relative h-[210px] md:h-[320px] w-full rounded-xl overflow-hidden 
				shadow-lg dark:shadow-2xl border border-zinc-200 dark:border-zinc-800 
				transition-transform duration-300 group-hover/card:scale-[1.02] bg-zinc-200 dark:bg-zinc-900
			">
				<img 
				:src="item.img" 
				class="h-full w-full object-cover pointer-events-none" 
				alt="cover" 
				/>
				<div class="absolute top-2 left-2 bg-white/90 text-black font-bold text-[10px] md:text-xs px-2 py-1 rounded backdrop-blur-sm shadow-sm">
				★ {{ item.rating }}
				</div>
			</div>

			<div class="flex flex-col items-center px-1">
				<h3 class="text-sm md:text-lg font-bold truncate w-full text-center text-zinc-900 dark:text-zinc-100 group-hover/card:text-yellow-600 dark:group-hover/card:text-yellow-400 transition-colors">
				{{ item.title }}
				</h3>
				<span class="text-[10px] md:text-sm text-zinc-500 dark:text-zinc-400">Манга / Сэйнэн</span>
				
				<button class="mt-2 w-full py-1.5 md:py-2 rounded-lg bg-zinc-100 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 text-zinc-900 dark:text-white text-xs md:text-sm font-medium transition-colors border border-zinc-200 dark:border-zinc-700">
				Читать
				</button>
			</div>
			</div>
		</template>
	</Carousel>
  </div>
</template>