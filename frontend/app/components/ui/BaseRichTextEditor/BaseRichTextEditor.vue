<script lang="ts" setup>
import { ref, watch, onMounted, reactive, nextTick } from "vue";

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
const emit = defineEmits(["update:modelValue"]);

const editorRef = ref<HTMLElement | null>(null);
const charCount = ref(0);

const toolbarItems = [
	{ command: "bold", icon: "ph:text-b-bold", label: "Жирный" },
	{ command: "italic", icon: "ph:text-italic-bold", label: "Курсив" },
	{ command: "strikeThrough", icon: "ph:text-strikethrough-bold", label: "Зачеркнутый" },
];

const activeStates = reactive<Record<string, boolean>>({
	bold: false,
	italic: false,
	strikeThrough: false,
});

const syncStates = () => {
	if (import.meta.server || !editorRef.value || props.disabled) return;
	for (const item of toolbarItems) {
		activeStates[item.command] = document.queryCommandState(item.command);
	}
};

const handleInput = () => {
	if (!editorRef.value) return;
	
	const content = editorRef.value.innerHTML;
	const text = editorRef.value.innerText.trim();

	if (props.maxLength && text.length > props.maxLength) {
		editorRef.value.innerHTML = props.modelValue;
		return;
	}

	charCount.value = text.length;
	emit("update:modelValue", content === "<br>" ? "" : content);
	syncStates();
};

const exec = (command: string) => {
	if (import.meta.server) return;
	document.execCommand(command, false);
	nextTick(() => editorRef.value?.focus());
	handleInput();
};

watch(() => props.modelValue, (val) => {
	if (editorRef.value && editorRef.value.innerHTML !== val) {
		editorRef.value.innerHTML = val || "";
		charCount.value = editorRef.value.innerText.trim().length;
	}
});

onMounted(() => {
	if (editorRef.value) {
		editorRef.value.innerHTML = props.modelValue || "";
		charCount.value = editorRef.value.innerText.trim().length;
	}
});
</script>

<template>
    <div class="space-y-1.5 w-full">
        <label v-if="label" :for="id" class="text-[10px] font-black uppercase tracking-wider text-zinc-500 ml-1">
            {{ label }}
        </label>

        <div 
            class="w-full bg-zinc-50 dark:bg-zinc-900/50 border rounded-2xl overflow-hidden transition-all duration-300"
            :class="[
                error 
                    ? 'border-red-500 ring-4 ring-red-500/5' 
                    : 'border-zinc-200 dark:border-zinc-800 focus-within:border-yellow-500/50 focus-within:ring-4 focus-within:ring-yellow-500/5'
            ]"
        >
            <div 
                ref="editorRef"
                :contenteditable="!disabled"
                class="w-full min-h-[140px] p-4 text-sm text-zinc-900 dark:text-zinc-200 focus:outline-none relative z-10 antialiased"
                :class="{ 'opacity-50 cursor-not-allowed': disabled }"
                @input="handleInput"
                @keyup="syncStates"
                @mouseup="syncStates"
            ></div>

            <div class="flex items-center gap-1 p-2 bg-zinc-100/50 dark:bg-zinc-800/30 border-t border-zinc-200 dark:border-zinc-800">
                <button 
                    v-for="item in toolbarItems"
                    :key="item.command"
                    type="button"
                    @click="exec(item.command)"
                    :disabled="disabled"
                    class="p-2 rounded-xl transition-all duration-200 flex items-center justify-center cursor-pointer disabled:cursor-not-allowed"
                    :class="[
                        activeStates[item.command] 
                            ? 'bg-yellow-600/40 text-white shadow-lg shadow-yellow-500/20' 
                            : 'text-zinc-500 hover:bg-zinc-200 dark:hover:bg-zinc-700'
                    ]"
                    :title="item.label"
                >
                    <Icon :name="item.icon" size="18" />
                </button>

                <div v-if="maxLength" class="ml-auto pr-2 text-[10px] font-black uppercase tracking-widest transition-colors"
                     :class="charCount >= maxLength ? 'text-red-500' : 'text-zinc-400'">
                    {{ charCount }} / {{ maxLength }}
                </div>
            </div>
        </div>

        <p v-if="error" class="text-xs text-red-500 font-medium ml-1">{{ error }}</p>
    </div>
</template>

<style scoped>
[contenteditable]:empty:before {
    content: v-bind('placeholder ? `"${placeholder}"` : ""');
    position: absolute;
    color: #71717a;
    pointer-events: none;
    font-style: italic;
}

[contenteditable] {
    word-break: break-word;
    white-space: pre-wrap;
}
</style>