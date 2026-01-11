import { useNotificationStore } from "~/stores/notification";
import { useApi } from "~/composables/api/useApi";
import type { UserTagDTO } from "~/types/backend/user";

export function useUserTag() {
	const { notify } = useNotificationStore();

	const updateUserTag = async (dto: UserTagDTO) => {
		try {
			const { data, error } = await useApi<any>("/api/user/tag", {
				method: "PUT",
				body: dto,
			});

			if (error.value) {
				throw new Error(error.value.message);
			}

			if (data.value?.message) {
				notify({
					type: "success",
					title: "User tag updated successfully",
					content: data.value.message,
				});
			}
		} catch (e: any) {
			console.error(e);
		}
	};

	return {
		updateUserTag,
	};
}
