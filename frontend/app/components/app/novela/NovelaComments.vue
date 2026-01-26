<script setup lang="ts">
import { ref } from "vue";
import { useNovelaComments } from "@/composables/api/novela/useNovelaComments";
import { useAuthStore } from "@/stores/auth";
import CommentItem from "./CommentItem.vue";
import AuthRequiredModal from "~/components/common/AuthRequiredModal.vue";
import BaseRichTextEditor from "~/components/ui/BaseRichTextEditor/BaseRichTextEditor.vue";
import { NOVELA_MAX_COMMENT_LENGTH } from "~/constants/data";

const props = defineProps<{ novelaId: number }>();

const { user, isAuthenticated } = useAuthStore();
const { notify } = useNotificationStore();
const { comments, fetchComments, addComment, deleteComment, isLoading } = useNovelaComments();

const isAuthModalOpen = ref(false);
const commentText = ref("");
const replyTo = ref<{ id: number; name: string } | null>(null);

const { status, refresh } = await useAsyncData(
	`comments-${props.novelaId}`,
	() => fetchComments(props.novelaId),
	{ 
        lazy: true, 
        server: true,
        watch: [() => props.novelaId] 
    }
);

const onReplyRequest = (payload: { id: number; name: string }) => {
    replyTo.value = payload;
    document.getElementById('comment-editor-section')?.scrollIntoView({ behavior: 'smooth', block: 'center' });
};

const handleSend = async () => {
    if (!isAuthenticated) {
        isAuthModalOpen.value = true;
        return;
    }
	if (!commentText.value.trim() || isLoading.value) return;

	try {
		await addComment({
			novela_id: props.novelaId,
			content: commentText.value,
			parent_id: replyTo.value?.id || null,
		});
		
		commentText.value = "";
		replyTo.value = null;
        notify({
            title: "Успех",
            content: "Комментарий успешно отправлен",
            type: "success",
        })
		await refresh(); 
	} catch (e) {
		console.error(e);
	}
};

const handleDelete = async (commentId: number) => {
    try {
        await deleteComment(commentId);
        notify({
            title: "Успех",
            content: "Комментарий успешно удален",
            type: "success",
        })
        comments.value = comments.value.filter(comment => comment.id !== commentId);
    } catch (e) {
        console.error(e);
    }
};

</script>

<template>
    <div class="space-y-12">
        <section id="comment-editor-section" class="space-y-4">
            <div v-if="replyTo" class="flex items-center justify-between bg-yellow-500/10 p-3 rounded-2xl border border-yellow-500/20">
                <span class="text-[10px] font-black text-yellow-600 uppercase tracking-widest">Ответ для: {{ replyTo.name }}</span>
                <button @click="replyTo = null" class="cursor-pointer text-yellow-600 hover:rotate-90 transition-transform">
                    <Icon name="ph:x-bold" size="14" />
                </button>
            </div>
            
            <BaseRichTextEditor
                id="comment-editor" 
                v-model="commentText"
                placeholder="Ваше мнение..."
                :disabled="isLoading"
                :maxLength="NOVELA_MAX_COMMENT_LENGTH"
            />
            
            <div class="flex justify-end">
                <button 
                    @click="handleSend"
                    :disabled="!commentText.trim() || isLoading"
                    class="px-10 py-3 bg-zinc-900 dark:bg-white text-white dark:text-zinc-900 font-black cursor-pointer rounded-2xl text-xs uppercase tracking-widest disabled:opacity-50 flex items-center gap-3 transition-all"
                >
                    <Icon v-if="isLoading" name="ph:circle-notch-bold" class="animate-spin" />
                    Отправить
                </button>
            </div>
        </section>

        <div class="relative min-h-[200px]">
            <div v-if="status === 'pending' && !comments.length" class="py-10 flex justify-center">
                <Icon name="ph:circle-notch-bold" size="32" class="animate-spin text-zinc-300" />
            </div>

            <div v-else-if="!comments.length" class="py-20 flex flex-col items-center justify-center border-2 border-dashed border-zinc-200 dark:border-zinc-800 rounded-[3rem] text-zinc-400">
                <Icon name="ph:chat-circle-dots-light" size="48" class="opacity-20 mb-4" />
                <p class="font-bold uppercase tracking-[0.2em] text-[10px]">Комментариев пока нет</p>
            </div>

            <TransitionGroup 
                v-else
                name="list" 
                tag="div" 
                class="space-y-6"
                :class="{ 'opacity-50 pointer-events-none': status === 'pending' }" 
            >
                <CommentItem 
                    v-for="comment in comments" 
                    :key="comment.id" 
                    :comment="comment"
                    :user-id="user?.id || 0"
                    @reply="onReplyRequest"
                    @delete="handleDelete"
                />
            </TransitionGroup>
        </div>

        <AuthRequiredModal v-model="isAuthModalOpen" action-text="оставлять комментарии" />
    </div>
</template>

<style scoped>
.list-enter-active, .list-leave-active { transition: all 0.5s ease; }
.list-enter-from, .list-leave-to { opacity: 0; transform: translateY(20px); }
</style>