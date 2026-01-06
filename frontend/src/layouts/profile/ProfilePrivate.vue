<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { RouterLink } from 'vue-router';
import { correctProfileImage } from '@/api/utils/str';
import Separator from '@/components/ui/Separator/Separtor.vue';
import { useProfile } from '@/composables/api/profile/useProfile';
import EditIcon from "@/assets/images/special/setting.png";
import { GetUserDto } from '@/types/backend/user';
import { useUserActivity } from '@/composables/api/profile/useUserActivity';
import TagSelector from '@/components/features/TagSelector/TagSelector.vue';
import UserExperiance from '@/components/features/UserExperience/UserExperience.vue';
// const { 
//     profileData, 
//     userRoleColor, 
//     userGender, 
//     isPublicAccount, 
//     isSelfProfile,
//     accountCreated,
//     // profileTabs 
// } = useProfile();

const props = defineProps<{
    profileData: GetUserDto;
}>()

const { profileTabs, userRoleColor, isPublicAccount, isSelfProfile, getGenderText, getLastLogin, getIsLogin } = useProfile();
const { activities, fetchByUserId, isLoadingActivities } = useUserActivity();


onMounted(async () => {
    if (props.profileData) {
      await fetchByUserId(props.profileData.id);
    }
});

const activeTab = ref('overview');
</script>

