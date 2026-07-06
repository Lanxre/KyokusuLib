<script setup lang="ts">
import { Button } from "@kyokusu-ui/vue";

defineProps<{
	saving: boolean;
}>();

const emit = defineEmits<{
	save: [];
	cancel: [];
}>();
</script>

<template>
	<div class="flex items-center justify-end gap-3">
		<Button
			type="button"
			label="Отмена"
			class-name="px-4 py-2 rounded-lg text-sm font-medium text-zinc-600 dark:text-zinc-400 bg-zinc-100 dark:bg-zinc-800 hover:bg-zinc-200 dark:hover:bg-zinc-700 transition cursor-pointer"
			@click="emit('cancel')"
		/>
		<Button
			type="submit"
			label="Сохранить"
			:disabled="saving"
			class-name="px-4 py-2 rounded-lg text-sm font-medium text-white bg-amber-500 hover:bg-amber-600 disabled:opacity-50 disabled:cursor-not-allowed transition cursor-pointer"
			@click="emit('save')"
		/>
	</div>
</template>
