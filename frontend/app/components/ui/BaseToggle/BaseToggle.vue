<script lang="ts" setup>
import { computed } from "vue";

const props = defineProps<{
	modelValue: boolean;
	label?: string;
	id?: string;
	disabled?: boolean;
}>();

const emit = defineEmits(["update:modelValue"]);

const toggle = () => {
	if (!props.disabled) {
		emit("update:modelValue", !props.modelValue);
	}
};

const switchClass = computed(() => {
	return props.modelValue
		? "bg-zinc-900 dark:bg-zinc-100"
		: "bg-zinc-200 dark:bg-zinc-800 border border-zinc-300 dark:border-zinc-700";
});

const thumbClass = computed(() => {
	return props.modelValue
		? "translate-x-[22px] bg-white dark:bg-zinc-900"
		: "translate-x-[2px] bg-zinc-400 dark:bg-zinc-500";
});
</script>

<template>
    <div class="flex items-center gap-3">
        <button 
            :id="id"
            type="button"
            role="switch"
            :aria-checked="modelValue"
            :disabled="disabled"
            @click="toggle"
            class="relative inline-flex h-6 w-12 flex-shrink-0 cursor-pointer rounded-full border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-zinc-900 dark:focus:ring-white focus:ring-offset-2 focus:ring-offset-white dark:focus:ring-offset-zinc-950 disabled:opacity-50 disabled:cursor-not-allowed"
            :class="switchClass"
        >
            <span class="sr-only">{{ label }}</span>
            <span 
                aria-hidden="true" 
                class="pointer-events-none inline-block h-5 w-5 transform rounded-full shadow ring-0 transition duration-200 ease-in-out mt-[1px]"
                :class="thumbClass"
            ></span>
        </button>
        <label 
            v-if="label" 
            :for="id" 
            class="text-sm font-medium text-zinc-700 dark:text-zinc-300 cursor-pointer select-none"
            :class="{ 'opacity-50 cursor-not-allowed': disabled }"
            @click="toggle"
        >
            {{ label }}
        </label>
    </div>
</template>