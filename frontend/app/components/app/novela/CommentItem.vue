<script setup lang="ts">
import { ref } from "vue";
import type { NovelaCommentResponse } from "@/types/backend/novela";
import { staticImage } from "@/utils/str";
import { parseDateToLocale } from "@/utils/date";
import BaseToolTip from "~/components/ui/BaseToolTip/BaseToolTip.vue";

interface Props {
	comment: NovelaCommentResponse;
	isReply?: boolean;
    userId?: number;
}

const props = defineProps<Props>();
const emit = defineEmits(["reply"]);

const isExpanded = ref(false);
const totalReplyCount = computed(() => {
    const getCount = (list: any[]) => 
        list.reduce((total, item) => total + 1 + getCount(item.replies || []), 0);
    
    return getCount(props.comment.replies || []);

});

</script>

<template>
    <div :class="['flex gap-4 group transition-all', isReply ? 'mt-4' : 'bg-white dark:bg-zinc-900/50 p-4 border border-zinc-200 dark:border-zinc-800 rounded-3xl shadow-sm']">
        <img 
            :src="staticImage(comment.user.picture)" 
            class="rounded-full bg-zinc-100 shrink-0 object-cover border border-zinc-200 dark:border-zinc-800 shadow-sm cursor-pointer" 
            :class="isReply ? 'w-8 h-8' : 'w-11 h-11'"
            @click="navigateTo(`/profile/${comment.user.id}`)"
        />
        
        <div class="flex-1 min-w-0 space-y-1.5">
            <div class="flex flex-row justify-between">
                <div class="flex items-center gap-3">
                    <span class="font-black text-zinc-900 dark:text-zinc-100 cursor-pointer text-sm" @click="navigateTo(`/profile/${comment.user.id}`)">
                        {{ comment.user.name }}
                    </span>
                    <span class="text-[9px] font-bold text-zinc-400 uppercase tracking-widest cursor-default">{{ parseDateToLocale(comment.created_at) }}</span>
                </div>
                <div class="flex">
                    <BaseToolTip v-if="comment.user.id === props.userId" text="Удалить комментарий" position="top">
                        <Icon name="ph:trash-bold" size="18" class="cursor-pointer text-red-700 hover:text-red-800 transition-colors" />
                    </BaseToolTip>
                    
                </div>
            </div>
            
            <div class="text-[14px] leading-relaxed text-zinc-700 dark:text-zinc-300 prose dark:prose-invert max-w-none break-words antialiased" v-html="comment.content"></div>
            
            <div class="flex items-center gap-4">
                <button 
                    @click="emit('reply', { id: comment.id, name: comment.user.name })" 
                    class="flex items-center gap-1.5 text-[10px] font-black uppercase tracking-widest text-zinc-400 hover:text-yellow-600 transition-colors cursor-pointer"
                >
                Ответить
                <Icon name="ph:arrow-bend-up-left-bold" size="14" />
                </button>

                <button 
                    v-if="comment.replies?.length" 
                    @click="isExpanded = !isExpanded"
                    class="flex items-center gap-1.5 text-[10px] font-black uppercase tracking-widest text-yellow-600/80 hover:text-yellow-500 transition-colors cursor-pointer"
                >
                    <Icon :name="isExpanded ? 'ph:caret-up-bold' : 'ph:caret-down-bold'" size="14" />
                    {{ isExpanded ? 'Скрыть' : `Ответы (${totalReplyCount})` }}
                </button>
            </div>

            <div 
                class="grid transition-[grid-template-rows] duration-300 ease-in-out"
                :class="isExpanded ? 'grid-template-rows-[1fr]' : 'grid-template-rows-[0fr]'"
            >
                <div class="overflow-hidden">
                    <div class="mt-2 border-l-2 border-zinc-100 dark:border-zinc-800 pl-4 space-y-2">
                        <CommentItem 
                            v-for="reply in comment.replies" 
                            :key="reply.id" 
                            :comment="reply" 
                            is-reply
                            @reply="(val) => emit('reply', val)"
                        />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.grid-template-rows-\[0fr\] { grid-template-rows: 0fr; }
.grid-template-rows-\[1fr\] { grid-template-rows: 1fr; }
</style>