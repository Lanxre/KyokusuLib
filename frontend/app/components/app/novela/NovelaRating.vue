<script setup lang="ts">
import { computed } from "vue";
import { useAuthStore } from "@/stores/auth";
import NovelaRatingModal from "./NovelaRatingModal.vue";
import AuthRequiredModal from "~/components/common/AuthRequiredModal.vue";

const props = defineProps<{
    novelaId: number;
	rating: number;
    userRating: number;
	count: number;
}>();

const { isAuthenticated } = useAuthStore();

const emit = defineEmits(['update:rated',]);

const isRatingModalOpen = ref(false);
const isAuthModalOpen = ref(false);

const starColor = computed(() => {
	if (props.rating >= 8) return "text-green-500";
	if (props.rating >= 6) return "text-yellow-500";
	return "text-red-500";
});

const handleOpenRating = () => {
  if (!isAuthenticated) {
    isAuthModalOpen.value = true;
    return;
  }
  isRatingModalOpen.value = !isRatingModalOpen.value;
};

const updateRating = (rating: number) => {
    emit('update:rated', rating);  
};

</script>

<template>
    <div class="flex px-1 py-1.5 rounded-lg">
        <div class="flex flex-row gap-2 leading-none">
            <span class="text-3xl font-bold text-zinc-900 dark:text-white text-right" :style="{ color: starColor }">{{ roundTo(rating, 1) }}</span>
            <div class="flex flex-col">
                <span class="text-[10px] text-zinc-400 uppercase">{{ count }} оценок</span>
                <button 
                    class="flex items-end gap-1 text-xs font-bold 
                    text-zinc-600 dark:text-white hover:text-yellow-600 
                    dark:hover:text-yellow-400 cursor-pointer"
                    @click="handleOpenRating"
                > 
                Оценить 
                <svg class="w-5 h-5 text-yellow-500 flex" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
                </button>
                <NovelaRatingModal 
                    v-model="isRatingModalOpen" 
                    :novela-id="novelaId"
                    :initial-rating="userRating"
                    @rated="updateRating"
                />
                <AuthRequiredModal 
                    v-model="isAuthModalOpen" 
                    action-text="ставить оценки" 
                />
            </div>
        </div>
    </div>
</template>