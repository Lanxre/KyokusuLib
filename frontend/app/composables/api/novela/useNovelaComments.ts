import type { NovelaCommentResponse } from "~/types/backend/novela";
import { $api } from "../useApi";

export function useNovelaComments() {
    const comments = useState<NovelaCommentResponse[]>('novela-comments', () => []);
    const isLoading = ref(false);

    const fetchComments = async (novelaId: number) => {
        isLoading.value = true;
        try {
            comments.value = await $api<NovelaCommentResponse[]>(`/api/novela/comments/${novelaId}`);
        } finally {
            isLoading.value = false;
        }
    };

    const addComment = async (payload: { novela_id: number; content: string; parent_id?: number | null }) => {
        return await $api('/api/novela/comments', {
            method: 'POST',
            body: payload
        });
    };

    const deleteComment = async (commentId: number) => {
        return await $api(`/api/novela/comments/${commentId}`, {
            method: 'DELETE'
        });
    };

    const updateComment = async (commentId: number, content: string) => {
        return await $api(`/api/novela/comments/${commentId}`, {
            method: 'PUT',
            body: { content, updated_at: new Date().toISOString() }
        });
    };

    return { comments, isLoading, fetchComments, addComment, deleteComment, updateComment };
}