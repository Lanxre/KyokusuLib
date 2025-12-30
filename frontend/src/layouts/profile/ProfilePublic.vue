<script setup lang="ts">
import { ref } from 'vue';
import { correctProfileImage } from '@/api/utils/str';
import Separator from '@/components/ui/Separator/Separtor.vue';
import { useProfile } from '@/composables/api/profile/useProfile';
import { useActivityStore } from '@/stores/activity';

const { profileData, profileTabs, userRoleColor, userGender } = useProfile();
const { isUserActive } = useActivityStore();
const activeTab = ref('overview');

</script>

<template>
    <div class="min-h-screen bg-zinc-50 dark:bg-radial-[at_center] dark:from-zinc-900 cursor-default dark:to-zinc-950 dark:to-90% text-zinc-900 dark:text-zinc-200 transition-colors duration-300 font-sans">
        
         <!-- Banner -->
        <div class="h-48 md:h-64 w-full bg-gradient-to-r from-zinc-800 to-zinc-700 relative overflow-hidden group">
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
                            :src="correctProfileImage(profileData?.picture || '')" 
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
                                <div class="flex mt-2 dark:bg-zinc-800 px-2 rounded-2xl border-2 border-white dark:border-zinc-700 font-semibold">Опыт пользователя: 0 </div>
                            </div>
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
                        
                        <p class="text-sm text-zinc-600 dark:text-zinc-400 whitespace-pre-wrap break-words leading-relaxed">
                            {{ profileData?.about || 'Пользователь предпочел не рассказывать о себе.' }}
                        </p>
                        
                        <div class="mt-6 pt-6 border-t border-zinc-200 dark:border-zinc-800 space-y-3">
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