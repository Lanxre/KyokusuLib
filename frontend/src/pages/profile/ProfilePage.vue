<script lang="ts" setup>
import FooterApp from "@/components/common/FooterApp.vue";
import HeaderApp from "@/components/common/HeaderApp.vue";
import { useProfile } from "@/composables/api/profile/useProfile";
import ProfilePrivate from "@/layouts/profile/ProfilePrivate.vue";
import ProfilePublic from "@/layouts/profile/ProfilePublic.vue";
import ProfileSelfContent from "@/layouts/profile/ProfileSelfContent.vue";

const { isSelfProfile, isPublicAccount } = useProfile();

</script>

<template>
    <div class="min-h-screen flex flex-col">
        <HeaderApp/>
       
        <main class="grow">
            <ProfileSelfContent v-if="isSelfProfile"/>
            <ProfilePublic v-else-if="!isSelfProfile && isPublicAccount"/>
            <ProfilePrivate v-else-if="!isSelfProfile && !isPublicAccount"/>
        </main>
        <FooterApp/>
    </div>
</template>