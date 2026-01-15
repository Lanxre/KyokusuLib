<script setup lang="ts">
import { useNovela } from "@/composables/api/novela/useNovela";
import { correctProfileImage } from "@/utils/str";

import NovelaStats from "./NovelaStats.vue";
import NovelaRating from "./NovelaRating.vue";
import ChapterList from "./ChapterList.vue";
import NovelaSettings from "./NovelaSettings.vue";

import EditIcon from "@/assets/images/special/setting.png";

import NovelaBookmarkButton from "./NovelaBookmarkButton.vue";
import NovelaLikeButton from "./NovelaLikeButton.vue";
import { useUserActivity } from "~/composables/api/profile/useUserActivity";
import { ACTIVITY_TYPES } from "~/constants/user-activity";
import { KyokusuAppRole } from "~/types/enums/role-enum";
import { useRolePermissions } from "~/composables/api/role/useRolePermissions";

const route = useRoute();

const { user } = useAuthStore();
const { novela, fetchNovela } = useNovela();
const { createUserActivity } = useUserActivity();
const { hasPermission } = useRolePermissions();


const activeTab = ref<"about" | "chapters">("about");

const novelaId = computed(() => route.params.id as string);

await useAsyncData(`novela-${novelaId.value}`, () => fetchNovela(novelaId.value));
const bookmarkInitial = ref(Boolean(novela.value?.bookmark));
const isOpenNovelaSettings = ref(false);

const totalChapters = computed(() => {
    if (!novela.value) return 0;
    return novela.value.volumes.reduce((acc, v) => acc + (v.chapters?.length || 0), 0);
});

const novelaInfo = computed(() => [
    { label: "Статус", value: novela.value?.status },
    { label: "Перевод", value: novela.value?.translation_status },
    { label: "Автор", value: novela.value?.authors?.[0]?.name, isLink: true },
]);

const updateCountBookmarks = async (categoryId: number) => {
    if (categoryId !== null && novela.value && user?.id) {
        if (!bookmarkInitial.value) {
            await createUserActivity({
                user_id: user.id,
                activity_type: ACTIVITY_TYPES.NOVELA_BOOKMARK,
                target_id: novela.value.id,
                metadata: {
                    novela_title: novela.value.title,
                    desc: 'Пользователь добавил сюжет в закладки'
                }
            });
        } else {
            await createUserActivity({
                user_id: user.id,
                activity_type: ACTIVITY_TYPES.NOVELA_BOOKMARK,
                target_id: novela.value.id,
                metadata: {
                    novela_title: novela.value.title,
                    desc: 'Пользователь обновил закладку сюжета',
                }
            });
        }
        

        novela.value.bookmark_count = (novela.value.bookmark_count || 0) + 1;
    } else if (novela.value && categoryId === null && user?.id) {
        await createUserActivity({
            user_id: user.id,
            activity_type: ACTIVITY_TYPES.NOVELA_BOOKMARK_REMOVE,
            target_id: novela.value.id,
            metadata: {
                novela_title: novela.value.title,
                desc: 'Пользователь удалил сюжет из закладок'
            }
        });
        bookmarkInitial.value = false;
        novela.value.bookmark_count = (novela.value.bookmark_count || 0) - 1;
    }
}

const updateCountLike = async (has_liked: boolean) => {
    if (novela.value && has_liked) {
        if (user?.id) {
            await createUserActivity({
                user_id: user.id,
                activity_type: ACTIVITY_TYPES.USER_NOVELA_LIKE,
                target_id: novela.value.id,
                metadata: {
                    name: novela.value.title,
                    desc: 'Пользователь поставил лайк сюжету'
                }
            })
        }
        novela.value.like_count = (novela.value.like_count || 0) + 1;
    } else if (novela.value && !has_liked) {
        if (user?.id) {
            await createUserActivity({
                user_id: user.id,
                activity_type: ACTIVITY_TYPES.USER_NOVELA_LIKE_REMOVE,
                target_id: novela.value.id,
                metadata: {
                    name: novela.value.title,
                    desc: 'Пользователь убрал лайк'
                }
            })
        }
        novela.value.like_count = (novela.value.like_count || 0) - 1;
    }
}

