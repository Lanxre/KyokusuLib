import type { GetUserDto } from "@/types/backend/user";
import { useApi } from "~/composables/api/useApi";

export function useUserApi() {
	const getUser = async (userId: number): Promise<GetUserDto | null> => {
		if (!userId) return null;

		try {
			const { data, error } = await useApi<GetUserDto>(`/api/user/${userId}`);

			if (error.value) {
				console.error("USER FETCHER ERROR:", error.value);
				return null;
			}
			return data.value!;
		} catch (e) {
			console.error("USER FETCHER EXCEPTION:", e);
			return null;
		}
	};

	return {
		getUser,
	};
}
