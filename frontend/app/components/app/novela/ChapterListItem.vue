<script setup lang="ts">
import type { NovelaChapter } from "~/types/backend/novela";
import { useReadProgress } from "~/composables/api/novela/useReadProgress";

const props = defineProps<{
	chapter: NovelaChapter;
	novelaId: number;
}>();

const { isChapterRead } = useReadProgress();
const read = computed(() => isChapterRead(props.chapter));
</script>

<template>
	<NuxtLink
		:to="`/novela/reader/${novelaId}/${chapter.id}`"
		:class="[
			'flex justify-between items-center p-4 border-b border-zinc-100 dark:border-zinc-800/50 last:border-0 transition-colors cursor-pointer group',
			read
				? 'opacity-50 hover:opacity-80 bg-zinc-50 dark:bg-zinc-900/30'
				: 'hover:bg-zinc-50 dark:hover:bg-zinc-800/50'
		]"
	>
		<div class="flex justify-center items-center gap-3">
			<span class="mt-1 text-zinc-400 font-mono text-sm w-8">#{{ chapter.number }}</span>
			<span
				:class="[
					'font-medium transition-colors',
					read
						? 'text-zinc-500 dark:text-zinc-500'
						: 'text-zinc-900 dark:text-zinc-200 group-hover:text-yellow-600 dark:group-hover:text-yellow-600'
				]"
			>
				{{ chapter.title }}
			</span>
		</div>
	</NuxtLink>
</template>
