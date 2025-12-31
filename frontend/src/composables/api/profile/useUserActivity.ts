import { onMounted, ref } from 'vue';
import { useApi } from '@/api/api';
import type { GetActivityResponse, UserActivity } from '@/types/backend/user_activity';

export function useUserActivity() {
    const activities = ref<UserActivity[]>([]);
    const isLoadingActivities = ref(false);
    const error = ref<string | null>(null);

    const fetchActivities = async () => {
        isLoadingActivities.value = true;
        error.value = null;

        try {
            const { data } = await useApi('/api/user/activities', { 
                credentials: 'include' 
            })
            .get()
            .json<GetActivityResponse>();
            
            activities.value = data.value?.message || [];
          
        } catch (e: any) {
            console.error(e);
            error.value = e.message || "Failed to fetch activities";
            activities.value = [];
        } finally {
            isLoadingActivities.value = false;
        }
    };

    onMounted(async () => {
        await fetchActivities();
    });

    return {
        activities,
        isLoadingActivities,
        error,
        fetchActivities
    };
}