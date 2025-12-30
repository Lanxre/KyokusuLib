<script lang="ts" setup>
import { ref, watch, onMounted, reactive } from 'vue';

interface Props {
    modelValue: string;
    id: string;
    label?: string;
    maxLength?: number;
    error?: string;
    disabled?: boolean;
    placeholder?: string;
}

const props = defineProps<Props>();
const emit = defineEmits(['update:modelValue']);

const editorRef = ref<HTMLElement | null>(null);
const charCount = ref(0);

const activeStates = reactive({
    bold: false,
    italic: false,
    strikeThrough: false
});

const checkActiveStates = () => {
    if (!editorRef.value || props.disabled) return;
    activeStates.bold = document.queryCommandState('bold');
    activeStates.italic = document.queryCommandState('italic');
    activeStates.strikeThrough = document.queryCommandState('strikeThrough');
};

const updateContent = () => {
    if (!editorRef.value) return;
    const content = editorRef.value.innerHTML;
    const text = editorRef.value.innerText || '';
    
    if (props.maxLength && text.length > props.maxLength) {
        editorRef.value.innerHTML = props.modelValue; 
        return;
    }
    
    charCount.value = text.length;
    emit('update:modelValue', content);
    checkActiveStates();
};

const execCommand = (command: string) => {
    document.execCommand(command, false);
    updateContent();
    editorRef.value?.focus();
    checkActiveStates();
};

watch(() => props.modelValue, (newVal) => {
    if (editorRef.value && editorRef.value.innerHTML !== newVal) {
        editorRef.value.innerHTML = newVal;
        charCount.value = editorRef.value.innerText.length;
    }
});

onMounted(() => {
    if (editorRef.value) {
        editorRef.value.innerHTML = props.modelValue;
        charCount.value = editorRef.value.innerText.length;
    }
});
</script>

<template>
    <div class="space-y-2">
        <div class="flex justify-between items-end ml-[3px]">
            <label v-if="label" :for="id" class="text-sm font-medium text-zinc-700 dark:text-zinc-300">
                {{ label }}
            </label>
        </div>

        <div 
            class="w-full bg-zinc-50 dark:bg-zinc-900/50 border rounded-lg overflow-hidden transition-colors"
            :class="[
                error 
                    ? 'border-red-500' 
                    : 'border-zinc-300 dark:border-zinc-700 focus-within:ring-2 focus-within:ring-zinc-500 dark:focus-within:ring-zinc-600 focus-within:border-transparent',
                disabled ? 'opacity-50 cursor-not-allowed' : ''
            ]"
        >
            <div 
                ref="editorRef"
                :contenteditable="!disabled"
                class="w-full min-h-[150px] p-3 text-sm text-zinc-900 dark:text-zinc-200 focus:outline-none empty:before:content-[attr(data-placeholder)] empty:before:text-zinc-400"
                :data-placeholder="placeholder"
                @input="updateContent"
                @keyup="checkActiveStates"
                @mouseup="checkActiveStates"
            ></div>
            <div class="flex items-center gap-1 p-2 border-t border-zinc-200 dark:border-zinc-800 bg-zinc-100 dark:bg-zinc-900">
                <button 
                    type="button"
                    @click="execCommand('bold')"
                    class="p-1.5 rounded text-zinc-700 dark:text-zinc-300 transition-colors"
                    :class="[
                        activeStates.bold 
                            ? 'bg-zinc-300 dark:bg-zinc-700' 
                            : 'hover:bg-zinc-200 dark:hover:bg-zinc-800'
                    ]"
                    title="Жирный"
                    :disabled="disabled"
                >
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 4h8a4 4 0 014 4 4 4 0 01-4 4H6z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 12h9a4 4 0 014 4 4 4 0 01-4 4H6z" />
                    </svg>
                </button>
                <button 
                    type="button"
                    @click="execCommand('italic')"
                    class="p-1.5 rounded text-zinc-700 dark:text-zinc-300 transition-colors"
                    :class="[
                        activeStates.italic 
                            ? 'bg-zinc-300 dark:bg-zinc-700' 
                            : 'hover:bg-zinc-200 dark:hover:bg-zinc-800'
                    ]"
                    title="Курсив"
                    :disabled="disabled"
                >
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <line x1="19" y1="4" x2="10" y2="4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        <line x1="14" y1="20" x2="5" y2="20" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        <line x1="15" y1="4" x2="9" y2="20" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>
                </button>
                <button 
                    type="button"
                    @click="execCommand('strikeThrough')"
                    class="p-1.5 rounded text-zinc-700 dark:text-zinc-300 transition-colors"
                    :class="[
                        activeStates.strikeThrough 
                            ? 'bg-zinc-300 dark:bg-zinc-700' 
                            : 'hover:bg-zinc-200 dark:hover:bg-zinc-800'
                    ]"
                    title="Зачеркнутый"
                    :disabled="disabled"
                >
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15h16" opacity="0" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 19c0-1.657-2.686-3-6-3s-6 1.343-6 3" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.5 5c0 1.657 2.686 3 6 3s6-1.343 6-3" />
                    </svg>
                </button>
            </div>
        </div>

        <div class="flex justify-between px-1">
            <p v-if="error" class="text-xs text-red-500 dark:text-red-400">{{ error }}</p>
            <div v-else></div>
            
            <p v-if="maxLength" class="text-xs text-zinc-400 ml-auto" :class="{ 'text-red-500': charCount >= maxLength }">
                {{ charCount }}/{{ maxLength }}
            </p>
        </div>
    </div>
</template>