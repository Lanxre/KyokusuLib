import { useApi } from "@/composables/api/useApi";
import type { BookmarkCategory, Bookmark } from "~/types/frontend/bookmarks";





export function useBookmark() {
  const loading = ref(false);

  const bookmarkCategories = [
    { id: 'planned', label: 'В планах' },
    { id: 'reading', label: 'Читаю' },
    { id: 'completed', label: 'Прочитано' },
    { id: 'on_hold', label: 'Отложено' },
    { id: 'dropped', label: 'Брошено' },
  ];

  const setBookmark = async (novelaId: number, category: BookmarkCategory) => {
    loading.value = true;
    try {
      const { data, error } = await useApi("/api/novela/bookmark", {
        method: "POST",
        body: {
          novela_id: novelaId,
          category: category,
        },
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

  const removeBookmark = async (novelaId: number) => {
    loading.value = true;
    try {
      const { data, error } = await useApi(`/api/novela/bookmark/${novelaId}`, {
        method: "DELETE",
      });

      if (error.value) throw error.value;
      return data.value;
    } catch (e) {
      console.error("Failed to remove bookmark:", e);
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return {
    loading,
    bookmarkCategories,
    setBookmark,
    removeBookmark,
  };
}