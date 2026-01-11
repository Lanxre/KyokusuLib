import { useApi } from "~/composables/api/useApi";
import type { NovelaDetails } from "@/types/backend/novela";

export function useNovela() {
    const novela = useState<NovelaDetails | null>("novela-data", () => null);
    const isLoading = useState("novela-loading", () => false);

    const fetchNovela = async (id: string) => {
        isLoading.value = true;
        try {
            const { data, error } = await useApi<NovelaDetails>(`/api/novela/${id}`);
            if (error.value) throw error.value;
            novela.value = data.value || null;
        } catch (e) {
            console.error(e);
        } finally {
            isLoading.value = false;
        }
    };

    return {
        novela,
        isLoading,
        fetchNovela,
    };
}