<script setup lang="ts">
interface Props {
	label: string;
	for?: string;
}

const props = defineProps<Props>();
</script>

<template>
    <label v-if="props.label" :for="props.for" class="text-[10px] font-black uppercase tracking-wider text-zinc-500 ml-1">
        {{ props.label }}
    </label>
</template>