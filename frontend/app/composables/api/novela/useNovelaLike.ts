import { useApi } from "@/composables/api/useApi";
import type { NovelaLikeRequest } from "~/types/frontend/novela-like";

export function useNovelaLike() {
  const loading = ref(false);

  const setNovelaLike = async (payload: NovelaLikeRequest) => {
    loading.value = true;
    try {
      const { data, error } = await useApi("/api/novela/like", {
        method: "POST",
        body: payload,
      });

      if (error.value) throw error.value;
      return data.value;
    } catch (e) {
      console.error("Failed to set bookmark:", e);
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return {
    loading,
    setNovelaLike,
  };
}