<script setup lang="ts">
import CatalogFilters from '@/components/app/catalog/CatalogFilters.vue';

const emit = defineEmits<{
	apply: [params: Record<string, any>];
	reset: [];
}>();
</script>

<template>
	<div
		class="min-h-150 bg-white dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 shadow-sm hidden lg:block"
	>
		<h2 class="text-sm font-bold text-zinc-900 dark:text-white mb-6 flex items-center gap-2">
			<Icon name="ph:funnel-bold" size="16" class="text-yellow-500" />
			Фильтры
		</h2>
		<CatalogFilters @apply="(params) => emit('apply', params)" @reset="() => emit('reset')" />
	</div>
</template>
