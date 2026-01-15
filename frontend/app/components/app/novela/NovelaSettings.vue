<script setup lang="ts">
import ModalWindow from '~/components/features/Modal/ModalWindow.vue';
import type { NovelaDetails } from '~/types/backend/novela';

interface Props {
  modelValue: boolean;
  novela: NovelaDetails;
}

const props = defineProps<Props>();
defineEmits(['update:modelValue']);
</script>

<template>
    <ModalWindow
        v-if="modelValue"
        :model-value="modelValue"
        @update:model-value="$emit('update:modelValue', $event)"
        title="Управление сюжетом"
        width="w-[95%] md:w-full md:max-w-2xl"
        center-title
    >
        <div class="space-y-6 pb-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="flex flex-col gap-1.5">
                    <label class="text-[10px] font-bold uppercase tracking-wider text-zinc-500 ml-1">Статус сюжета</label>
                    <select class="w-full p-3 text-sm rounded-xl bg-zinc-100 dark:bg-zinc-800 border-none outline-none appearance-none cursor-pointer">
                        <option :selected="novela.status === 'ongoing'">Продолжается</option>
                        <option :selected="novela.status === 'completed'">Завершен</option>
                    </select>
                </div>
                
                <div class="flex flex-col gap-1.5">
                    <label class="text-[10px] font-bold uppercase tracking-wider text-zinc-500 ml-1">Перевод</label>
                    <select class="w-full p-3 text-sm rounded-xl bg-zinc-100 dark:bg-zinc-800 border-none outline-none appearance-none cursor-pointer">
                        <option :selected="novela.translation_status === 'ongoing'">В процессе</option>
                        <option :selected="novela.translation_status === 'completed'">Готов</option>
                    </select>
                </div>
            </div>

            <div class="p-4 rounded-2xl bg-zinc-50 dark:bg-zinc-900/50 border border-zinc-100 dark:border-zinc-800 flex justify-between items-center">
                <div class="flex flex-col">
                    <span class="text-sm font-semibold">Скрытый контент</span>
                    <span class="text-xs text-zinc-500">Произведение не будет видно в поиске</span>
                </div>
                <!-- Здесь ваш компонент BaseToggle -->
            </div>

            <div class="flex flex-col md:flex-row gap-3 pt-2">
                <button 
                    @click="$emit('update:modelValue', false)"
                    class="flex-1 py-3.5 bg-zinc-900 dark:bg-white text-white dark:text-zinc-900 font-bold rounded-xl active:scale-[0.98] transition-transform text-sm"
                >
                    Сохранить изменения
                </button>
                <button 
                    @click="$emit('update:modelValue', false)"
                    class="md:flex-none px-6 py-3.5 bg-zinc-100 dark:bg-zinc-800 text-zinc-500 font-semibold rounded-xl text-sm"
                >
                    Отмена
                </button>
            </div>
        </div>
    </ModalWindow>
</template>