<template>
    <div class="min-h-screen bg-zinc-50 dark:bg-radial-[at_center] dark:from-zinc-900 cursor-default dark:to-zinc-950 dark:to-90% text-zinc-900 dark:text-zinc-200 transition-colors duration-300 font-sans">
        
        <div class="h-48 md:h-64 w-full bg-linear-to-r from-zinc-800 to-zinc-700 relative overflow-hidden group">
            <img 
                v-if="profileData?.banner" 
                :src="correctProfileImage(profileData.banner)" 
                alt="Profile Banner"
                class="absolute inset-0 w-full h-full object-cover transition-transform duration-700"
            />
            <div class="absolute inset-0 bg-black/20"></div>
            <div class="absolute -bottom-10 -right-10 w-64 h-64 bg-white/5 rounded-full blur-3xl pointer-events-none"></div>
        </div>

        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-10">
            <div class="relative -mt-16 md:-mt-20 mb-8 flex flex-col md:flex-row items-end md:items-end gap-6">
                
                <div class="relative shrink-0">
                    <div class="w-32 h-32 md:w-40 md:h-40 rounded-2xl border-4 border-white dark:border-zinc-900 bg-zinc-800 overflow-hidden shadow-xl">
                        <img 
                            :src="correctProfileImage(profileData?.picture || '')" 
                            class="w-full h-full object-cover"
                            alt="Avatar"
                        />
                    </div>
                    <div 
                        class="absolute bottom-2 right-2 w-5 h-5 border-4 border-white dark:border-zinc-900 rounded-full transition-colors duration-300"
                        :title="getIsLogin(profileData?.last_login) ? 'Онлайн' : 'Оффлайн'"
                        :class="getIsLogin(profileData?.last_login) ? 'bg-green-500' : 'bg-red-500'"
                    ></div>
                </div>

                <div class="flex-1 pb-2 w-full md:w-auto">
                    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
                        <div>
                            <h1 class="text-3xl md:text-4xl font-bold text-zinc-900 dark:text-white flex items-center gap-3">
                                {{ profileData?.name }}
                                <span :class="['text-xs px-2 py-0.5 rounded border uppercase tracking-wider font-semibold', userRoleColor]">
                                    {{ profileData?.role || 'User' }}
                                </span>
                            </h1>
                            <div class="flex flex-row gap-4 justify-start items-center">
                                <p class="text-left ml-1 text-zinc-500 dark:text-zinc-400 font-medium mt-1">ID: #{{ profileData?.id }}</p>
                                <div class="flex flex-row gap-4">
                                    <TagSelector v-if="profileData != null && profileData.settings.is_show_tag && profileData?.active_tag" v-model="profileData.active_tag" :tags="profileData.tags" /> 
                                    <UserExperiance 
                                        v-if="profileData?.user_level" 
                                        :level="profileData.user_level.level" 
                                        :currentExp="profileData?.user_level.experience" 
                                        :expToNextLevel="profileData?.user_level.xp_needed_for_next"
                                    /> 
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <Separator class="mb-8"/>

            <div v-if="!isPublicAccount && !isSelfProfile" class="flex flex-col items-center justify-center py-20 text-center bg-white dark:bg-zinc-900/30 border border-zinc-200 dark:border-zinc-800 rounded-xl">
                <div class="w-20 h-20 bg-zinc-100 dark:bg-zinc-800 rounded-full flex items-center justify-center mb-6">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10 text-zinc-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                        <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                    </svg>
                </div>
                <h3 class="text-xl font-bold text-zinc-900 dark:text-white mb-2">Это закрытый профиль</h3>
                <p class="text-zinc-500 dark:text-zinc-400 max-w-sm">
                    Только одобренные подписчики могут видеть информацию, активность и закладки этого пользователя.
                </p>
            </div>

            <div v-else class="grid grid-cols-1 lg:grid-cols-12 gap-8">
                <div class="lg:col-span-4 xl:col-span-3 space-y-6">
                    
                    <div class="bg-white dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 shadow-sm">
                        <h3 class="text-lg font-semibold text-zinc-900 dark:text-white mb-4">О себе</h3>
                        <p class="text-sm text-zinc-600 dark:text-zinc-400 whitespace-pre-wrap wrap-break-words leading-relaxed">
                            {{ profileData?.about || 'Пользователь предпочел не рассказывать о себе.' }}
                        </p>
                        
                        <div class="mt-6 pt-6 border-t border-zinc-200 dark:border-zinc-800 space-y-3">
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500">Дата регистрации</span>
                                <span class="text-zinc-700 dark:text-zinc-300 font-medium">{{  }}</span>
                            </div>
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500">Пол</span>
                                <span class="text-zinc-700 dark:text-zinc-300 font-medium capitalize">
                                    {{  }}
                                </span>
                            </div>
                        </div>
                    </div>

                    <div class="grid grid-cols-2 gap-3">
                        <div class="bg-zinc-100 dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 p-4 rounded-xl text-center">
                            <div class="text-2xl font-bold text-zinc-900 dark:text-white">0</div>
                            <div class="text-xs text-zinc-500 uppercase tracking-wide mt-1">Прочитано</div>
                        </div>
                        <div class="bg-zinc-100 dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 p-4 rounded-xl text-center">
                            <div class="text-2xl font-bold text-zinc-900 dark:text-white">0</div>
                            <div class="text-xs text-zinc-500 uppercase tracking-wide mt-1">Отзывов</div>
                        </div>
                    </div>
                </div>

                <div class="lg:col-span-8 xl:col-span-9">
                    <div class="flex gap-1 border-b border-zinc-200 dark:border-zinc-800 mb-6 overflow-x-auto no-scrollbar">
                        <button 
                            v-for="tab in profileTabs" 
                            :key="tab.id"
                            @click="activeTab = tab.id"
                            class="px-4 py-3 text-sm font-medium border-b-2 transition-colors whitespace-nowrap"
                            :class="[
                                activeTab === tab.id 
                                    ? 'border-zinc-900 dark:border-white text-zinc-900 dark:text-white' 
                                    : 'border-transparent text-zinc-500 hover:text-zinc-700 dark:hover:text-zinc-300'
                            ]"
                        >
                            {{ tab.label }}
                        </button>
                    </div>

                    <div class="min-h-[300px]">
                        <transition name="fade" mode="out-in">
                            <div v-if="activeTab === 'overview'" key="overview" class="space-y-6">
                                <div class="bg-white dark:bg-zinc-900/30 border border-zinc-200 dark:border-zinc-800 rounded-xl p-8 flex flex-col items-center justify-center text-center">
                                    <div class="w-16 h-16 bg-zinc-100 dark:bg-zinc-800 rounded-full flex items-center justify-center mb-4">
                                        <svg class="w-8 h-8 text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                                        </svg>
                                    </div>
                                    <h4 class="text-lg font-medium text-zinc-900 dark:text-white">Активность отсутствует</h4>
                                    <p class="text-zinc-500 text-sm mt-1">Здесь будет отображаться последняя активность пользователя.</p>
                                </div>
                            </div>

                            <div v-else-if="activeTab === 'bookmarks'" key="bookmarks">
                                <div class="p-8 rounded border border-dashed border-zinc-300 dark:border-zinc-700 text-zinc-500 text-center">
                                    Список закладок пуст
                                </div>
                            </div>

                            <div v-else key="comments">
                                <div class="p-8 rounded border border-dashed border-zinc-300 dark:border-zinc-700 text-zinc-500 text-center">
                                    Нет комментариев
                                </div>
                            </div>
                        </transition>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.no-scrollbar::-webkit-scrollbar {
    display: none;
}
.no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>