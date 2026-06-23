<script setup lang="ts">
import { ref } from 'vue';
import { Separator, Button } from '@kyokusu-ui/vue';
import { useNovelaFilters } from '~/composables/api/novela/useNovelaFilters';
import FilterSearch from './FilterSearch.vue';
import FilterGenre from './FilterGenre.vue';
import FilterCategory from './FilterCategory.vue';
import FilterStatus from './FilterStatus.vue';
import FilterChapters from './FilterChapters.vue';
import type { NovelaFilterOption } from '~/types/frontend/query/novela-filters';

interface Props {
	genreOptions?: NovelaFilterOption[];
	categoryOptions?: NovelaFilterOption[];
}

const props = withDefaults(defineProps<Props>(), {
	genreOptions: () => [],
	categoryOptions: () => [],
});

const emit = defineEmits<{
	apply: [params: Record<string, any>];
	reset: [];
}>();

const {
	filters,
	hasActiveFilters,
	activeFilterCount,
	resetFilters,
	buildQueryParams,
	debouncedSearch,
} = useNovelaFilters();

const localSearch = ref('');

function onSearchUpdate(value: string) {
	localSearch.value = value;
	debouncedSearch(value);
}

function handleReset() {
	localSearch.value = '';
	resetFilters();
	emit('reset');
}

function handleApply() {
	emit('apply', buildQueryParams());
}
</script>

<template>
	<div class="space-y-6">
		<FilterSearch
			v-model="localSearch"
			@update="onSearchUpdate"
		/>

		<Separator />

		<FilterStatus
			v-model:status="filters.status"
			v-model:translationStatus="filters.translationStatus"
			v-model:type="filters.type"
		/>

		<Separator />

		<FilterGenre
			v-if="genreOptions.length"
			:options="genreOptions"
			:model-value="filters.genres"
			@update:model-value="(val) => { filters.genres = val; }"
		/>

		<template v-if="genreOptions.length && categoryOptions.length">
			<Separator />
		</template>

		<FilterCategory
			v-if="categoryOptions.length"
			:options="categoryOptions"
			:model-value="filters.categories"
			@update:model-value="(val) => { filters.categories = val; }"
		/>

		<template v-if="categoryOptions.length">
			<Separator />
		</template>

		<FilterChapters
			v-model:chapters-from="filters.chaptersFrom"
			v-model:chapters-to="filters.chaptersTo"
		/>

		<Separator />

		<div class="space-y-2.5">
			<Button
				variant="primary"
				class="w-full flex items-center justify-center gap-4"
				@click="handleApply"
			>
				<Icon name="ph:magnifying-glass" size="18" />
				Применить
			</Button>

			<button
				v-if="hasActiveFilters"
				@click="handleReset"
				class="w-full py-2 text-sm font-medium text-zinc-500 hover:text-zinc-700 dark:hover:text-zinc-300 transition-colors cursor-pointer"
			>
				Сбросить фильтры ({{ activeFilterCount }})
			</button>
		</div>
	</div>
</template>
