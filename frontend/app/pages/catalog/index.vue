<script setup lang="ts">
import { PageLayout, Separator } from '@kyokusu-ui/vue';
import { useCatalog } from '@/composables/api/novela/useCatalog';
import CatalogNovelaList from '@/components/app/catalog/CatalogNovelaList.vue';
import CatalogFilterSidebar from '@/components/app/catalog/CatalogFilterSidebar.vue';
import CatalogMobileFilters from '@/components/app/catalog/CatalogMobileFilters.vue';

const {
	isFilterOpen,
	toggleFilters,
	closeFilters,
	hasMore,
	allNovels,
	isFetchingMore,
	isInitialLoading,
	loadNextPage,
	onApplyFilters,
	onResetFilters,
} = await useCatalog();
</script>

<template>
	<PageLayout>
		<div
			class="
				min-h-full bg-zinc-50 dark:bg-[#0f0f0f] text-zinc-900 dark:text-zinc-200 transition-colors duration-300
				border border-zinc-200 dark:border-zinc-800 rounded-3xl sm:rounded-[2.5rem] shadow-sm overflow-hidden
			"
		>
			<div class="w-full max-w-8xl mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-12">
				<div class="mb-10">
					<h1 class="text-3xl sm:text-4xl font-black uppercase tracking-tight text-zinc-900 dark:text-white mb-4 flex items-center gap-4">
						<Icon name="ph:books-bold" size="36" class="text-yellow-400/30 dark:text-yellow-500/20 -rotate-6" />
						<span class="relative">
							Каталог
							<Icon name="ph:book-open-bold" size="44" class="absolute -top-2 left-1/2 -translate-x-1/2 text-yellow-400/10 dark:text-yellow-500/10 -z-10 rotate-12" />
						</span>
					</h1>
					<Separator />
				</div>

				<div class="grid grid-cols-1 lg:grid-cols-[7fr_3fr] gap-12">
					<CatalogNovelaList
						:novels="allNovels"
						:is-initial-loading="isInitialLoading"
						:is-fetching-more="isFetchingMore"
						:has-more="hasMore"
						@toggle-filters="toggleFilters"
						@intersect="loadNextPage"
					/>

					<CatalogFilterSidebar
						@apply="onApplyFilters"
						@reset="onResetFilters"
					/>
				</div>
			</div>
		</div>

		<CatalogMobileFilters
			:is-open="isFilterOpen"
			@close="closeFilters"
			@apply="onApplyFilters"
			@reset="onResetFilters"
		/>
	</PageLayout>
</template>