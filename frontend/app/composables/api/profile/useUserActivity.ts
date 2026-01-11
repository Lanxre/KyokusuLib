import { useApi } from "~/composables/api/useApi";
import type { GetActivityResponse, UserActivity } from "@/types/backend/user_activity";

export function useUserActivity() {
    const activities = useState<UserActivity[]>('user-activities', () => []);
    const isLoadingActivities = useState('is-loading-activities', () => false);

    const fetchActivities = async (userId?: string | number) => {
        isLoadingActivities.value = true;
        const url = userId ? `/api/user/activities/${userId}` : "/api/user/activities";
        
        try {
            const { data, error } = await useApi<GetActivityResponse>(url);
            if (error.value) throw error.value;
            
            activities.value = data.value?.message || [];
            return activities.value;
        } catch (e) {
            console.error("Failed to fetch activities:", e);
            activities.value = [];
        } finally {
            isLoadingActivities.value = false;
        }
    };

	const fetchByUserId = async (userId: number) => {
        isLoadingActivities.value = true;
        try {
            const { data } = await useApi<GetActivityResponse>(`/api/user/activities/${userId}`)
            activities.value = data.value?.message || [];
			return activities.value;
        } catch (e: any) {
            console.error(e);
            activities.value = [];
        } finally {
            isLoadingActivities.value = false;
        }
    };

    return {
        activities,
        isLoadingActivities,
        fetchActivities,
		fetchByUserId,
    };
}