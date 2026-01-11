<script lang="ts" setup>
import { ref, computed } from "vue";

const props = defineProps<{
	modelValue: string;
	id: string;
	label?: string;
	type?: string;
	placeholder?: string;
	error?: string;
	disabled?: boolean;
	autocomplete?: string;
	description?: string;
}>();

defineEmits(["update:modelValue"]);

const showPassword = ref(false);

const currentType = computed(() => {
	if (props.type === "password" && showPassword.value) {
		return "text";
	}
	return props.type || "text";
});

const togglePasswordVisibility = () => {
	showPassword.value = !showPassword.value;
};
</script>

<template>
    <div class="space-y-2">
        <label v-if="label" :for="id" class="text-sm font-medium text-dark dark:text-stone-300 ml-[3px]">
            {{ label }}
        </label>
        
        <div class="relative">
            <input 
                :id="id"
                :type="currentType"
                :value="modelValue"
                :disabled="disabled"
                :placeholder="placeholder"
                :autocomplete="autocomplete"
                @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
                class="w-full bg-zinc-200 dark:bg-zinc-900 text-zinc-900 border rounded-md py-2 px-3 text-sm text-zinc-200 placeholder-zinc-400 dark:placeholder-zinc-300 focus:outline-none focus:ring-2 transition-all disabled:opacity-50"
                :class="[
                    error 
                        ? 'border-red-500 focus:ring-red-500' 
                        : 'border-zinc-700 focus:ring-zinc-600 focus:border-transparent',
                    type === 'password' ? 'pr-10 block p-2.5 pl-10' : type === 'email' ? 'pr-10 block p-2.5 pl-10' : ''
                ]"
            />

            <div v-if="type === 'email'" class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none text-zinc-900 group-focus-within:text-white transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></svg>
            </div>

            <div v-if="type === 'password'" class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none text-zinc-900 group-focus-within:text-white transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
            </div>

            <button 
                v-if="type === 'password'"
                type="button"
                @click="togglePasswordVisibility"
                class="absolute inset-y-0 right-0 flex items-center pr-3 text-zinc-900 hover:text-white cursor-pointer transition-colors focus:outline-none"
            >
                <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"/>
                    <circle cx="12" cy="12" r="3"/>
                </svg>
                
                <svg v-else xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M9.88 9.88a3 3 0 1 0 4.24 4.24"/>
                    <path d="M10.73 5.08A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"/>
                    <path d="M6.61 6.61A13.526 13.526 0 0 0 2 12s3 7 10 7c.68 0 1.35-.06 1.99-.17"/>
                    <line x1="1" x2="23" y1="1" y2="23"/>
                </svg>
            </button>
        </div>

        <p v-if="error" class="text-xs text-red-400 ml-[3px] font-medium animate-pulse">
            {{ error }}
        </p>
        <p v-else-if="description" class="text-xs text-stone-500 ml-[3px]">
            {{ description }}
        </p>
    </div>
</template>