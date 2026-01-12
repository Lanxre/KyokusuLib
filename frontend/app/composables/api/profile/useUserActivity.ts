import { useApi } from "~/composables/api/useApi";
import type { CreateUserActivity, GetActivityResponse, UserActivity } from "@/types/backend/user_activity";
import type { ResponseMessage } from "~/types/backend/response_message";

export function useUserActivity() {

    const ACTIVITIES_LIMIT_QUERY = 20;
    const activities = useState<UserActivity[]>('user-activities', () => []);
    const isLoadingActivities = useState('is-loading-activities', () => false);

    const fetchActivities = async (userId?: string | number) => {
        isLoadingActivities.value = true;
        const url = userId ? `/api/user/activities/${userId}` : "/api/user/activities";
        
        try {
            const { data, error } = await useApi<GetActivityResponse>(url, { query: { limit: ACTIVITIES_LIMIT_QUERY } });
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

    const createUserActivity = async (payload: CreateUserActivity) => {
        try {
            const { data, error } = await useApi<ResponseMessage>("/api/user/activities", {
                method: "POST",
                body: payload,
            });
            if (error.value) throw error.value;
            return data.value;
        } catch (e) {
            console.error("Failed to create activity:", e);
            throw e;
        }
    };

    return {
        activities,
        isLoadingActivities,
        fetchActivities,
        createUserActivity
    };
}