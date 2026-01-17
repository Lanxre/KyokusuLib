import { useApi, $api } from "~/composables/api/useApi";
import type { NovelaDetails } from "@/types/backend/novela";
import { useNotificationStore } from "~/stores/notification";
import type { NovelsQueryParams } from "~/types/frontend/query/novela-query";

export function useNovela() {
	const novels = useState<NovelaDetails[]>("novels-data", () => []);
	const novela = useState<NovelaDetails | null>("novela-data", () => null);
	const { notify } = useNotificationStore();
	const isLoading = useState("novela-loading", () => false);
	const isUpdating = useState("novela-updating", () => false);

	const fetchNovels = async (params: NovelsQueryParams) => {
		isLoading.value = true;
        try {
            const { data, error } = await useApi<NovelaDetails[]>('/api/novela', {
                params: params
            });
            
            if (error.value) throw error.value;
			novels.value = data.value!;
            return data.value;
        } catch (e) {
            console.error(e);
        } finally {
            isLoading.value = false;
        }
	}

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
		novels,
		isLoading,
		isUpdating,
		fetchNovela,
		fetchNovels,
		updateNovela,
	};
}