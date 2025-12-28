import { ref, computed, watch, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useUserApi } from "@/api/user/userApi";
import { useAuthStore } from "@/stores/auth";
import { GenderSetting } from "@/types/enums/gender-enum";
import { KyokusuAppRole } from "@/types/enums/role-enum";
import type { User, GetUserDto } from "@/types/backend/user";

export function useProfile() {
    const route = useRoute();
    const router = useRouter();
    const userApi = useUserApi();
    const authStore = useAuthStore();

    const isLoading = ref(false);
    const fetchedUser = ref<GetUserDto | null>(null);

    const profileTabs = [
        { id: 'overview', label: 'Обзор' },
        { id: 'bookmarks', label: 'Закладки' },
        { id: 'comments', label: 'Комментарии' },
    ];

    const isSelfProfile = computed(() => {
        const routeId = Number(route.params.id);
        return authStore.isAuthenticated && authStore.user?.id === routeId;
    });

    const profileData = computed<User | GetUserDto | null>(() => {
        if (isSelfProfile.value) {
            return authStore.user;
        }
        return fetchedUser.value;
    });

    const isPublicAccount = computed(() => !!profileData.value?.is_public);

    const getRoleColor = (role?: string) => {
        switch (role) {
            case KyokusuAppRole.ADMIN: 
                return 'text-red-400 bg-red-400/10 border-red-400/20';
            case KyokusuAppRole.MODERATOR: 
                return 'text-blue-400 bg-blue-400/10 border-blue-400/20';
            default: 
                return 'text-zinc-400 bg-zinc-400/10 border-zinc-400/20';
        }
    };

    const getGenderText = (gender?: string) => {
        switch (gender) {
            case GenderSetting.MALE: return 'Мужской';
            case GenderSetting.FEMALE: return 'Женский';
            case GenderSetting.HIDDEN: 
            default: return 'Не указан';
        }
    };

    const accountCreated = computed(() => {
        if (isSelfProfile.value && profileData.value) {
            return new Date((profileData.value as User).create_at).toLocaleDateString('ru-RU');
        }
        return 'Неизвестно';
    });

    const userRoleColor = computed(() => getRoleColor(profileData.value?.role));
    
    const userGender = computed(() => getGenderText(profileData.value?.gender));

    const loadProfileData = async (id: number) => {
        if (isSelfProfile.value) {
            fetchedUser.value = null; 
            return;
        }

        isLoading.value = true;
        try {
            const data = await userApi.getUser(id);

             if (!data) {
                router.replace({ name: "notFound" });
                return;
            }

            fetchedUser.value = data || null;
        } catch (error) {
            console.error(error);
            fetchedUser.value = null;
        } finally {
            isLoading.value = false;
        }
    };

    const init = () => {
        const id = Number(route.params.id);
        if (!isNaN(id)) {
            loadProfileData(id);
        }
    };

    onMounted(init);

    watch(() => route.params.id, init, { immediate: true });

    return {
        profileData,
        isSelfProfile,
        isPublicAccount,
        isLoading,
        profileTabs,
        
        accountCreated,
        userRoleColor,
        userGender,

        getRoleColor,
        getGenderText,
        loadProfileData
    };
}