<script lang="ts" setup>
interface Option {
	value: string;
	label: string;
}

defineProps<{
	modelValue: string;
	label?: string;
	id: string;
	options: Option[];
	placeholder?: string;
}>();

defineEmits(["update:modelValue"]);
</script>

<template>
    <div class="space-y-2">
        <label v-if="label" :for="id" class="text-sm font-medium text-dark dark:text-stone-300 ml-[3px]">
            {{ label }}
        </label>
        <div class="relative">
            <select 
                :id="id"
                :value="modelValue"
                @change="$emit('update:modelValue', ($event.target as HTMLSelectElement).value)"
                class="w-full bg-zinc-950/50 border border-zinc-700 rounded-md py-2 px-3 text-sm text-zinc-200 darl:text-zinc-400 focus:outline-none focus:ring-2 focus:ring-zinc-600 focus:border-transparent appearance-none transition-all cursor-pointer"
            >
                <option v-if="placeholder" value="" disabled selected>{{ placeholder }}</option>
                <option 
                    v-for="option in options" 
                    :key="option.value" 
                    :value="option.value"
                    class="bg-zinc-900/70 cursor-pointer"
                >
                    {{ option.label }}
                </option>
            </select>
            
            <div class="absolute inset-y-0 right-0 flex items-center px-2 pointer-events-none">
                <svg class="w-4 h-4 text-stone-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
            </div>
        </div>
    </div>
</template>