<script setup lang="ts">
import { ref } from "vue";
import { onClickOutside } from "@vueuse/core";
import SearchSelect from "~/components/ui/SearchSelect/SearchSelect.vue";
import { REPORT_COMMNET_REASONS } from "~/constants/data";

const props = defineProps<{ commentId: number }>();
const emit = defineEmits(["close", "submit"]);

const target = ref(null);
const selectedReason = ref<string | null>(null);
const error = ref(false);

onClickOutside(target, () => emit("close"));

const handleSend = () => {
    if (!selectedReason.value) {
        error.value = true;
        return;
    }
    emit("submit", { commentId: props.commentId, reason: selectedReason.value });
    emit("close");
};
</script>

<template>
    <div ref="target" class="absolute right-0 top-8 z-50 w-64 bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-2xl p-4 animate-in fade-in zoom-in duration-200">
        <div class="space-y-4">
            <div class="flex flex-col gap-1">
                <span class="text-[10px] font-black uppercase tracking-widest text-zinc-400 ml-1">Причина жалобы</span>
                <SearchSelect
                    id="report-reason"
                    v-model="selectedReason"
                    :selects="REPORT_COMMNET_REASONS"
                    placeholder="Выберите причину"
                    :error="error && !selectedReason ? 'Обязательно выберите причину' : ''"
                />
            </div>

            <div class="flex flex-col gap-2">
                <button 
                    @click="handleSend"
                    class="w-full py-1.5 bg-red-600 hover:bg-red-700 text-white text-[10px] font-black uppercase tracking-widest rounded-xl transition-colors cursor-pointer"
                >
                    Отправить репорт
                </button>
                <button 
                    @click="emit('close')"
                    class="w-full py-1.5 bg-zinc-100 dark:bg-zinc-800 text-zinc-500 text-xs font-bold uppercase rounded-xl hover:bg-zinc-200 dark:hover:bg-zinc-700 transition-colors cursor-pointer"
                >
                    Отмена
                </button>
            </div>
        </div>
    </div>
</template>