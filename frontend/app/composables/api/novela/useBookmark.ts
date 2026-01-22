import { $api } from "@/composables/api/useApi";
import type { NovelaDetails } from "~/types/backend/novela";
import type { BookmarkCategory } from "~/types/frontend/bookmarks";

export function useBookmark() {
  const userBookmarkNovels = useState<NovelaDetails[]>("user-category-bookmarks", () => []);
  const loading = ref(false);

  const bookmarkCategories = [
    { id: 'planned', label: 'В планах', icon: "ph:book-open-bold" },
    { id: 'reading', label: 'Читаю', icon: "ph:calendar-blank-bold" },
    { id: 'completed', label: 'Прочитано', icon: "ph:check-circle-bold" },
    { id: 'on_hold', label: 'Отложено', icon: "ph:pause-circle-bold" },
    { id: 'dropped', label: 'Брошено', icon: "ph:prohibit-bold" },
  ];

  const setBookmark = async (novelaId: number, category: BookmarkCategory) => {
    loading.value = true;
    try {
      const data = await $api("/api/novela/bookmark", {
        method: "POST",
        body: {
          novela_id: novelaId,
          category: category,
        },
      });

      return data;
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
      const data = await $api(`/api/novela/bookmark/${novelaId}`, {
        method: "DELETE",
      });

      return data;
    } catch (e) {
      console.error("Failed to remove bookmark:", e);
      throw e;
    } finally {
      loading.value = false;
    }
  };

  const fetchUserBookmarkNovels = async (userID: number, category: BookmarkCategory) => {
		loading.value = true;
		try {
			const data = await $api<NovelaDetails[]>("/api/novela/bookmarks", {
				query: {  user_id: userID, category: category },
			});

      userBookmarkNovels.value = data || [];

		} catch (e) {
			userBookmarkNovels.value = [];
		} finally {
			loading.value = false;
		}
	};

  return {
    loading,
    bookmarkCategories,
    setBookmark,
    removeBookmark,
    fetchUserBookmarkNovels,
    userBookmarkNovels
  };
}