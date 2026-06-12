<script setup lang="ts">
import { staticImage } from '@/utils/str'
import type { TotalNovelaStatistics } from '@/types/backend/statistics'

defineProps<{
  items: TotalNovelaStatistics[]
  offset: number
}>()

const roundTo = (num: number, decimals = 1) => Math.round(num * Math.pow(10, decimals)) / Math.pow(10, decimals)
</script>

<template>
  <div class="overflow-x-auto">
    <table class="w-full text-left border-collapse">
      <thead>
        <tr class="border-b border-zinc-200 dark:border-zinc-800 text-sm font-semibold text-zinc-500 dark:text-zinc-400">
          <th class="py-4 px-4 w-12 text-center">
            #
          </th>
          <th class="py-4 px-4">
            Новелла
          </th>
          <th class="py-4 px-4 text-center">
            Рейтинг
          </th>
          <th class="py-4 px-4 text-center">
            <Icon name="ph:bookmark-simple-bold" size="18" class="inline-block" />
          </th>
          <th class="py-4 px-4 text-center">
            <Icon name="ph:eye-bold" size="18" class="inline-block" />
          </th>
          <th class="py-4 px-4 text-center">
            <Icon name="ph:star-bold" size="18" class="inline-block" />
          </th>
          <th class="py-4 px-4 text-center">
            <Icon name="ph:chat-circle-bold" size="18" class="inline-block" />
          </th>
          <th class="py-4 px-4 text-center">
            <Icon name="ph:books-bold" size="18" class="inline-block" />
          </th>
          <th class="py-4 px-4 text-center">
            <Icon name="ph:file-text-bold" size="18" class="inline-block" />
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(item, index) in items"
          :key="item.novela.novelaId"
          class="border-b border-zinc-100 dark:border-zinc-800/50 hover:bg-zinc-50 dark:hover:bg-zinc-800/30 transition-colors"
        >
          <td class="py-3 px-4 text-center font-medium text-zinc-500">
            {{ offset + index + 1 }}
          </td>
          <td class="py-3 px-4">
            <NuxtLink :to="`/novela/${item.novela.novelaId}`" class="flex items-center gap-3 group">
              <img
                :src="staticImage(item.novela.posterUrl)"
                :alt="item.novela.title"
                class="w-10 h-14 rounded-lg object-cover shadow-sm group-hover:shadow-md transition-shadow"
              >
              <span class="font-semibold text-zinc-900 dark:text-zinc-100 group-hover:text-yellow-500 dark:group-hover:text-yellow-400 transition-colors line-clamp-2">
                {{ item.novela.title }}
              </span>
            </NuxtLink>
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            <div class="flex items-center justify-center gap-1">
              <Icon name="ph:star-fill" class="text-yellow-500" size="16" />
              <span>{{ roundTo(item.rating, 1) }}</span>
            </div>
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            {{ item.bookmarkCount }}
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            {{ item.readCount }}
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            {{ item.ratingCount }}
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            {{ item.commentCount }}
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            {{ item.volumeCount }}
          </td>
          <td class="py-3 px-4 text-center text-zinc-600 dark:text-zinc-400 font-medium">
            {{ item.chapterCount }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
