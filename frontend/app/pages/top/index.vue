<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useTotalStatistics } from "~/composables/api/statistics/total-stats";
import { StatisticsPeriodEnum } from "@/types/enums/statistics";
import { Separator } from "@kyokusu-ui/vue";
import StatisticsPeriodSelector from "@/components/app/statistics/StatisticsPeriodSelector.vue";
import StatisticsTable from "@/components/app/statistics/StatisticsTable.vue";
import StatisticsPagination from "@/components/app/statistics/StatisticsPagination.vue";
import StatisticsTopThree from "@/components/app/statistics/StatisticsTopThree.vue";

useSeoMeta({ title: 'Статистика - KyokusuLib' });

const { novelaTotalStats, fetchTotalStatistics, isLoading, total } = useTotalStatistics();

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
				<h1 class="text-3xl sm:text-4xl font-black uppercase tracking-tight text-zinc-900 dark:text-white mb-4 flex items-center justify-center gap-4">
					<Icon name="ph:star-bold" size="36" class="text-yellow-400/30 dark:text-yellow-500/20 -rotate-12" />
					<span class="relative">
						Топы
						<Icon name="ph:star-bold" size="48" class="absolute -top-3 left-1/2 -translate-x-1/2 text-yellow-400/10 dark:text-yellow-500/10 -z-10 rotate-45" />
					</span>
					<Icon name="ph:star-bold" size="36" class="text-yellow-400/30 dark:text-yellow-500/20 rotate-12" />
				</h1>
				<Separator />
			</div>

			<div class="flex flex-col justify-around gap-4 p-4 rounded-4xl border bg-white dark:bg-zinc-900/50 dark:border-zinc-800 mb-4">
                <h3 class="text-[50px] flex justify-center sm:text-4xl font-black uppercase tracking-tight text-zinc-900 dark:text-white mb-4">
    			    Лидеры
    			</h3>
			    <StatisticsTopThree :items="novelaTotalStats" />
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
					:total="total"
				/>

				<div v-if="novelaTotalStats.length > 0" class="mt-6 pt-6">
					<StatisticsPagination
						:current-page="currentPage"
						:has-more="hasMore"
						:loading="isLoading"
						:total="total"
						@prev="prevPage"
						@next="nextPage"
					/>
				</div>
			</div>

		</div>
	</div>
</template>
