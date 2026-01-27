<script setup lang="ts">
import { ref } from "vue";

interface Props {
	text: string;
	position?: "top" | "bottom" | "left" | "right";
	delay?: number;
}

const props = withDefaults(defineProps<Props>(), {
	position: "top",
	delay: 200,
});

const isVisible = ref(false);
let timeout: ReturnType<typeof setTimeout> | null = null;

const show = () => {
	timeout = setTimeout(() => {
		isVisible.value = true;
	}, props.delay);
};

const hide = () => {
	if (timeout) clearTimeout(timeout);
	isVisible.value = false;
};

const posClasses = {
	top: "bottom-full left-1/2 -translate-x-1/2 mb-2",
	bottom: "top-full left-1/2 -translate-x-1/2 mt-2",
	left: "right-full top-1/2 -translate-y-1/2 mr-2",
	right: "left-full top-1/2 -translate-y-1/2 ml-2",
};
</script>

<template>
  <div 
    class="relative inline-block w-fit" 
    @mouseenter="show" 
    @mouseleave="hide"
  >
    <slot />

    <Transition name="tooltip">
      <div 
        v-if="isVisible && text"
        class="absolute z-100 whitespace-nowrap pointer-events-none"
        :class="posClasses[position]"
      >
        <div class="bg-white dark:bg-zinc-700 text-zinc-900 dark:text-white text-[10px] font-black uppercase tracking-widest px-3 py-1.5 rounded-lg shadow-2xl border border-white/10 dark:border-black/5 antialiased">
          {{ text }}
          
          <div 
            class="absolute w-2 h-2 bg-inherit rotate-45 border-inherit"
            :class="{
              'top-full -translate-y-1/2 left-1/2 -translate-x-1/2 border-b border-r': position === 'top',
              'bottom-full translate-y-1/2 left-1/2 -translate-x-1/2 border-t border-l': position === 'bottom',
              'left-full -translate-x-1/2 top-1/2 -translate-y-1/2 border-t border-r': position === 'left',
              'right-full translate-x-1/2 top-1/2 -translate-y-1/2 border-b border-l': position === 'right',
            }"
          ></div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.tooltip-enter-active,
.tooltip-leave-active {
  transition: opacity 0.1s ease, transform 0.1s ease;
}

.tooltip-enter-from,
.tooltip-leave-to {
  opacity: 0;
  transform: translate(v-bind('position === "top" || position === "bottom" ? "-50%, 5px" : "5px, -50%"')) scale(0.85);
}
</style>