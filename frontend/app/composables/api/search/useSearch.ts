import { ref, watch, computed } from 'vue';
import { useStorage } from '@vueuse/core';
import { useDebounceFn } from '@vueuse/core';
import { $api } from '@/composables/api/useApi';
import { useAuthors } from '@/composables/api/authors/useAuthors';
import type { NovelaDetails, NovelaAuthorDetails } from '@/types/backend/novela';
import type { GetUserDto } from '@/types/backend/user';
import type { Team } from '@/types/frontend/teams';
import type { MostSearched } from '@/types/frontend/search/searches';


export type SearchCategory = 'ranobe' | 'users' | 'teams' | 'authors';
export type SearchResultItem = NovelaDetails | GetUserDto | Team | NovelaAuthorDetails | any;

export interface UseSearchOptions {
    immediateWatch?: boolean;
}

export function useSearch(options: UseSearchOptions = { immediateWatch: false }) {
    const { user } = useAuthStore();
    const isOpen = useState('search-is-open', () => false);
    const searchQuery = useState('search-query', () => '');

    const genres = useState<MostSearched[]>('search-genres', () => []);
    const categories = useState<MostSearched[]>('search-categories', () => []);
  
    const activeCategory = useState<SearchCategory>('search-category', () => 'ranobe');
    const isSearching = useState('search-is-searching', () => false);
    const searchResults = useState<SearchResultItem[]>('search-results', () => []);
    
    const recentSearches = useStorage<string[]>('recent-searches', []);
    const popularSearches = ref<MostSearched[]>([{
      id: 1,
      type: 'genres',
      label: 'Детектив',
    }, {
      id: 2,
      type: 'genres',
      label: 'Романтика',
    }, {
      id: 3,
      type: 'categories',
      label: 'Перерождение',
    }, {
      id: 4,
      type: 'categories',
      label: 'Академия',
    }]);
    

    const hasActiveFilters = computed(() => genres.value.length > 0 || categories.value.length > 0);
    const isValidQuery = computed(() => searchQuery.value.trim().length >= 2);
    const canSearch = computed(() => isValidQuery.value || hasActiveFilters.value);

    const setQueryParam = (item: MostSearched) => {
        const target = item.type === 'genres' ? genres : categories;
        if (!target.value.some(v => v.id === item.id)) {
            target.value.push(item);
        }
    };

    const removeQueryParam = (item: MostSearched) => {
        const target = item.type === 'genres' ? genres : categories;
        target.value = target.value.filter(v => v.id !== item.id);
    };

    const clearFilters = () => {
        genres.value = [];
        categories.value = [];
    };

    const openSearch = () => {
        isOpen.value = true;
    };

    const closeSearch = () => {
        isOpen.value = false;
        searchQuery.value = '';
        clearFilters();
        searchResults.value = [];
    };

    const addRecentSearch = (query: string) => {
        if (!query.trim()) return;
        const current = [...recentSearches.value];
        const index = current.indexOf(query);
        if (index > -1) current.splice(index, 1);
        current.unshift(query);
        recentSearches.value = current.slice(0, 5);
    };

    const removeRecentSearch = (query: string) => {
        recentSearches.value = recentSearches.value.filter(q => q !== query);
    };

    const clearRecentSearches = () => {
        recentSearches.value = [];
    };

    const performSearch = async () => {
        if (!canSearch.value) {
            searchResults.value = [];
            return;
        }

        const query = searchQuery.value.trim();
        isSearching.value = true;
        
        try {
            if (activeCategory.value === 'ranobe') {
                const apiQuery: Record<string, any> = { 
                    search: query, 
                    limit: 10 
                };

                if (genres.value.length > 0) {
                    apiQuery.genres = genres.value.map(v => v.label).join(',');
                }

                if (categories.value.length > 0) {
                    apiQuery.categories = categories.value.map(v => v.label).join(',');
                }

                const data = await $api<NovelaDetails[]>("/api/novela", {
                    query: apiQuery
                });
                searchResults.value = data || [];
            } else {

                if (searchQuery.value.trim().length < 2) {
                  searchResults.value = [];
                  return;
                };
                
                const endpoint = {
                    authors: "/api/author",
                    users: "/api/user",
                    teams: "/api/teams"
                }[activeCategory.value];

                const data = await $api<any[]>(endpoint, {
                    query: { 
                        search: query, 
                        limit: 10,
                        ...(activeCategory.value === 'teams' && user?.id ? { user_id: user.id } : {})
                    }
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

    const debouncedSearch = useDebounceFn(performSearch, 500);

    if (options.immediateWatch) {
        watch([searchQuery, activeCategory, genres, categories], () => {
            if (canSearch.value) {
                debouncedSearch();
            } else {
                searchResults.value = [];
            }
        }, { deep: true });
    }

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
        genres,
        categories,
        hasActiveFilters,
        addRecentSearch,
        removeRecentSearch,
        clearRecentSearches,
        setQueryParam,
        removeQueryParam,
        clearFilters,
        performSearch
    };
}
