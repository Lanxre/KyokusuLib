<script setup lang="ts">
import { MultiSelect as BaseMultiSelect } from '@kyokusu-ui/vue';
import type { NovelaFilterOption } from '~/types/frontend/query/novela-filters';

defineProps<{
	options: NovelaFilterOption[];
	modelValue: string[];
}>();

const emit = defineEmits<{
	'update:modelValue': [value: string[]];
}>();
</script>

<template>
	<BaseMultiSelect
		id="filter-categories"
		label="Категории"
		:model-value="modelValue"
		:options="options"
		placeholder="Выберите категории..."
		@update:model-value="(val: string[]) => emit('update:modelValue', val)"
	/>
</template>
