<script setup lang="ts">
import { computed } from "vue";
import { useReaderSettings } from "~/composables/ui/useReaderSettings";

const { fontSize } = useReaderSettings();

const fillPercent = computed(() => ((fontSize.value - 12) / (40 - 12)) * 100);
</script>

<template>
	<div class="rounded-2xl bg-zinc-50/80 dark:bg-zinc-800/40 border border-zinc-100 dark:border-zinc-800 p-5 space-y-4">
		<div class="flex items-center justify-between">
			<div class="flex items-center gap-2.5">
				<div class="w-8 h-8 rounded-lg bg-yellow-500/10 flex items-center justify-center">
					<Icon name="ph:text-aa-bold" size="16" class="text-yellow-600 dark:text-yellow-400" />
				</div>
				<label class="text-sm font-bold text-zinc-500 dark:text-zinc-400 uppercase tracking-wider">Размер шрифта</label>
			</div>
			<div class="flex items-baseline gap-0.5">
				<span class="text-lg font-black font-mono text-yellow-500 tabular-nums">{{ fontSize }}</span>
				<span class="text-xs font-bold text-zinc-400">px</span>
			</div>
		</div>

		<div class="px-0.5">
			<div class="relative">
				<input
					v-model.number="fontSize"
					type="range"
					min="12"
					max="40"
					step="1"
					class="reader-slider w-full"
					:style="{ '--fill': `${fillPercent}%` }"
				/>
			</div>
			<div class="flex justify-between mt-2 text-[10px] text-zinc-400 font-bold uppercase tracking-tight select-none">
				<span>12</span>
				<span>40</span>
			</div>
		</div>
	</div>
</template>

<style scoped>
.reader-slider {
	-webkit-appearance: none;
	appearance: none;
	height: 6px;
	cursor: pointer;
	border-radius: 999px;
	outline: none;
	background: transparent;
}

.reader-slider::-webkit-slider-runnable-track {
	height: 6px;
	border-radius: 999px;
	background: linear-gradient(to right, #eab308 var(--fill, 0%), #e4e4e7 var(--fill, 0%));
}

:root.dark .reader-slider::-webkit-slider-runnable-track {
	background: linear-gradient(to right, #eab308 var(--fill, 0%), #27272a var(--fill, 0%));
}

.reader-slider::-webkit-slider-thumb {
	-webkit-appearance: none;
	appearance: none;
	width: 18px;
	height: 18px;
	border-radius: 50%;
	background: #eab308;
	cursor: pointer;
	margin-top: -6px;
	box-shadow: 0 2px 8px rgba(234, 179, 8, 0.45);
	transition: transform 0.15s ease, box-shadow 0.15s ease;
}

.reader-slider::-webkit-slider-thumb:hover {
	transform: scale(1.2);
	box-shadow: 0 3px 12px rgba(234, 179, 8, 0.55);
}

.reader-slider::-moz-range-track {
	height: 6px;
	border-radius: 999px;
	background: #e4e4e7;
}

:root.dark .reader-slider::-moz-range-track {
	background: #27272a;
}

.reader-slider::-moz-range-thumb {
	width: 18px;
	height: 18px;
	border-radius: 50%;
	background: #eab308;
	cursor: pointer;
	border: none;
	box-shadow: 0 2px 8px rgba(234, 179, 8, 0.45);
}

.reader-slider::-moz-range-progress {
	height: 6px;
	border-radius: 999px;
	background: #eab308;
}
</style>
