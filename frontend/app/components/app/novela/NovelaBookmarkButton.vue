<script setup lang="ts">
import { ref, computed } from 'vue';
import { onClickOutside } from '@vueuse/core';
import { useBookmark } from '@/composables/api/novela/useBookmark';

interface Props {
  modelValue: string | undefined;
  novelaId: number;
}

const props = defineProps<Props>();
const emit = defineEmits(['update:modelValue']);

const { setBookmark, removeBookmark, loading, bookmarkCategories } = useBookmark();

const isPopoverOpen = ref(false);
const containerRef = ref(null);

onClickOutside(containerRef, () => {
  isPopoverOpen.value = false;
});

const currentLabel = computed(() => {
  if (!props.modelValue) return 'В закладки';
  return bookmarkCategories.find(c => c.id === props.modelValue)?.label || 'В закладки';
});

const handleSelect = async (categoryId: string) => {
  try {
    await setBookmark(props.novelaId, categoryId as any);
    emit('update:modelValue', categoryId);
    isPopoverOpen.value = false;
  } catch (e) {
    console.error(e);
  }
};

const handleRemove = async () => {
  try {
    await removeBookmark(props.novelaId);
    emit('update:modelValue', null);
    isPopoverOpen.value = false;
  } catch (e) {
    console.error(e);
  }
};
</script>

<template>
  <div ref="containerRef" class="relative inline-block w-full">
    <button 
      @click="isPopoverOpen = !isPopoverOpen"
      type="button"
      :disabled="loading"
      class="w-full py-2.5 cursor-pointer rounded-xl bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 hover:bg-zinc-50 dark:hover:bg-zinc-800 transition-all flex items-center justify-center px-2 gap-2 text-sm font-semibold group"
      :class="{ 'border-yellow-500/50 ring-2 ring-yellow-500/10': modelValue }"
    >
      <svg 
        class="w-5 h-5 transition-colors" 
        :class="modelValue ? 'text-yellow-500' : 'text-zinc-400'"
        fill="none" viewBox="0 0 24 24" stroke="currentColor"
      >
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
      </svg>
      <span class="truncate">{{ currentLabel }}</span>
      
      <svg class="w-4 h-4  opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
      </svg>
    </button>

    <Transition name="popover">
      <div 
        v-if="isPopoverOpen"
        class="absolute left-0 right-0 mt-2 p-1.5 z-50 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl shadow-xl overflow-hidden"
      >
        <div class="flex flex-col gap-0.5">
          <button
            v-for="cat in bookmarkCategories"
            :key="cat.id"
            @click="handleSelect(cat.id)"
            class="w-full py-2 rounded-lg text-sm font-medium flex items-center justify-center gap-2 cursor-pointer"
            :class="modelValue === cat.id 
              ? 'bg-yellow-50 dark:bg-yellow-500/10 text-yellow-600 dark:text-yellow-500' 
              : 'hover:bg-zinc-100 dark:hover:bg-zinc-800 text-zinc-600 dark:text-zinc-400'"
          >
            {{ cat.label }}
            <div v-if="modelValue === cat.id" class="w-1.5 h-1.5 rounded-full bg-yellow-500"></div>
          </button>

          <div v-if="modelValue" class="h-px bg-zinc-100 dark:bg-zinc-800 my-1"></div>

          <button
            v-if="modelValue"
            @click="handleRemove"
            class="w-full flex justify-center rounded-lg text-sm font-medium text-red-500 hover:bg-red-50 dark:hover:bg-red-500/10 transition-colors cursor-pointer"
          >
            Удалить из закладок
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.popover-enter-active, .popover-leave-active {
  transition: all 0.2s ease;
}
.popover-enter-from, .popover-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>