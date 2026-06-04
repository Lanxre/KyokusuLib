import { useApi, $api } from "~/composables/api/useApi";
import type { NovelaDetails } from "@/types/backend/novela";
import { useNotificationStore } from "~/stores/notification";
import type { NovelsQueryParams } from "~/types/frontend/query/novela-query";

export function useNovela() {
	const novela = useState<NovelaDetails | null>("novela-data", () => null);
	const { notify } = useNotificationStore();
	const isLoading = useState("novela-loading", () => false);
	const isUpdating = useState("novela-updating", () => false);

	const fetchNovels = async (params: NovelsQueryParams) => {
		isLoading.value = true;
		try {
			const { data, error } = await useApi<NovelaDetails[]>("/api/novela", {
				params: params,
			});

			if (error.value) throw error.value;
			return data.value;
		} catch (e) {
			console.error(e);
		} finally {
			isLoading.value = false;
		}
	};

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

	const addVolume = async (novelaId: number, volumeNumber: number, title: string) => {
		try {
			const res = await $api<{message: string}>(`/api/novela/${novelaId}/volumes`, {
				method: "POST",
				body: { volume_number: volumeNumber, title },
			});
			notify({
				type: "success",
				title: "Успех",
				content: res.message || "Том успешно добавлен",
			});
			await fetchNovela(novelaId);
		} catch (e) {
			console.error(e);
		}
	};

	const addChapter = async (novelaId: number, volumeId: number, chapterNumber: number, title: string, content: string) => {
		try {
			const res = await $api<{id: string; message: string; status: string}>(`/api/novela/volumes/${volumeId}/chapters`, {
				method: "POST",
				body: { chapter_number: chapterNumber, title, content },
			});
			notify({
				type: "success",
				title: "Успех",
				content: res.message || "Глава успешно добавлена",
			});
			await fetchNovela(novelaId);
			return res;
		} catch (e) {
			console.error(e);
		}
	};

	const addChapterImage = async (chapterId: string, imageUrl: string, caption: string, position: number) => {
		try {
			await $api(`/api/novela/chapters/${chapterId}/images`, {
				method: "POST",
				body: { image_url: imageUrl, caption, position },
			});
		} catch (e) {
			console.error(e);
		}
	};

	return {
		novela,
		isLoading,
		isUpdating,
		fetchNovela,
		fetchNovels,
		updateNovela,
		addVolume,
		addChapter,
		addChapterImage,
	};
}
