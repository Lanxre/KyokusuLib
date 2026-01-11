<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
	rating: number;
	count: number;
}>();

const starColor = computed(() => {
	if (props.rating >= 8) return "text-green-500";
	if (props.rating >= 6) return "text-yellow-500";
	return "text-red-500";
});
</script>

<template>
    <div class="flex items-center px-2 py-1.5 rounded-lg">
        <div class="flex flex-row items-center gap-2 leading-none">
            <span class="text-3xl font-bold text-zinc-900 dark:text-white text-right" :style="{ color: starColor }">{{ rating.toFixed(1) }}</span>
            <span class="text-[10px] text-zinc-400 uppercase">{{ count }} оценок</span>
        </div>
    </div>
</template>