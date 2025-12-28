<script lang="ts" setup>
import FooterApp from "@/components/common/FooterApp.vue";
import HeaderApp from "@/components/common/HeaderApp.vue";
import SettingsContent from "@/layouts/settings/SettingsContent.vue";
</script>

<template>
    <div class="min-h-screen flex flex-col">
        <HeaderApp/>
       
        <main class="flex-grow">
            <SettingsContent/>
        </main>
        <FooterApp/>
    </div>
</template>