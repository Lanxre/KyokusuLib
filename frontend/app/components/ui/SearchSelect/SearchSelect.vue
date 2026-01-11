<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted } from "vue";

interface Props {
	modelValue: string;
	id: string;
	label?: string;
	selects: string[];
	placeholder?: string;
	error?: string;
	disabled?: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:modelValue"]);

const isOpen = ref(false);
const searchQuery = ref("");
const wrapperRef = ref<HTMLElement | null>(null);

const filteredSearchItems = computed(() => {
	if (!searchQuery.value) return props.selects;
	return props.selects.filter((select) =>
		select.toLowerCase().includes(searchQuery.value.toLowerCase()),
	);
});

const toggleDropdown = () => {
	if (props.disabled) return;
	isOpen.value = !isOpen.value;
	if (isOpen.value) {
		searchQuery.value = "";
	}
};

const selectItem = (item: string) => {
	emit("update:modelValue", item);
	isOpen.value = false;
	searchQuery.value = "";
};

const handleClickOutside = (event: MouseEvent) => {
	if (wrapperRef.value && !wrapperRef.value.contains(event.target as Node)) {
		isOpen.value = false;
	}
};

onMounted(() => {
	document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
	document.removeEventListener("click", handleClickOutside);
});
</script>

<template>
    <div class="space-y-2 relative" ref="wrapperRef">
        <label v-if="label" :for="id" class="text-sm font-medium text-zinc-700 dark:text-zinc-300 ml-[3px]">
            {{ label }} 
        </label>
        
        <div class="relative mt-1">
            <div 
                @click="toggleDropdown"
                class="w-full bg-zinc-50 dark:bg-zinc-900/50 border rounded-lg py-2.5 px-3 text-sm text-left flex items-center justify-between cursor-pointer transition-all"
                :class="[
                    error 
                        ? 'border-red-500 focus:ring-red-500' 
                        : 'border-zinc-300 dark:border-zinc-700 hover:border-zinc-400 dark:hover:border-zinc-600',
                    disabled ? 'opacity-50 cursor-not-allowed' : ''
                ]"
            >
                <span :class="modelValue ? 'text-zinc-900 dark:text-zinc-200' : 'text-zinc-400 dark:text-zinc-500'">
                    {{ modelValue || placeholder || 'Выберите страну' }}
                </span>
                <svg 
                    class="w-4 h-4 text-zinc-500 transition-transform duration-200"
                    :class="{ 'rotate-180': isOpen }"
                    fill="none" viewBox="0 0 24 24" stroke="currentColor"
                >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
            </div>

            <div 
                v-if="isOpen" 
                class="absolute z-50 w-full mt-1 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-700 rounded-lg shadow-lg max-h-60 overflow-hidden flex flex-col"
            >
                <div class="p-2 border-b border-zinc-100 dark:border-zinc-800">
                    <input 
                        v-model="searchQuery"
                        type="text"
                        class="w-full bg-zinc-100 dark:bg-zinc-800 border-none rounded-md py-1.5 px-3 text-sm text-zinc-900 dark:text-zinc-200 placeholder-zinc-400 focus:ring-2 focus:ring-zinc-500 dark:focus:ring-zinc-600 focus:outline-none"
                        placeholder="Поиск..."
                        @click.stop
                    />
                </div>
                
                <ul class="overflow-y-auto flex-1 p-1">
                    <li 
                        v-for="item in filteredSearchItems" 
                        :key="item"
                        @click="selectItem(item)"
                        class="px-3 py-2 text-sm rounded-md cursor-pointer transition-colors"
                        :class="[
                            modelValue === item 
                                ? 'bg-zinc-100 dark:bg-zinc-800 text-zinc-900 dark:text-white font-medium' 
                                : 'text-zinc-700 dark:text-zinc-300 hover:bg-zinc-50 dark:hover:bg-zinc-800/50'
                        ]"
                    >
                        {{ item }}
                    </li>
                    <li v-if="filteredSearchItems.length === 0" class="px-3 py-2 text-sm text-zinc-400 text-center">
                        Ничего не найдено
                    </li>
                </ul>
            </div>
        </div>

        <p v-if="error" class="text-xs text-red-500 dark:text-red-400 ml-[3px]">{{ error }}</p>
    </div>
</template>