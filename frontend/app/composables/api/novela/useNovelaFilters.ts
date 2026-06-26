import { reactive, computed, toRefs } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import type { NovelaFilterState } from '~/types/frontend/query/novela-filters';
import { DEFAULT_FILTER_STATE } from '~/types/frontend/query/novela-filters';

export function useNovelaFilters() {
	const filters = reactive<NovelaFilterState>({
		...DEFAULT_FILTER_STATE,
	});

	const hasActiveFilters = computed(() => {
		return (
			filters.search !== '' ||
			filters.genres.length > 0 ||
      filters.categories.length > 0 ||
      filters.country.length > 0 ||
			filters.status !== '' ||
			filters.translationStatus !== '' ||
			filters.type !== '' ||
			filters.chaptersFrom !== null ||
			filters.chaptersTo !== null ||
			filters.yearFrom !== null ||
			filters.yearTo !== null
		);
	});

	const activeFilterCount = computed(() => {
		let count = 0;
		if (filters.search !== '') count++;
		count += filters.genres.length;
    count += filters.categories.length;
    count += filters.country.length;
		if (filters.status !== '') count++;
		if (filters.translationStatus !== '') count++;
		if (filters.type !== '') count++;
		if (filters.chaptersFrom !== null || filters.chaptersTo !== null) count++;
		if (filters.yearFrom !== null || filters.yearTo !== null) count++;
		return count;
	});

	function buildQueryParams() {
		const params: Record<string, any> = {};

		if (filters.search) params.search = filters.search;
		if (filters.genres.length > 0) params.genres = filters.genres;
    if (filters.categories.length > 0) params.categories = filters.categories;
    if (filters.country.length > 0) params.country = filters.country;
		if (filters.status) params.status = filters.status;
		if (filters.translationStatus) params.translation_status = filters.translationStatus;
		if (filters.type) params.type = filters.type;
		if (filters.chaptersFrom) params.chapters_from = filters.chaptersFrom;
		if (filters.chaptersTo) params.chapters_to = filters.chaptersTo;
		if (filters.yearFrom) params.year_from = filters.yearFrom;
		if (filters.yearTo) params.year_to = filters.yearTo;

		return params;
	}

	function resetFilters() {
		Object.assign(filters, DEFAULT_FILTER_STATE);
	}

	function setGenre(genre: string, selected: boolean) {
		if (selected) {
			if (!filters.genres.includes(genre)) {
				filters.genres.push(genre);
			}
		} else {
			filters.genres = filters.genres.filter((g) => g !== genre);
		}
	}

	function setCategory(category: string, selected: boolean) {
		if (selected) {
			if (!filters.categories.includes(category)) {
				filters.categories.push(category);
			}
		} else {
			filters.categories = filters.categories.filter((c) => c !== category);
		}
	}

	const debouncedSearch = useDebounceFn((value: string) => {
		filters.search = value;
	}, 400);

	return {
		filters,
		hasActiveFilters,
		activeFilterCount,
		buildQueryParams,
		resetFilters,
		setGenre,
		setCategory,
		debouncedSearch,
	};
}
