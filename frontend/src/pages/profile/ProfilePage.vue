<script lang="ts" setup>
import FooterApp from "@/components/common/FooterApp.vue";
import HeaderApp from "@/components/common/HeaderApp.vue";
import { useProfile } from "@/composables/api/profile/useProfile";
import ProfilePrivate from "@/layouts/profile/ProfilePrivate.vue";
import ProfilePublic from "@/layouts/profile/ProfilePublic.vue";
import ProfileSelfContent from "@/layouts/profile/ProfileSelfContent.vue";
import { GetUserDto } from "@/types/backend/user";
import { onMounted } from "vue";

const { profileData, isSelfProfile, isPublicAccount, init } = useProfile();

onMounted(init);

</script>

<template>
    <div class="min-h-screen flex flex-col">
        <HeaderApp/>
       
        <main class="grow">
            <ProfileSelfContent v-if="isSelfProfile"/>
            <ProfilePublic v-else-if="!isSelfProfile && isPublicAccount" :profileData="profileData as GetUserDto"/>
            <ProfilePrivate v-else-if="!isSelfProfile && !isPublicAccount"  :profileData="profileData as GetUserDto"/>
        </main>
        <FooterApp/>
    </div>
</template>