<script setup lang="ts">
import type { NovelaChapter } from "~/types/backend/novela";

defineProps<{
	chapter: NovelaChapter;
}>();
</script>

<template>
	<div
		class="flex justify-between items-center p-4 border-b border-zinc-100 dark:border-zinc-800/50 last:border-0 hover:bg-zinc-50 dark:hover:bg-zinc-800/50 transition-colors cursor-pointer group"
	>
		<div class="flex justify-center items-center gap-3">
			<span class="mt-1 text-zinc-400 font-mono text-sm w-8">#{{ chapter.number }}</span>
			<span
				class="text-zinc-900 dark:text-zinc-200 font-medium group-hover:text-yellow-600 dark:group-hover:text-yellow-600 transition-colors"
			>
				{{ chapter.title }}
			</span>
		</div>
	</div>
</template>
