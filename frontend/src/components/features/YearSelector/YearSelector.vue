<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';

import { YEAR_MIN, YEAR_MAX } from '@/const';

interface Props {
    modelValue: string | number;
    id: string;
    label?: string;
    placeholder?: string;
    error?: string;
    disabled?: boolean;
    from?: number;
    to?: number;
}


const props = withDefaults(defineProps<Props>(), {
    from: YEAR_MIN,
    to: YEAR_MAX
});

const emit = defineEmits(['update:modelValue']);

const isOpen = ref(false);
const searchQuery = ref('');
const wrapperRef = ref<HTMLElement | null>(null);
const searchInputRef = ref<HTMLInputElement | null>(null);

const allYears = computed(() => {
    const years = [];
    for (let i = props.to; i >= props.from; i--) {
        years.push(i.toString());
    }
    return years;
});

const filteredYears = computed(() => {
    if (!searchQuery.value) return allYears.value;
    return allYears.value.filter(year => year.includes(searchQuery.value));
});

const toggleDropdown = () => {
    if (props.disabled) return;
    isOpen.value = !isOpen.value;
    if (isOpen.value) {
        searchQuery.value = '';
        setTimeout(() => searchInputRef.value?.focus(), 50);
    }
};

const selectYear = (year: string) => {
    emit('update:modelValue', year);
    isOpen.value = false;
    searchQuery.value = '';
};

const handleClickOutside = (event: MouseEvent) => {
    if (wrapperRef.value && !wrapperRef.value.contains(event.target as Node)) {
        isOpen.value = false;
    }
};

onMounted(() => document.addEventListener('click', handleClickOutside));
onUnmounted(() => document.removeEventListener('click', handleClickOutside));
</script>

<template>
    <div class="space-y-2 relative" ref="wrapperRef">
        <label v-if="label" :for="id" class="text-sm font-medium text-zinc-700 dark:text-zinc-300 ml-[3px]">
            {{ label }}
        </label>
        
        <div class="relative mt-1">
            <!-- Триггер -->
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
                    {{ modelValue || placeholder || 'Выберите год' }}
                </span>
                
                <svg 
                    class="w-4 h-4 text-zinc-500 transition-transform duration-200"
                    :class="{ 'rotate-180': isOpen }"
                    fill="none" viewBox="0 0 24 24" stroke="currentColor"
                >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
            </div>

            <!-- Выпадающий список -->
            <div 
                v-if="isOpen" 
                class="absolute z-50 w-full mt-1 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-700 rounded-lg shadow-xl max-h-60 overflow-hidden flex flex-col"
            >
                <!-- Поле поиска -->
                <div class="p-2 border-b border-zinc-100 dark:border-zinc-800 sticky top-0 bg-white dark:bg-zinc-900 z-10">
                    <input 
                        ref="searchInputRef"
                        v-model="searchQuery"
                        type="text"
                        class="w-full bg-zinc-100 dark:bg-zinc-800 border-none rounded-md py-1.5 px-3 text-sm text-zinc-900 dark:text-zinc-200 placeholder-zinc-400 focus:ring-2 focus:ring-zinc-500 dark:focus:ring-zinc-600 focus:outline-none"
                        placeholder="Найти год..."
                        @click.stop
                    />
                </div>
                
                <!-- Список -->
                <ul class="overflow-y-auto flex-1 p-1 custom-scrollbar">
                    <li 
                        v-for="year in filteredYears" 
                        :key="year"
                        @click="selectYear(year)"
                        class="px-3 py-2 text-sm rounded-md cursor-pointer transition-colors"
                        :class="[
                            modelValue.toString() === year 
                                ? 'bg-zinc-100 dark:bg-zinc-800 text-zinc-900 dark:text-white font-medium' 
                                : 'text-zinc-700 dark:text-zinc-300 hover:bg-zinc-50 dark:hover:bg-zinc-800/50'
                        ]"
                    >
                        {{ year }}
                    </li>
                    <li v-if="filteredYears.length === 0" class="px-3 py-2 text-sm text-zinc-400 text-center">
                        Год не найден
                    </li>
                </ul>
            </div>
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