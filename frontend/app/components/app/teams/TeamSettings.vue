<script setup lang="ts">
import { ModalWindow, Input as BaseInput, RichText as BaseRichTextEditor } from "@kyokusu-ui/vue";
import UploadIcon from "@/assets/images/special/add.png";
import { useTeamSettings } from "@/composables/api/teams/useTeamSettings";
import type { Team } from "@/types/frontend/teams";

defineOptions({
  inheritAttrs: false
});

const props = defineProps<{
    modelValue: boolean;
    team: Team;
}>();

const emit = defineEmits(["update:modelValue", "updated"]);

const {
    form,
    errors,
    isUpdating,
    fileInput,
    previewUrl,
    bannerInput,
    previewBannerUrl,
    handleImageClick,
    handleImageChange,
    removeImage,
    handleBannerClick,
    handleBannerChange,
    removeBanner,
    submitUpdate,
} = useTeamSettings(props.team);

const handleUpdate = async () => {
    const updatedTeam = await submitUpdate(props.team.slug);
    if (updatedTeam) {
        emit("updated", updatedTeam);
        emit("update:modelValue", false);
    }
};
</script>

<template>
    <ModalWindow
        v-if="modelValue"
        :model-value="modelValue"
        @update:model-value="$emit('update:modelValue', $event)"
        title="Настройки команды"
        width="w-[95%] md:w-full md:max-w-3xl"
        center-title
    >
        <div class="max-h-[85vh] overflow-y-auto pr-2 custom-scrollbar">
            <div class="space-y-8 pb-6 px-2">
                <div class="flex flex-col items-center sm:flex-row gap-6 pb-8 border-b border-zinc-100 dark:border-zinc-800">
                    <div class="relative group cursor-pointer shrink-0" @click="handleImageClick">
                        <div class="w-32 h-32 rounded-2xl overflow-hidden border-2 border-dashed border-zinc-200 dark:border-zinc-700 group-hover:border-yellow-500 transition-colors bg-zinc-50 dark:bg-zinc-900/50 flex items-center justify-center">
                            <img v-if="previewUrl" :src="previewUrl" class="w-full h-full object-cover" />
                            <div v-else class="flex flex-col items-center text-zinc-400">
                                <img :src="UploadIcon" class="w-8 h-8 dark:invert opacity-20 mb-2"/>
                                <span class="text-[10px] font-bold uppercase">Загрузить</span>
                            </div>
                        </div>
                        <input ref="fileInput" type="file" accept="image/*" class="hidden" @change="handleImageChange" />
                    </div>
                    <div class="flex flex-col items-start gap-2 w-full text-center sm:text-left">
                        <h4 class="text-sm font-bold">Аватар команды</h4>
                        <p class="text-xs text-zinc-500">JPG/PNG, до 5МБ.</p>
                        <button v-if="previewUrl" @click.stop="removeImage" class="text-xs font-bold text-red-500 uppercase mt-2 cursor-pointer">Удалить</button>
                    </div>
                </div>

                <div class="flex flex-col items-center sm:flex-row gap-6 pb-8 border-b border-zinc-100 dark:border-zinc-800">
                    <div class="relative group cursor-pointer shrink-0 w-full sm:w-64 h-32" @click="handleBannerClick">
                        <div class="w-full h-full rounded-2xl overflow-hidden border-2 border-dashed border-zinc-200 dark:border-zinc-700 group-hover:border-yellow-500 transition-colors bg-zinc-50 dark:bg-zinc-900/50 flex items-center justify-center">
                            <img v-if="previewBannerUrl" :src="previewBannerUrl" class="w-full h-full object-cover" />
                            <div v-else class="flex flex-col items-center text-zinc-400">
                                <img :src="UploadIcon" class="w-8 h-8 dark:invert opacity-20 mb-2"/>
                                <span class="text-[10px] font-bold uppercase">Загрузить баннер</span>
                            </div>
                        </div>
                        <input ref="bannerInput" type="file" accept="image/*" class="hidden" @change="handleBannerChange" />
                    </div>
                    <div class="flex flex-col items-start gap-2 w-full text-center sm:text-left">
                        <h4 class="text-sm font-bold">Баннер команды</h4>
                        <p class="text-xs text-zinc-500">JPG/PNG, до 5МБ. Рекомендуемое соотношение 3:1.</p>
                        <button v-if="previewBannerUrl" @click.stop="removeBanner" class="text-xs font-bold text-red-500 uppercase mt-2 cursor-pointer">Удалить</button>
                    </div>
                </div>

                <div class="grid grid-cols-1 gap-6 text-left">
                    <div>
                        <BaseRichTextEditor
                            id="edit-team-description"
                            label="Описание команды"
                            v-model="form.description"
                            :maxLength="2000"
                            placeholder="Расскажите о вашей команде..."
                            :disabled="isUpdating"
                            :error="errors.description"
                        />
                    </div>
                </div>

                <div class="pt-6 border-t border-zinc-100 dark:border-zinc-800">
                    <h3 class="text-base font-bold text-zinc-900 dark:text-white mb-4">Названия ролей</h3>
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <BaseInput
                            id="edit-team-owner-role"
                            label="Владелец"
                            v-model="form.owner_role_name"
                            placeholder="Например: Создатель"
                            :disabled="isUpdating"
                            :error="errors.owner_role_name"
                        />
                        <BaseInput
                            id="edit-team-moderator-role"
                            label="Модератор"
                            v-model="form.moderator_role_name"
                            placeholder="Например: Редактор"
                            :disabled="isUpdating"
                            :error="errors.moderator_role_name"
                        />
                        <BaseInput
                            id="edit-team-member-role"
                            label="Участник"
                            v-model="form.member_role_name"
                            placeholder="Например: Переводчик"
                            :disabled="isUpdating"
                            :error="errors.member_role_name"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-3 pt-6 border-t border-zinc-100 dark:border-zinc-800">
                    <button 
                        @click="handleUpdate"
                        :disabled="isUpdating"
                        class="flex-1 py-2 cursor-pointer bg-zinc-900 dark:bg-white hover:bg-zinc-700 dark:hover:bg-zinc-200 text-white dark:text-zinc-900 font-black rounded-2xl active:scale-[0.98] transition-all text-sm disabled:opacity-50 flex items-center justify-center gap-2"
                    >
                        <div v-if="isUpdating" class="w-4 h-4 border-2 border-zinc-500 border-t-transparent rounded-full animate-spin"></div>
                        {{ isUpdating ? 'Сохранение...' : 'Сохранить изменения' }}
                    </button>
                    <button 
                        @click="$emit('update:modelValue', false)"
                        :disabled="isUpdating"
                        class="md:flex-none px-10 py-2 cursor-pointer bg-zinc-50 dark:bg-zinc-800 text-zinc-500 font-bold rounded-2xl hover:bg-zinc-100 dark:hover:bg-zinc-700 transition-colors text-sm"
                    >
                        Отмена
                    </button>
                </div>
            </div>
        </div>
    </ModalWindow>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #3f3f46; border-radius: 10px; }
</style>