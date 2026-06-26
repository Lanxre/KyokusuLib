<script setup lang="ts">
import {
	STATUS_OPTIONS,
	TRANSLATION_STATUS_OPTIONS,
	NOVELA_TYPE_OPTIONS,
} from '~/types/frontend/query/novela-filters';
import { Select } from '@kyokusu-ui/vue';

const status = defineModel<string>('status', { default: '' });
const translationStatus = defineModel<string>('translationStatus', { default: '' });
const type = defineModel<string>('type', { default: '' });
</script>

<template>
	<div class="space-y-3">
		<div class="space-y-1.5">
			<label class="text-xs font-bold uppercase tracking-wider text-zinc-500 dark:text-zinc-400 ml-0.5">
				Тип
			</label>
			<Select
				v-model="type"
				:options="NOVELA_TYPE_OPTIONS"
				placeholder="Все типы"
			/>
		</div>

		<div class="space-y-1.5">
			<label class="text-xs font-bold uppercase tracking-wider text-zinc-500 dark:text-zinc-400 ml-0.5">
				Статус выхода
			</label>
			<Select
				v-model="status"
				:options="STATUS_OPTIONS"
				placeholder="Все статусы"
			/>
		</div>

		<div class="space-y-1.5">
			<label class="text-xs font-bold uppercase tracking-wider text-zinc-500 dark:text-zinc-400 ml-0.5">
				Статус перевода
			</label>
			<Select
				v-model="translationStatus"
				:options="TRANSLATION_STATUS_OPTIONS"
				placeholder="Все статусы"
			/>
		</div>
	</div>
</template>