const updateRating = (rating: number) => {
    if (!novela.value) return;

    const oldCount = novela.value.rating_count || 0;
    const oldAverage = novela.value.rating || 0;
    const previousUserRating = novela.value.user_rating || 0;

    if (previousUserRating === rating) return;

    let newCount = oldCount;
    let newAverage = oldAverage;

    if (previousUserRating === 0) {
        newCount = oldCount + 1;
        newAverage = ((oldAverage * oldCount) + rating) / newCount;
    } else {
        const oldSum = oldAverage * oldCount;
        newAverage = (oldSum - previousUserRating + rating) / oldCount;
    }

    novela.value.rating_count = newCount;
    novela.value.rating = roundTo(newAverage, 1);
    novela.value.user_rating = rating; 

    if (user?.id) {
        createUserActivity({
            user_id: user.id,
            activity_type: ACTIVITY_TYPES.USER_NOVELA_RATING,
            target_id: novela.value.id,
            metadata: {
                name: novela.value.title,
                rating: rating,
                desc: 'Пользователь оценил произведение'
            }
        });
    }
}
</script>

<template>
    <div class="min-h-screen bg-zinc-50 dark:bg-[#0f0f0f] text-zinc-900 dark:text-zinc-200 flex flex-col font-sans transition-colors duration-300">
        
        <main class="flex-grow" v-if="novela">
            <div class="relative h-[300px] md:h-[450px] w-full overflow-hidden">
                <div class="absolute inset-0 bg-black/40 z-10 backdrop-blur-sm"></div>
                <div class="absolute inset-0 bg-gradient-to-t from-zinc-50 dark:from-[#0f0f0f] via-transparent to-transparent z-20"></div>
                <img :src="correctProfileImage(novela.poster_url || '')" class="w-full h-full object-cover blur-md scale-110" alt="Backdrop" />
            </div>

            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 -mt-64 md:-mt-106 relative z-30 pb-12">
                <div class="flex flex-col md:flex-row gap-8">
                    
                    <div class="w-full md:w-[300px] flex-shrink-0 flex flex-col gap-6">
                        <div class="relative rounded-2xl overflow-hidden shadow-2xl aspect-[2/3] border border-white/10 group">
                            <img :src="correctProfileImage(novela.poster_url || '')" class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105" :alt="novela.title" />
                            <div class="absolute top-3 left-3 bg-zinc-900/80 backdrop-blur px-2 py-1 rounded-lg text-xs font-bold text-white">
                                {{ novela.age_rating }}
                            </div>
                            <div class="absolute top-3 right-3 bg-zinc-900/80 backdrop-blur px-2 py-1 rounded-lg text-xs font-bold text-white">
                                {{ novela.type }}
                            </div>
                        </div>

                        <div class="flex flex-col gap-3">
                            <button class="w-full py-1.5 cursor-pointer rounded-xl font-bold shadow-lg transition-all active:scale-95 bg-white text-zinc-900 hover:bg-yellow-500">
                                Читать
                            </button>
                            <div class="grid grid-cols-2 gap-3">
                                <NovelaBookmarkButton 
                                    v-model="novela.bookmark" 
                                    :novela-id="novela.id" 
                                    @update:model-value="updateCountBookmarks"
                                />
                                <NovelaLikeButton 
                                    v-model="novela.has_liked"
                                    :novela-id="novela.id"
                                    @update:model-value="updateCountLike"
                                />
                            </div>
                        </div>

                        <div class="hidden md:flex flex-col gap-4 bg-white dark:bg-zinc-900/50 border border-zinc-200 dark:border-zinc-800 p-6 rounded-2xl backdrop-blur-sm">
                            <div v-for="info in novelaInfo" :key="info.label" class="flex justify-between items-center text-sm">
                                <span class="text-zinc-500">{{ info.label }}</span>
                                <span :class="['font-semibold', info.isLink ? 'text-blue-500 hover:underline cursor-pointer' : 'text-zinc-800 dark:text-zinc-200']">
                                    {{ info.value || '—' }}
                                </span>
                            </div>
                            <div class="h-px bg-zinc-200 dark:bg-zinc-800 my-1"></div>
                            <div class="flex flex-wrap gap-2">
                                <span v-for="genre in novela.genres" :key="genre" class="px-3 py-1 rounded-lg bg-zinc-100 dark:bg-zinc-800 text-[11px] font-bold uppercase tracking-wider text-zinc-600 dark:text-zinc-400 hover:bg-zinc-200 dark:hover:bg-zinc-700 cursor-pointer transition-colors">
                                    {{ genre }}
                                </span>
                            </div>
                        </div>
                    </div>

                    <div class="flex-1 min-w-0 md:pt-12">
                        <div class="md:mb-48">
                            <div class="flex flex-row">
                                <div class="flex flex-col gap-3 mb-12">
                                    <h1 class="text-4xl md:text-4xl shadow-text font-black tracking-tight leading-none">{{ novela.title }}</h1>
                                    <h2 v-if="novela.alternative_titles?.length" class="text-lg md:text-xl text-zinc-500 dark:text-zinc-400 font-medium italic">
                                        {{ novela.alternative_titles.join(' • ') }}
                                    </h2>
                                </div>
                               <div 
                                    v-if="hasPermission(KyokusuAppRole.MODERATOR)" 
                                    class="flex items-start mt-3 px-4 w-auto"
                                >
                                        <button 
                                            class="p-2 -m-4 cursor-pointer group outline-none"
                                            @click="isOpenNovelaSettings = true"
                                        >
                                            <img 
                                                :src="EditIcon" 
                                                class="w-5 h-5 md:w-4 md:h-4 dark:invert opacity-60 group-hover:opacity-100 transition-opacity"
                                            />
                                        </button>
                                </div>

                                <NovelaSettings
                                    v-model="isOpenNovelaSettings"
                                    :novela="novela"
                                />
                            </div>
                            
                            <div class="flex flex-wrap items-center gap-6 ml-2">
                                <NovelaStats :novela="novela" />
                                <NovelaRating 
                                    :novela-id="novela.id"
                                    :rating="novela.rating" 
                                    :count="novela.rating_count"
                                    :user-rating="novela.user_rating || 1"
                                    @update:rated="updateRating"
                                />
                            </div>
                        </div>

                        <div class="md:hidden mb-8 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 p-5 rounded-2xl space-y-4">
                            <div v-for="info in novelaInfo" :key="info.label" class="flex justify-between text-sm">
                                <span class="text-zinc-500">{{ info.label }}</span>
                                <span class="font-bold">{{ info.value || '—' }}</span>
                            </div>
                            <div class="flex flex-wrap gap-2">
                                <span v-for="genre in novela.genres" :key="genre" class="px-2 py-1 rounded bg-zinc-100 dark:bg-zinc-800 text-[10px] font-bold text-zinc-500">
                                    {{ genre }}
                                </span>
                            </div>
                        </div>

                        <div class="flex gap-8 border-b border-zinc-200 dark:border-zinc-800 mb-8">
                            <button 
                                v-for="tab in ['about', 'chapters']" 
                                :key="tab"
                                @click="activeTab = tab as any"
                                class="pb-4 text-lg font-bold transition-all relative cursor-pointer capitalize"
                                :class="activeTab === tab ? 'text-zinc-900 dark:text-white' : 'text-zinc-400'"
                            >
                                {{ tab === 'about' ? 'Описание' : 'Главы' }}
                                <span v-if="tab === 'chapters'" class="ml-2 text-xs opacity-50">{{ totalChapters }}</span>
                                <div v-if="activeTab === tab" class="absolute bottom-0 left-0 w-full h-1 bg-yellow-500 rounded-t-full"></div>
                            </button>
                        </div>

                        <div class="min-h-[300px]">
                            <transition name="fade" mode="out-in">
                                <div v-if="activeTab === 'about'" key="about" class="space-y-10">
                                    <div class="prose dark:prose-invert max-w-none text-zinc-700 dark:text-zinc-300 leading-relaxed text-base md:text-lg whitespace-pre-line" v-html="novela.description"></div>
                                    
                                    <div v-if="novela.categories?.length">
                                        <h3 class="text-xl font-black mb-4 uppercase tracking-tighter">Категории</h3>
                                        <div class="flex flex-wrap gap-3">
                                            <span v-for="cat in novela.categories" :key="cat" class="px-4 py-2 rounded-xl border border-zinc-200 dark:border-zinc-800 bg-white dark:bg-zinc-900/50 text-sm font-medium hover:border-yellow-500 cursor-pointer transition-all">
                                                #{{ cat }}
                                            </span>
                                        </div>
                                    </div>
                                </div>
                                <div v-else key="chapters">
                                    <ChapterList :volumes="novela.volumes" />
                                </div>
                            </transition>
                        </div>
                    </div>
                </div>
            </div>
        </main>

        <div v-else class="flex-grow flex items-center justify-center">
            <div class="w-12 h-12 border-4 border-zinc-300 border-t-yellow-500 rounded-full animate-spin"></div>
        </div>
    </div>
</template>

<style scoped>
.shadow-text {
    text-shadow: 4px 4px 2px rgba(0,0,0,0.6);
}
.fade-enter-active, .fade-leave-active {
    transition: all 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
    opacity: 0;
    transform: translateY(10px);
}
</style>