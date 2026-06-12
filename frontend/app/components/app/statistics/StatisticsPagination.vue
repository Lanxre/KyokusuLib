<script setup lang="ts">
defineProps<{
  currentPage: number
  hasMore: boolean
  loading: boolean
}>()

const emit = defineEmits<{
  (e: 'prev'): void
  (e: 'next'): void
}>()
</script>

<template>
  <div class="flex items-center justify-between">
    <button
      :disabled="currentPage <= 1 || loading"
      class="px-4 py-2 rounded-xl text-sm font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed bg-zinc-100 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 text-zinc-900 dark:text-zinc-100"
      @click="emit('prev')"
    >
      < Назад
    </button>

    <span class="text-sm font-medium text-zinc-500 dark:text-zinc-400">
      Страница {{ currentPage }}
    </span>

    <button
      :disabled="!hasMore || loading"
      class="px-4 py-2 rounded-xl text-sm font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed bg-zinc-100 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 text-zinc-900 dark:text-zinc-100"
      @click="emit('next')"
    >
      Вперёд >
    </button>
  </div>
</template>
