<script setup lang="ts">
import { ref } from "vue";
import { useAutoDragCarousel } from "@/composables/ui/useAutoDragCarousel";

interface Props {
	items: any[];
}
defineProps<Props>();

const containerRef = ref<HTMLElement | null>(null);

const { onMouseDown, onMouseLeaveOrUp, onMouseMove, stopAutoPlay } =
	useAutoDragCarousel(containerRef, 5);
</script>

<template>
  <div class="relative w-full px-4 md:px-8 overflow-hidden">
    <div
      ref="containerRef"
      class="
        flex items-center gap-3 md:gap-6 overflow-x-auto 
        snap-x snap-mandatory 
        scroll-smooth 
        scrollbar-hide 
        cursor-grab select-none
        h-full py-4
      "
      @mousedown="onMouseDown"
      @mouseleave="onMouseLeaveOrUp"
      @mouseup="onMouseLeaveOrUp"
      @mousemove="onMouseMove"
      @mouseenter="stopAutoPlay" 
      @touchstart="stopAutoPlay"
    >
      <div
        v-for="(item, index) in items"
        :key="index"
        class="flex-shrink-0 snap-start"
      >
        <slot name="card" :item="item" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>