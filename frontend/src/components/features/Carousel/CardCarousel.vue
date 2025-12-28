<script setup lang="ts">
import { ref } from "vue";
import { useDraggableScroll } from "@/composables/ui/useDraggablescroll";

interface Props {
	items: any[];
}
defineProps<Props>();

const containerRef = ref<HTMLElement | null>(null);

const { onMouseDown, onMouseLeaveOrUp, onMouseMove } =
	useDraggableScroll(containerRef);
</script>

<template>
  <div class="relative w-full group">
    <div
      ref="containerRef"
      class="flex gap-4 overflow-x-auto snap-x snap-mandatory scroll-smooth pb-4 px-4 scrollbar-hide cursor-grab select-none"
      @mousedown="onMouseDown"
      @mouseleave="onMouseLeaveOrUp"
      @mouseup="onMouseLeaveOrUp"
      @mousemove="onMouseMove"
    >
      <div
        v-for="(item, index) in items"
        :key="index"
        class="flex-shrink-0 snap-center"
      >
        <slot name="card" :item="item" />
      </div>

    </div>

  </div>
</template>

<style scoped>
/* Utility to hide scrollbar but allow scrolling */
.scrollbar-hide {
  -ms-overflow-style: none;  /* IE and Edge */
  scrollbar-width: none;  /* Firefox */
}
.scrollbar-hide::-webkit-scrollbar {
  display: none; /* Chrome, Safari and Opera */
}
</style>