import { ref, watch } from 'vue';
import { useStorage } from '@vueuse/core';
import { useDebounceFn } from '@vueuse/core';
import { $api } from '@/composables/api/useApi';
import { useAuthors } from '@/composables/api/authors/useAuthors';
import type { NovelaDetails } from '@/types/backend/novela';

export type SearchCategory = 'ranobe' | 'users' | 'teams' | 'authors';

export function useSearch() {
    const isOpen = useState('search-is-open', () => false);
    const searchQuery = useState('search-query', () => '');
    const activeCategory = useState<SearchCategory>('search-category', () => 'ranobe');
    const isSearching = useState('search-is-searching', () => false);
    const searchResults = useState<any[]>('search-results', () => []);
    
    const recentSearches = useStorage<string[]>('recent-searches', []);
    const popularSearches = ref<string[]>(['Система', 'Магия', 'Перерождение', 'Академия']);
    
    const { searchAuthors, foundAuthors } = useAuthors();

    const openSearch = () => {
        isOpen.value = true;
    };

    const closeSearch = () => {
        isOpen.value = false;
        searchQuery.value = '';
        searchResults.value = [];
    };

    const addRecentSearch = (query: string) => {
        if (!query.trim()) return;
        const index = recentSearches.value.indexOf(query);
        if (index > -1) {
            recentSearches.value.splice(index, 1);
        }
        recentSearches.value.unshift(query);
        if (recentSearches.value.length > 5) {
            recentSearches.value.pop();
        }
    };

    const removeRecentSearch = (query: string) => {
        const index = recentSearches.value.indexOf(query);
        if (index > -1) {
            recentSearches.value.splice(index, 1);
        }
    };

    const clearRecentSearches = () => {
        recentSearches.value = [];
    };

    const performSearch = async () => {
        const query = searchQuery.value.trim();
        if (!query || query.length < 2) {
            searchResults.value = [];
            return;
        }

        isSearching.value = true;
        try {
            if (activeCategory.value === 'ranobe') {
                const data = await $api<NovelaDetails[]>("/api/novela", {
                    query: { search: query, limit: 10 }
                });
                searchResults.value = data || [];
            } else if (activeCategory.value === 'authors') {
                await searchAuthors(query);
                searchResults.value = foundAuthors.value;
            } else if (activeCategory.value === 'users') {
                const data = await $api<any[]>("/api/user", {
                    query: { search: query, limit: 10 }
                });
                searchResults.value = data || [];
            } else if (activeCategory.value === 'teams') {
                const data = await $api<any[]>("/api/teams", {
                    query: { search: query, limit: 10 }
                });
                searchResults.value = data || [];
            }
        } catch (error) {
            console.error('Search failed:', error);
            searchResults.value = [];
        } finally {
            isSearching.value = false;
        }
    };

    const debouncedSearch = useDebounceFn(() => {
        performSearch();
    }, 500);

    watch([searchQuery, activeCategory], () => {
        if (searchQuery.value.trim().length >= 2) {
            isSearching.value = true;
            debouncedSearch();
        } else {
            searchResults.value = [];
        }
    });

    return {
        isOpen,
        openSearch,
        closeSearch,
        searchQuery,
        activeCategory,
        isSearching,
        searchResults,
        recentSearches,
        popularSearches,
        addRecentSearch,
        removeRecentSearch,
        clearRecentSearches,
        performSearch
    };
}
