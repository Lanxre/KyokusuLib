<script setup lang="ts">
import { computed } from "vue";
import type { NovelaBookmarkCategory } from "@/types/backend/novela";

const props = defineProps<{
	bookmarkDetails: NovelaBookmarkCategory;
	isExpanded: boolean;
}>();


const emit = defineEmits(['toggle']);

const statusConfig: Record<string, { label: string; color: string; icon: string }> = {
	reading: { label: "Читают", color: "bg-emerald-500", icon: "M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" },
	completed: { label: "Прочитано", color: "bg-blue-500", icon: "M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" },
	planned: { label: "В планах", color: "bg-zinc-400", icon: "M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" },
	on_hold: { label: "Отложено", color: "bg-amber-500", icon: "M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z" },
	dropped: { label: "Брошено", color: "bg-red-500", icon: "M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" },
};

const maxCount = computed(() => {
	if (!props.bookmarkDetails.nc_items?.length) return 0;
	return Math.max(...props.bookmarkDetails.nc_items.map((item) => item.count));
});

const getPercentage = (count: number) => {
	if (maxCount.value === 0) return 0;
	return (count / maxCount.value) * 100;
};
</script>

<template>
	<div class="w-full bg-white dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 rounded-2xl overflow-hidden transition-all duration-300">
		<button
			@click="$emit('toggle')"
			class="w-full flex items-center justify-between p-4 cursor-pointer hover:bg-zinc-50 dark:hover:bg-zinc-800/50 transition-colors outline-none"
		>
			<div class="flex items-center gap-3">
				<div class="flex flex-col items-start">
					<span class="text-[10px] font-black uppercase tracking-[0.2em] text-zinc-500">В библиотеках</span>
					<div class="flex items-baseline gap-2">
						<span class="text-2xl font-black text-zinc-900 dark:text-white">
							{{ bookmarkDetails.total }}
						</span>
						<span class="text-sm font-medium text-zinc-400">человек</span>
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
						v-for="item in bookmarkDetails.nc_items" 
						:key="item.value"
						class="flex items-center gap-3 group"
					>
						<div class="flex items-center gap-2 min-w-[100px]">
							<svg class="w-4 h-4 shrink-0" :class="statusConfig[item.value]?.color.replace('bg-', 'text-')" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
								<path stroke-linecap="round" stroke-linejoin="round" :d="statusConfig[item.value]?.icon" />
							</svg>
							<span class="text-xs font-bold text-zinc-500 dark:text-zinc-400">
								{{ statusConfig[item.value]?.label || item.value }}
							</span>
						</div>
						
						<div class="flex-1 h-1.5 bg-zinc-100 dark:bg-zinc-800 rounded-full overflow-hidden">
							<div 
								class="h-full transition-all duration-700 ease-out"
								:class="statusConfig[item.value]?.color"
								:style="{ width: `${getPercentage(item.count)}%` }"
							></div>
						</div>

						<div class="flex items-center justify-end gap-1.5 min-w-[45px] group-hover:translate-x-0.5 transition-transform duration-300">
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