import { ref } from 'vue';
import { useApi } from '@/api/api';
import { useNotificationStore } from '@/stores/notification';

export interface UserTag {
    id: number;
    tag: string;
}

export function useUserTags() {
    const { notify } = useNotificationStore();
    const availableTags = ref<UserTag[]>([]);
    const isLoadingTags = ref(false);

    const fetchTags = async () => {
        isLoadingTags.value = true;
        try {
            const { data } = await useApi('/api/user/tags', { credentials: 'include' })
                .get()
                .json<UserTag[]>();
            
            if (data.value) {
                availableTags.value = data.value;
            }
        } catch (e) {
            console.error("Failed to fetch tags", e);
        } finally {
            isLoadingTags.value = false;
        }
    };

    const updateUserTag = async (tagName: string) => {
        try {
            const { error } = await useApi('/api/profile/update', { credentials: 'include' })
                .put({ tag_name: tagName })
                .json();

            if (error.value) throw new Error("Failed to update tag");

            notify({
                title: "Успешно",
                content: "Тег обновлен",
                type: "success"
            });
            return true;
        } catch (e) {
            notify({
                title: "Ошибка",
                content: "Не удалось обновить тег",
                type: "error"
            });
            return false;
        }
    };

    return {
        availableTags,
        isLoadingTags,
        fetchTags,
        updateUserTag
    };
}