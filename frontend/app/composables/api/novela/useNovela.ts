import { useApi, $api } from "~/composables/api/useApi";
import type { NovelaDetails } from "@/types/backend/novela";
import { useNotificationStore } from "~/stores/notification";

export function useNovela() {
	const novela = useState<NovelaDetails | null>("novela-data", () => null);
	const { notify } = useNotificationStore();
	const isLoading = useState("novela-loading", () => false);
	const isUpdating = useState("novela-updating", () => false);

	const fetchNovela = async (id: string | number) => {
		isLoading.value = true;
		try {
			const { data, error } = await useApi<NovelaDetails>(`/api/novela/${id}`);
			if (error.value) throw error.value;
			novela.value = data.value!;
			return data.value;
		} finally {
			isLoading.value = false;
		}
	};

	const updateNovela = async (id: number, payload: any, file?: File | null) => {
		isUpdating.value = true;
		const formData = new FormData();

		formData.append("data", JSON.stringify(payload));
		if (file) formData.append("poster", file);

		try {
			await $api(`/api/novela/${id}`, {
				method: "PUT",
				body: formData,
			});

			notify({
				type: "success",
				title: "Новела обновлена",
				content: "Новела успешно обновлена",
			});

		} finally {
			isUpdating.value = false;
		}
	};

	return {
		novela,
		isLoading,
		isUpdating,
		fetchNovela,
		updateNovela,
	};
}