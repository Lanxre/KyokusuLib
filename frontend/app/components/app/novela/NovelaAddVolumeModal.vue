<script setup lang="ts">
import { ref } from "vue";
import { ModalWindow } from "@kyokusu-ui/vue";
import { useNovela } from "~/composables/api/novela/useNovela";

import type { NovelaVolume } from "@/types/backend/novela";

defineOptions({ inheritAttrs: false });

const props = defineProps<{
    modelValue: boolean;
    novelaId: number;
    volumes: NovelaVolume[];
}>();

const emit = defineEmits(["update:modelValue"]);

const { addVolume } = useNovela();
const loading = ref(false);

const volumeNumber = ref<number>(1);
const title = ref("");

const handleSubmit = async () => {
    loading.value = true;
    try {
        await addVolume(props.novelaId, volumeNumber.value, title.value);
        emit("update:modelValue", false);
        volumeNumber.value = 1;
        title.value = "";
    } finally {
        loading.value = false;
    }
};

const maxVolumeNumber = computed(() => {
  if (props.volumes === undefined) return 0;
  return props.volumes[props.volumes.length - 1]?.number! + 1;
});

</script>

<template>
    <ModalWindow
        v-if="modelValue"
        :model-value="modelValue"
        @update:model-value="$emit('update:modelValue', $event)"
        title="Добавить том"
        width="w-full max-w-md"
        center-title
    >
        <div class="py-6 space-y-4">
            <div class="space-y-2 flex flex-col gap-0.5">
                <label class="text-sm font-semibold text-zinc-700 dark:text-zinc-300 ml-1">
                        Номер тома
                </label>
                <input
                    v-model.number="maxVolumeNumber"
                    type="number"
                    :min="maxVolumeNumber"
                    class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800/50 border border-zinc-200 dark:border-zinc-700 rounded-xl text-zinc-900 dark:text-zinc-100 outline-none focus:ring-2 focus:ring-yellow-500 transition-all"
                    placeholder="Например, 1"
                />
            </div>

            <div class="space-y-2 flex flex-col gap-0.5">
                <label
                    class="text-sm font-semibold text-zinc-700 dark:text-zinc-300 ml-1">
                    Название (опционально)
                </label>
                <input
                    v-model="title"
                    type="text"
                    class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800/50 border border-zinc-200 dark:border-zinc-700 rounded-xl text-zinc-900 dark:text-zinc-100 outline-none focus:ring-2 focus:ring-yellow-500 transition-all"
                    placeholder="Например, Начало пути"
                />
            </div>

            <div class="pt-4 flex justify-end gap-3">
                <button
                    @click="$emit('update:modelValue', false)"
                    class="px-4 py-2 text-sm font-semibold text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-xl transition-colors cursor-pointer"
                >
                    Отмена
                </button>
                <button
                    @click="handleSubmit"
                    :disabled="loading || volumeNumber < 1"
                    class="px-4 py-2 text-sm font-bold bg-yellow-600 hover:bg-yellow-700 text-white rounded-xl transition-all disabled:opacity-50 flex items-center gap-2 cursor-pointer"
                >
                    <span
                        v-if="loading"
                        class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"
                    ></span>
                    Добавить
                </button>
            </div>
        </div>
    </ModalWindow>
</template>
