<script setup lang="ts">
import { useProfile } from "@/composables/api/profile/useProfile";
import ProfileSelfContent from "@/components/app/profile/ProfileSelfContent.vue";
import ProfilePublic from "@/components/app/profile/ProfilePublic.vue";
import ProfilePrivate from "@/components/app/profile/ProfilePrivate.vue";

definePageMeta({
	validate: async (route) => {
		return /^\d+$/.test(route.params.id as string);
	},
});

const { profileData, isSelfProfile, isPublicAccount, isLoading } = useProfile();

useSeoMeta({
	title: () =>
		profileData.value?.name ? `${profileData.value.name} - Профиль` : "Профиль",
	ogTitle: () => profileData.value?.name,
	ogImage: () => profileData.value?.picture,
});
</script>

<template>
    <div v-if="isLoading" class="flex justify-center py-20">
        <svg class="animate-spin h-10 w-10 text-zinc-500"></svg>
    </div>
    <div v-else="profileData" class="bg-white dark:bg-[#0f0f0f]">
        <ProfileSelfContent v-if="isSelfProfile"/>
        <ProfilePublic v-else-if="isPublicAccount" :profileData="profileData"/>
        <ProfilePrivate v-else :profileData="profileData"/>
    </div>
</template>