<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useTotalStatistics } from "~/composables/api/statistics/total-stats";
import { StatisticsPeriodEnum } from "@/types/enums/statistics";
import { Separator } from "@kyokusu-ui/vue";
import StatisticsPeriodSelector from "@/components/app/statistics/StatisticsPeriodSelector.vue";
import StatisticsTable from "@/components/app/statistics/StatisticsTable.vue";
import StatisticsPagination from "@/components/app/statistics/StatisticsPagination.vue";

useSeoMeta({ title: 'Статистика - KyokusuLib' });

const { novelaTotalStats, fetchTotalStatistics, isLoading } = useTotalStatistics();

const currentPeriod = ref<StatisticsPeriodEnum | undefined>(undefined);
const currentPage = ref(1);
const limit = 20;

const offset = computed(() => (currentPage.value - 1) * limit);
const hasMore = computed(() => novelaTotalStats.value.length >= limit);

const loadData = async () => {
	await fetchTotalStatistics(currentPeriod.value, limit, offset.value);
};

onMounted(() => {
	loadData();
});

const setPeriod = (period: StatisticsPeriodEnum | undefined) => {
	if (currentPeriod.value === period) return;
	currentPeriod.value = period;
	currentPage.value = 1;
	loadData();
};

const nextPage = () => {
	if (!hasMore.value) return;
	currentPage.value++;
	loadData();
};

const prevPage = () => {
	if (currentPage.value <= 1) return;
	currentPage.value--;
	loadData();
};
</script>

<template>
	<div class="min-h-screen bg-zinc-50 dark:bg-[#0f0f0f] text-zinc-900 dark:text-zinc-200 transition-colors duration-300">
		<div class="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-12">

			<div class="mb-8">
				<h1 class="text-3xl sm:text-4xl font-black uppercase tracking-tight text-zinc-900 dark:text-white mb-4">
					Статистика
				</h1>
				<Separator />
			</div>

			<div class="bg-white dark:bg-zinc-900/40 rounded-[2.5rem] border border-zinc-200 dark:border-zinc-800 p-6 shadow-sm mb-8">
				<div class="mb-6">
					<StatisticsPeriodSelector
						:model-value="currentPeriod"
						@update:model-value="setPeriod"
					/>
				</div>

				<div v-if="isLoading" class="flex justify-center items-center py-12">
					<Icon name="ph:spinner-gap-bold" size="32" class="animate-spin text-zinc-400" />
					<span class="ml-3 text-zinc-500 font-medium">Загрузка...</span>
				</div>

				<div v-else-if="novelaTotalStats.length === 0" class="flex flex-col items-center justify-center py-16 text-center">
					<Icon name="ph:chart-bar-bold" size="48" class="text-zinc-300 dark:text-zinc-700 mb-4" />
					<p class="text-zinc-500 dark:text-zinc-400 font-medium">Нет данных для отображения</p>
				</div>

				<StatisticsTable
					v-else
					:items="novelaTotalStats"
					:offset="offset"
				/>

				<div v-if="novelaTotalStats.length > 0" class="mt-6 pt-6 border-t border-zinc-200 dark:border-zinc-800">
					<StatisticsPagination
						:current-page="currentPage"
						:has-more="hasMore"
						:loading="isLoading"
						@prev="prevPage"
						@next="nextPage"
					/>
				</div>
			</div>

		</div>
	</div>
</template>
