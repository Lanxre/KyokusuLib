<script lang="ts" setup>
import { toRef } from 'vue';
import { useModalLogic } from '@/composables/ui/useModal';

interface Props {
    modelValue: boolean;
    title?: string;
    width?: string;
}

const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    width: 'max-w-lg'
});

const emit = defineEmits(['update:modelValue', 'close']);

const close = () => {
    emit('update:modelValue', false);
    emit('close');
};

useModalLogic(toRef(props, 'modelValue'), close);
</script>

<template>
    <Teleport to="body">
        <Transition name="modal">
            <div v-if="modelValue" class="relative z-100" aria-labelledby="modal-title" role="dialog" aria-modal="true">
                
                <div class="fixed inset-0 bg-black/60 backdrop-blur-sm transition-opacity" @click="close"></div>

                <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
                    <div class="flex min-h-full items-center justify-center p-4 text-center sm:p-0">
                        <div 
                            class="relative transform overflow-hidden rounded-xl bg-white dark:bg-[#18181b] text-left shadow-xl transition-all sm:my-8 w-full border border-zinc-200 dark:border-zinc-800"
                            :class="width"
                            @click.stop
                        >
                            <div class="flex items-center justify-between px-6 py-4 border-b border-zinc-200 dark:border-zinc-800">
                                <h3 class="text-lg font-semibold leading-6 text-zinc-900 dark:text-white" id="modal-title">
                                    {{ title }}
                                </h3>
                                <button 
                                    @click="close"
                                    class="text-zinc-400 hover:text-zinc-500 dark:hover:text-zinc-300 transition-colors focus:outline-none"
                                >
                                    <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                                    </svg>
                                </button>
                            </div>

                            <div class="px-6 py-4">
                                <slot></slot>
                            </div>

                            <div v-if="$slots.footer" class="bg-zinc-50 dark:bg-zinc-900/50 px-6 py-3 border-t border-zinc-200 dark:border-zinc-800 flex justify-end gap-3">
                                <slot name="footer"></slot>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </Transition>
    </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
    transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}

.modal-enter-active .transform,
.modal-leave-active .transform {
    transition: all 0.3s ease-out;
}

.modal-enter-from .transform {
    opacity: 0;
    transform: translateY(10px) scale(0.95);
}

.modal-leave-to .transform {
    opacity: 0;
    transform: translateY(10px) scale(0.95);
}
</style>