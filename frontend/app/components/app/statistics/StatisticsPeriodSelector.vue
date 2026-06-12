<script setup lang="ts">
import { StatisticsPeriodEnum } from '@/types/enums/statistics'

defineProps<{
  modelValue: StatisticsPeriodEnum | undefined
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: StatisticsPeriodEnum | undefined): void
}>()

const periods = [
  { label: 'За всё время', value: undefined },
  { label: 'За день', value: StatisticsPeriodEnum.Day },
  { label: 'За неделю', value: StatisticsPeriodEnum.Week },
  { label: 'За месяц', value: StatisticsPeriodEnum.Month },
]
</script>

<template>
  <div class="flex flex-wrap gap-2">
    <button
      v-for="period in periods"
      :key="String(period.value)"
      type="button"
      class="px-4 py-2 rounded-xl text-sm font-medium transition-colors"
      :class="[
        modelValue === period.value
          ? 'bg-zinc-200 dark:bg-zinc-700 text-zinc-900 dark:text-white ring-1 ring-zinc-300 dark:ring-zinc-600'
          : 'text-zinc-500 hover:bg-zinc-100 dark:text-zinc-400 dark:hover:bg-zinc-800'
      ]"
      @click="emit('update:modelValue', period.value)"
    >
      {{ period.label }}
    </button>
  </div>
</template>
