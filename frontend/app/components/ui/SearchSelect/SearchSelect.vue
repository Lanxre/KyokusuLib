<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from "vue";

interface Props {
	modelValue: string | number | null;
	id: string;
	label?: string;
	selects: any[];
	placeholder?: string;
	error?: string;
	disabled?: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:modelValue"]);

const isOpen = ref(false);
const searchQuery = ref("");
const wrapperRef = ref<HTMLElement | null>(null);
const searchInputRef = ref<HTMLInputElement | null>(null);

const normalizedOptions = computed(() => {
	return props.selects.map(item => 
		typeof item === 'string' ? { id: item, label: item } : item
	);
});

const currentLabel = computed(() => {
	const found = normalizedOptions.value.find(opt => opt.id === props.modelValue);
	return found ? found.label : '';
});

const filteredSearchItems = computed(() => {
	if (!searchQuery.value) return normalizedOptions.value;
	return normalizedOptions.value.filter((item) =>
		item.label.toLowerCase().includes(searchQuery.value.toLowerCase()),
	);
});

const toggleDropdown = () => {
	if (props.disabled) return;
	isOpen.value = !isOpen.value;
	if (isOpen.value) {
		searchQuery.value = "";
		nextTick(() => searchInputRef.value?.focus());
	}
};

const selectItem = (item: any) => {
	emit("update:modelValue", item.id);
	isOpen.value = false;
	searchQuery.value = "";
};

const handleOutsideClick = (event: MouseEvent) => {
	if (wrapperRef.value && !wrapperRef.value.contains(event.target as Node)) {
		isOpen.value = false;
	}
};

onMounted(() => {
	document.addEventListener("mousedown", handleOutsideClick);
});

onUnmounted(() => {
	document.removeEventListener("mousedown", handleOutsideClick);
});
</script>

<template>
    <div class="space-y-1 relative w-full" ref="wrapperRef">
        <label v-if="label" :for="id" class="text-[10px] font-black uppercase tracking-wider text-zinc-500 ml-1">
            {{ label }} 
        </label>
        
        <div class="relative">
            <div 
                @click="toggleDropdown"
                class="w-full bg-zinc-50 dark:bg-zinc-900/50 border rounded-xl py-3 px-4 text-sm text-left flex items-center justify-between cursor-pointer transition-all"
                :class="[
                    error 
                        ? 'border-red-500' 
                        : 'border-zinc-200 dark:border-zinc-800 hover:border-zinc-300 dark:hover:border-zinc-700',
                    isOpen ? 'border-yellow-500/50 ring-4 ring-yellow-500/5' : '',
                    disabled ? 'opacity-50 cursor-not-allowed' : ''
                ]"
            >
                <span :class="currentLabel ? 'text-zinc-900 dark:text-zinc-200 font-default' : 'text-zinc-400 dark:text-zinc-500'">
                    {{ currentLabel || placeholder || 'Выберите из списка' }}
                </span>
                <svg 
                    class="w-4 h-4 text-zinc-400 transition-transform duration-300"
                    :class="{ 'rotate-180 text-yellow-500': isOpen }"
                    fill="none" viewBox="0 0 24 24" stroke="currentColor"
                >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
            </div>

            <Transition name="dropdown">
                <div 
                    v-if="isOpen" 
                    class="absolute z-[100] w-full mt-2 bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-2xl overflow-hidden flex flex-col"
                >
                    <div class="p-2 border-b border-zinc-100 dark:border-zinc-800">
                        <div class="relative">
                            <input 
                                ref="searchInputRef"
                                v-model="searchQuery"
                                type="text"
                                class="w-full bg-zinc-100 dark:bg-zinc-800/50 border-none rounded-xl py-2 pl-9 pr-4 text-sm text-zinc-900 dark:text-zinc-200 placeholder-zinc-500 focus:ring-2 focus:ring-yellow-500/20 outline-none"
                                placeholder="Поиск..."
                                @click.stop
                            />
                            <svg class="absolute left-3 top-2.5 h-4 w-4 text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                            </svg>
                        </div>
                    </div>
                    
                    <ul class="overflow-y-auto max-h-64 p-1.5 custom-scrollbar">
                        <li 
                            v-for="item in filteredSearchItems" 
                            :key="item.id"
                            @click.stop="selectItem(item)"
                            class="px-3 py-2.5 text-sm rounded-xl cursor-pointer transition-all flex items-center justify-between group"
                            :class="[
                                modelValue === item.id 
                                    ? 'bg-yellow-50 dark:bg-yellow-500/10 text-yellow-600 dark:text-yellow-500 font-bold' 
                                    : 'text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 hover:text-zinc-900 dark:hover:text-zinc-200'
                            ]"
                        >
                            {{ item.label }}
                            <div v-if="modelValue === item.id" class="w-1.5 h-1.5 rounded-full bg-yellow-500"></div>
                        </li>
                        <li v-if="filteredSearchItems.length === 0" class="px-4 py-8 text-sm text-zinc-400 text-center italic">
                            Ничего не найдено
                        </li>
                    </ul>
                </div>
            </Transition>
        </div>

        <p v-if="error" class="text-xs text-red-500 font-medium ml-1 mt-1">{{ error }}</p>
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
    background: #3f3f46;
    border-radius: 10px;
}

.dropdown-enter-active, .dropdown-leave-active {
    transition: all 0.2s ease;
}
.dropdown-enter-from, .dropdown-leave-to {
    opacity: 0;
    transform: translateY(-8px);
}
</style>