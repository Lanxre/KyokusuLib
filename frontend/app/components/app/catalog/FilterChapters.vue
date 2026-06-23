<script setup lang="ts">
import { Input as BaseInput } from '@kyokusu-ui/vue';

const chaptersFrom = defineModel<number | null>('chaptersFrom', { default: null });
const chaptersTo = defineModel<number | null>('chaptersTo', { default: null });
</script>

<template>
	<div class="space-y-2">
		<label class="text-xs font-bold uppercase tracking-wider text-zinc-500 dark:text-zinc-400">
			Количество глав
		</label>
		<div class="flex items-center gap-2">
			<BaseInput
				id="filter-chapters-from"
				v-model="chaptersFrom"
				type="number"
				placeholder="От"
				class="[appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
			/>
			<span class="text-zinc-400 text-sm font-medium shrink-0">—</span>
			<BaseInput
				id="filter-chapters-to"
				v-model="chaptersTo"
				type="number"
				placeholder="До"
				class="[appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
			/>
		</div>
	</div>
</template>
