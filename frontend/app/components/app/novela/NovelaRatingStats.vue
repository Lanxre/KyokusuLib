<script setup lang="ts">
import { computed } from "vue";
import type { NovelaRatingCategory } from "@/types/backend/novela";

const props = defineProps<{
	ratingDetails: NovelaRatingCategory;
	isExpanded: boolean;
}>();

const maxCount = computed(() => {
	if (!props.ratingDetails.nc_items?.length) return 0;
	return Math.max(...props.ratingDetails.nc_items.map((item) => item.count));
});

const getPercentage = (count: number) => {
	if (maxCount.value === 0) return 0;
	return (count / maxCount.value) * 100;
};

const emit = defineEmits(['toggle']);
</script>

<template>
	<div class="w-full bg-white dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 rounded-2xl overflow-hidden transition-all duration-300">
		<button 
		@click="$emit('toggle')"
			class="w-full flex items-center justify-between p-4 cursor-pointer hover:bg-zinc-50 dark:hover:bg-zinc-800/50 transition-colors outline-none"
		>
			<div class="flex items-center gap-3">
				<div class="flex flex-col items-start">
					<span class="text-[10px] font-black uppercase tracking-[0.2em] text-zinc-500">Рейтинг</span>
					<div class="flex items-baseline gap-2">
						<span class="text-2xl font-black text-zinc-900 dark:text-white">
							{{ ratingDetails.total_rating.toFixed(1) }}
						</span>
						<span class="text-sm font-medium text-zinc-400">
							/ {{ ratingDetails.total }} оценок
						</span>
					</div>
				</div>
			</div>
			
			<div 
				class="w-8 h-8 rounded-full flex items-center justify-center bg-zinc-100 dark:bg-zinc-800 transition-transform duration-300"
				:class="{ 'rotate-180': isExpanded }"
			>
				<svg class="w-4 h-4 text-zinc-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
					<path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
				</svg>
			</div>
		</button>

		<Transition name="expand">
			<div v-if="isExpanded" class="px-4 pb-4 overflow-hidden">
				<div class="pt-4 border-t border-zinc-100 dark:border-zinc-800 space-y-3">
					<div 
						v-for="item in ratingDetails.nc_items" 
						:key="item.value"
						class="flex items-center gap-1 group"
					>
						<div class="flex items-center gap-1.5 min-w-11.25">
							<svg class="w-3 h-3 text-yellow-500 shrink-0" fill="currentColor" viewBox="0 0 20 20">
								<path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
							</svg>
							<span class="text-xs font-bold text-zinc-500 dark:text-zinc-400">{{ item.value }}</span>
						</div>
						
						<div class="flex-1 h-1.5 bg-zinc-100 dark:bg-zinc-800 rounded-full overflow-hidden">
							<div 
								class="h-full bg-yellow-500 transition-all duration-700 ease-out"
								:style="{ width: `${getPercentage(item.count)}%` }"
							></div>
						</div>

						<div class="flex items-center justify-center  gap-1.5 min-w-11.25 group-hover:translate-x-0.5 transition-transform duration-300">
							<span class="text-[11px] font-black text-zinc-400 group-hover:text-zinc-900 dark:group-hover:text-zinc-200 transition-colors">
								{{ item.count }}
							</span>
							<svg class="w-3 h-3 text-zinc-300 dark:text-zinc-600" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
								<path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
							</svg>
						</div>
					</div>
				</div>
			</div>
		</Transition>
	</div>
</template>

<style scoped>
.expand-enter-active,
.expand-leave-active {
	transition: all 0.3s ease-in-out;
	max-height: 500px;
}

.expand-enter-from,
.expand-leave-to {
	max-height: 0;
	opacity: 0;
}
</style>