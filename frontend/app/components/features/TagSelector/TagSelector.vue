<script setup lang="ts">
import { useProfile } from "@/composables/api/profile/useProfile";
import { useUserTag } from "@/composables/api/profile/useUserTag";
import type { UserTagDTO } from "@/types/backend/user";
import { ref, onMounted, onUnmounted } from "vue";

const props = defineProps<{
	modelValue: string;
	tags: UserTagDTO[];
}>();

const emit = defineEmits<(e: "update:modelValue", value: string) => void>();

const { isSelfProfile } = useProfile();
const { updateUserTag } = useUserTag();
const isOpen = ref(false);
const containerRef = ref<HTMLElement | null>(null);

const toggleDropdown = async () => {
	isOpen.value = !isOpen.value;
};

const selectTag = async (dto: UserTagDTO) => {
	isOpen.value = false;
	emit("update:modelValue", dto.tag);

	await updateUserTag(dto);
};

const handleClickOutside = (event: MouseEvent) => {
	if (
		containerRef.value &&
		!containerRef.value.contains(event.target as Node)
	) {
		isOpen.value = false;
	}
};

onMounted(() => {
	document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
	document.removeEventListener("click", handleClickOutside);
});
</script>

<template>
    <div class="relative" ref="containerRef">
        <div 
            @click="toggleDropdown"
            class="flex items-center justify-center dark:bg-zinc-800 mt-2 px-3 py-0.5 h-8 rounded-2xl border-2 border-white dark:border-zinc-700 font-semibold cursor-pointer hover:border-zinc-500 transition-colors select-none text-sm"
        >
            <span>{{ modelValue || 'Выбрать тег' }}</span>
            <svg
                v-if="isSelfProfile"
                class="w-3 h-3 ml-2 transition-transform duration-200"
                :class="{ 'rotate-180': isOpen }"
                fill="none" viewBox="0 0 24 24" stroke="currentColor"
            >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
        </div>

        <!-- Выпадающий список -->
        <Transition name="fade">
            <div 
                v-if="isOpen && isSelfProfile"
                class="absolute left-0 mt-2 w-48 bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-700 rounded-xl shadow-xl py-1 z-50 max-h-60 overflow-y-auto custom-scrollbar"
            >
                
                <div v-if="props.tags.length === 0" class="px-2 py-2 text-xs text-zinc-500 text-center">
                    Нет тегов
                </div>

                <template v-else>
                    <div 
                        v-for="tagItem in props.tags" 
                        :key="tagItem.id"
                        @click="selectTag(tagItem)"
                        class="px-2 ml-2 mr-2 rounded-2xl text-sm text-zinc-700 dark:text-zinc-300 hover:bg-zinc-100 dark:hover:bg-zinc-800 cursor-pointer transition-colors"
                        :class="{ 'bg-zinc-100 dark:bg-zinc-500/40 font-medium': tagItem.tag === modelValue }"
                    >
                        {{ tagItem.tag }}
                    </div>
                </template>
            </div>
        </Transition>
    </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}

.custom-scrollbar::-webkit-scrollbar {
    width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
    background: #52525b;
    border-radius: 2px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: #71717a;
}
</style>