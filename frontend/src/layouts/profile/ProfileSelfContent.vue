<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { correctProfileImage } from '@/api/utils/str';
import Separator from '@/components/ui/Separator/Separtor.vue';
import EditIcon from "@/assets/images/special/setting.png";
import { useProfile } from '@/composables/api/profile/useProfile';
import { useUserActivity } from '@/composables/api/profile/useUserActivity';
import { useActivityStore } from '@/stores/activity';
import TabActivity from '@/layouts/profile/activity/TabActivity.vue';
import TagSelector from '@/components/features/TagSelector/TagSelector.vue';
import { useInterfaceSettings } from '@/composables/api/settings/useInterfaceSettings';

const { profileData, accountCreated, profileTabs, userRoleColor, userGender, lastLogin } = useProfile();
const { activities, fetchActivities, isLoadingActivities } = useUserActivity();
const { isUserActive } = useActivityStore();
const { isShowTag } = useInterfaceSettings();
const activeTab = ref('overview');

onMounted(async () => {
    await fetchActivities();
});

</script>

<template>
    <div class="min-h-screen bg-zinc-50 dark:bg-radial-[at_center] dark:from-zinc-900 cursor-default dark:to-zinc-950 dark:to-90% text-zinc-900 dark:text-zinc-200 transition-colors duration-300 font-sans">
        
       <!-- Banner -->
        <div class="h-48 md:h-64 w-full bg-linear-to-r from-zinc-800 to-zinc-700 relative overflow-hidden group">
            <!-- Изображение баннера -->
            <img 
                v-if="profileData?.banner" 
                :src="correctProfileImage(profileData.banner)" 
                alt="Profile Banner"
                class="absolute inset-0 w-full h-full transition-transform duration-700"
            />

            <!-- Затемнение (Overlay) -->
            <div class="absolute inset-0 bg-black/20"></div>
            
            <!-- Декоративный элемент -->
            <div class="absolute -bottom-10 -right-10 w-64 h-64 bg-white/5 rounded-full blur-3xl pointer-events-none"></div>
        </div>

        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-10">
            <!-- Profile Header -->
            <div class="relative -mt-16 md:-mt-20 mb-8 flex flex-col md:flex-row items-end md:items-end gap-6">
                
                <!-- Avatar -->
                <div class="relative shrink-0">
                    <div class="w-32 h-32 md:w-40 md:h-40 rounded-2xl border-4 border-white dark:border-zinc-900 bg-zinc-800 overflow-hidden shadow-xl">
                        <img 
                            :src="correctProfileImage(profileData!.picture)" 
                            class="w-full h-full object-cover"
                            alt="Avatar"
                        />
                    </div>
                    <!-- Online Status -->
                    <div 
                        class="absolute bottom-2 right-2 w-5 h-5 border-4 border-white dark:border-zinc-900 rounded-full"
                        title="Онлайн"
                        :class="{ 'bg-green-500' : isUserActive, 'bg-red-500' : !isUserActive }"
                    ></div>
                </div>

                <!-- Main Info -->
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
                                    <TagSelector v-if="isShowTag && profileData?.active_tag" v-model="profileData.active_tag" :tags="profileData.tags" />
                                    <div class="flex dark:bg-zinc-800 px-2 mt-1 rounded-2xl border-2 border-white dark:border-zinc-700 font-semibold">Опыт пользователя: 0 </div>
                                </div>
                            </div>
                        </div>

                        <div class="flex gap-3 w-full md:w-auto">
                            <RouterLink to="/settings" class="flex items-center gap-2 px-4 py-2 bg-zinc-200 dark:bg-zinc-800 hover:bg-zinc-300 dark:hover:bg-zinc-700 rounded-lg text-sm font-medium transition-colors border border-transparent dark:border-zinc-700">
                                <img :src="EditIcon" class="w-4 h-4 dark:invert opacity-70"/>
                                <span>Редактировать</span>
                            </RouterLink>
                        </div>
                    </div>
                </div>
            </div>

            <Separator class="mb-8"/>

            <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
                <!-- Left Column (Sidebar) -->
                <div class="lg:col-span-4 xl:col-span-3 space-y-6">
                    
                    <!-- About Card -->
                    <div class="bg-white dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 shadow-sm">
                        <h3 class="text-lg font-semibold text-zinc-900 dark:text-white mb-4">О себе</h3>
                        
                        <p class="text-sm text-zinc-600 dark:text-zinc-400 whitespace-pre-wrap wrap-break-words leading-relaxed">
                            {{ profileData?.about || 'Пользователь предпочел не рассказывать о себе.' }}
                        </p>
                        
                        <div class="mt-6 pt-6 border-t border-zinc-200 dark:border-zinc-800 space-y-3">
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500">Дата регистрации</span>
                                <span class="text-zinc-700 dark:text-zinc-300 font-medium">{{ accountCreated }}</span>
                            </div>
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500">Последняя активность</span>
                                <span class="text-zinc-700 dark:text-zinc-300 font-medium">{{ lastLogin }}</span>
                            </div>
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500">Пол</span>
                                <span class="text-zinc-700 dark:text-zinc-300 font-medium capitalize">
                                    {{ userGender }}
                                </span>
                            </div>
                        </div>
                    </div>

                    <!-- Mini Stats -->
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

                <!-- Right Column (Content) -->
                <div class="lg:col-span-8 xl:col-span-9">
                    
                    <!-- Tabs Navigation -->
                    <div class="flex gap-1 border-b border-zinc-200 dark:border-zinc-800 mb-6 overflow-x-auto no-scrollbar">
                        <button 
                            v-for="tab in profileTabs" 
                            :key="tab.id"
                            @click="activeTab = tab.id"
                            class="px-4 py-3 text-sm font-medium border-b-2 transition-colors whitespace-nowrap cursor-pointer"
                            :class="[
                                activeTab === tab.id 
                                    ? 'border-zinc-900 dark:border-white text-zinc-900 dark:text-white' 
                                    : 'border-transparent text-zinc-500 hover:text-zinc-700 dark:hover:text-zinc-300'
                            ]"
                        >
                            {{ tab.label }}
                        </button>
                    </div>

                    <!-- Content Area -->
                    <div class="min-h-[300px]">
                        <transition name="fade" mode="out-in">
                            <div v-if="activeTab === 'overview'" key="overview" class="space-y-6">
                                <TabActivity
                                    :activities="activities"
                                    :isLoading="isLoadingActivities"
                                />
                            </div>

                            <div v-else-if="activeTab === 'bookmarks'" key="bookmarks">
                                <div class="p-4 rounded border border-dashed border-zinc-300 dark:border-zinc-700 text-zinc-500 text-center">
                                    Список закладок пуст
                                </div>
                            </div>

                            <div v-else key="comments">
                                <div class="p-4 rounded border border-dashed border-zinc-300 dark:border-zinc-700 text-zinc-500 text-center">
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