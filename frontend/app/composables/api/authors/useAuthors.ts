import { ref } from "vue";
import { $api } from "@/composables/api/useApi";
import type { NovelaAuthorDetails } from "~/types/backend/novela";

export function useAuthors() {
	const isSearching = ref(false);
	const foundAuthors = ref<{ id: number; label: string }[]>([]);

	const searchAuthors = async (query: string) => {
		if (!query || query.length < 2) {
			foundAuthors.value = [];
			return;
		}

		isSearching.value = true;
		try {
			const data = await $api<NovelaAuthorDetails[]>("/api/author", {
				query: { search: query },
			});

			foundAuthors.value = data.map((a) => ({
				id: a.id,
				label: a.name,
			}));
		} catch (e) {
			foundAuthors.value = [];
		} finally {
			isSearching.value = false;
		}
	};

	return {
		isSearching,
		foundAuthors,
		searchAuthors,
	};
}
