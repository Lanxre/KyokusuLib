<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from "vue";

interface Option {
	id: string | number;
	label: string;
}

interface Props {
	modelValue: (string | number | Option)[];
	options: (string | Option)[];
	id: string;
	label?: string;
	placeholder?: string;
	error?: string;
	disabled?: boolean;
	loading?: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:modelValue", "search"]);

const isOpen = ref(false);
const searchQuery = ref("");
const wrapperRef = ref<HTMLElement | null>(null);
const inputRef = ref<HTMLInputElement | null>(null);

const normalizedOptions = computed(() => {
	return props.options.map((opt) =>
		typeof opt === "string" ? { id: opt, label: opt } : opt,
	);
});

const normalizedSelected = computed(() => {
	return props.modelValue.map((val) => {
		if (typeof val === "object" && val !== null) return val as Option;

		const found = normalizedOptions.value.find((opt) => opt.id === val);
		return found || { id: val, label: String(val) };
	});
});

const availableOptions = computed(() => {
	const selectedIds = normalizedSelected.value.map((s) => s.id);
	return normalizedOptions.value.filter((opt) => !selectedIds.includes(opt.id));
});

const filteredOptions = computed(() => {
	if (!searchQuery.value) return availableOptions.value;
	return availableOptions.value.filter((opt) =>
		opt.label.toLowerCase().includes(searchQuery.value.toLowerCase()),
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

const selectOption = (option: Option) => {
    if (props.modelValue.some(item => (typeof item === 'object' ? item.id : item) === option.id)) {
        return;
    }
    emit("update:modelValue", [...props.modelValue, option]);
    searchQuery.value = "";
    nextTick(() => inputRef.value?.focus());
};

const removeOption = (index: number) => {
	if (props.disabled) return;
	const newSelected = [...props.modelValue];
	newSelected.splice(index, 1);
	emit("update:modelValue", newSelected);
};

const handleInput = (e: Event) => {
	const target = e.target as HTMLInputElement;
	emit("search", target.value);
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

onMounted(() => document.addEventListener("mousedown", handleClickOutside));
onUnmounted(() =>
	document.removeEventListener("mousedown", handleClickOutside),
);
</script>

<template>
    <div class="space-y-1 relative w-full" ref="wrapperRef">
        <label v-if="label" :for="id" class="text-[10px] font-black uppercase tracking-wider text-zinc-500 ml-1">
            {{ label }}
        </label>

        <div 
            class="w-full bg-zinc-50 dark:bg-zinc-900/50 border rounded-xl mt-1.5 py-2 px-3 min-h-[46px] flex flex-wrap items-center gap-2 transition-all cursor-text relative"
            :class="[
                error 
                    ? 'border-red-500' 
                    : 'border-zinc-200 dark:border-zinc-800 focus-within:border-yellow-500/50 focus-within:ring-4 focus-within:ring-yellow-500/5',
                disabled ? 'opacity-50 cursor-not-allowed' : ''
            ]"
            @click="openDropdown"
        >
            <div 
                v-for="(item, index) in normalizedSelected" 
                :key="item.id"
                class="flex h-7 items-center gap-2 cursor-default bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 pl-3 pr-1.5 rounded-lg text-xs font-bold border border-zinc-200 dark:border-zinc-700 shadow-sm"
            >
                <span>{{ item.label }}</span>
                <button 
                    type="button" 
                    @click.stop="removeOption(index)"
                    class="text-zinc-400 hover:text-red-500 transition-colors cursor-pointer"
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
                class="flex-1 bg-transparent border-none outline-none text-sm text-zinc-900 dark:text-zinc-200 placeholder-zinc-400 min-w-[120px] h-7"
                :placeholder="modelValue.length === 0 ? placeholder : ''"
                :disabled="disabled"
                @input="handleInput"
                @keydown.backspace="handleBackspace"
                @focus="isOpen = true"
            />

            <div v-if="loading" class="absolute right-3">
                <div class="w-4 h-4 border-2 border-yellow-500 border-t-transparent rounded-full animate-spin"></div>
            </div>
        </div>

        <div v-if="isOpen" class="absolute z-[100] w-full mt-2 bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-2xl overflow-hidden flex flex-col">
            <ul class="overflow-y-auto max-h-60 p-2 custom-scrollbar">
                <li 
                    v-for="option in filteredOptions" 
                    :key="option.id"
                    @click.stop="selectOption(option)"
                    class="px-4 py-2.5 text-sm rounded-xl cursor-pointer transition-all text-zinc-700 dark:text-zinc-300 hover:bg-zinc-100 dark:hover:bg-zinc-800 font-medium"
                >
                    {{ option.label }}
                </li>
                <li v-if="filteredOptions.length === 0 && !loading" class="px-4 py-8 text-sm text-zinc-400 text-center italic">
                    {{ searchQuery ? 'Ничего не найдено' : 'Начните вводить...' }}
                </li>
            </ul>
        </div>
    </div>
</template>