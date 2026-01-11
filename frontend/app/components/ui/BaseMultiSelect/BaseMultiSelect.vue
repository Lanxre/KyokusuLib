<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from "vue";

interface Props {
	modelValue: string[];
	options: string[];
	id: string;
	label?: string;
	placeholder?: string;
	error?: string;
	disabled?: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:modelValue"]);

const isOpen = ref(false);
const searchQuery = ref("");
const wrapperRef = ref<HTMLElement | null>(null);
const inputRef = ref<HTMLInputElement | null>(null);

const availableOptions = computed(() => {
	return props.options.filter((option) => !props.modelValue.includes(option));
});

const filteredOptions = computed(() => {
	if (!searchQuery.value) return availableOptions.value;
	return availableOptions.value.filter((option) =>
		option.toLowerCase().includes(searchQuery.value.toLowerCase()),
	);
});

const openDropdown = () => {
	if (props.disabled) return;
	isOpen.value = true;
	nextTick(() => inputRef.value?.focus());
};

const closeDropdown = () => {
	isOpen.value = false;
	searchQuery.value = "";
};

const selectOption = (option: string) => {
	emit("update:modelValue", [...props.modelValue, option]);
	searchQuery.value = "";
	inputRef.value?.focus();
};

const removeOption = (index: number) => {
	if (props.disabled) return;
	const newSelected = [...props.modelValue];
	newSelected.splice(index, 1);
	emit("update:modelValue", newSelected);
};

const handleBackspace = () => {
	if (!searchQuery.value && props.modelValue.length > 0) {
		removeOption(props.modelValue.length - 1);
	}
};

const handleClickOutside = (event: MouseEvent) => {
	if (wrapperRef.value && !wrapperRef.value.contains(event.target as Node)) {
		closeDropdown();
	}
};

onMounted(() => document.addEventListener("click", handleClickOutside));
onUnmounted(() => document.removeEventListener("click", handleClickOutside));
</script>

<template>
    <div class="space-y-1 relative" ref="wrapperRef">
        <label v-if="label" :for="id" class="text-sm font-medium text-zinc-700 dark:text-zinc-300 ml-[3px]">
            {{ label }}
        </label>

        <div 
            class="w-full bg-zinc-50 dark:bg-zinc-900/50 border rounded-lg mt-1 py-1.5 px-2 min-h-8 flex flex-wrap items-center gap-2 transition-all cursor-text relative"
            :class="[
                error 
                    ? 'border-red-500 focus-within:ring-2 focus-within:ring-red-500' 
                    : 'border-zinc-300 dark:border-zinc-700 focus-within:ring-2 focus-within:ring-zinc-500 dark:focus-within:ring-zinc-600 focus-within:border-transparent',
                disabled ? 'opacity-50 cursor-not-allowed' : ''
            ]"
            @click="openDropdown"
        >
            <div 
                v-for="(item, index) in modelValue" 
                :key="item"
                class="flex h-8 items-center gap-2 cursor-default bg-zinc-200 dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 px-4 py-1 rounded-4xl text-[15px] group transition-colors hover:bg-red-100 dark:hover:bg-orange-700/30"
            >
                <span>{{ item }}</span>
                <button 
                    type="button" 
                    @click.stop="removeOption(index)"
                    :disabled="disabled"
                    class="text-zinc-500 cursor-pointer hover:text-red-500 transition-colors focus:outline-none"
                >
                    <svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>

            <input 
                :id="id"
                ref="inputRef"
                v-model="searchQuery"
                type="text"
                class="flex-1 bg-transparent border-none outline-none text-sm text-zinc-900 dark:text-zinc-200 placeholder-zinc-400 min-w-[100px] h-7"
                :placeholder="modelValue.length === 0 ? placeholder : ''"
                :disabled="disabled"
                @keydown.backspace="handleBackspace"
                @focus="isOpen = true"
            />
        </div>

        <div 
            v-if="isOpen && (filteredOptions.length > 0 || searchQuery)" 
            class="absolute z-50 w-full mt-1 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-700 rounded-lg shadow-xl max-h-60 overflow-hidden flex flex-col"
        >
            <ul class="overflow-y-auto flex-1 p-1 custom-scrollbar">
                <li 
                    v-for="option in filteredOptions" 
                    :key="option"
                    @click="selectOption(option)"
                    class="px-3 py-2 text-sm rounded-md cursor-pointer transition-colors text-zinc-700 dark:text-zinc-300 hover:bg-zinc-100 dark:hover:bg-zinc-800"
                >
                    {{ option }}
                </li>
                <li v-if="filteredOptions.length === 0" class="px-3 py-2 text-sm text-zinc-400 text-center">
                    Ничего не найдено
                </li>
            </ul>
        </div>

        <p v-if="error" class="text-xs text-red-500 dark:text-red-400 ml-[3px]">{{ error }}</p>
    </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
    width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
    background: #52525b;
    border-radius: 2px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: #71717a;
}
</style>