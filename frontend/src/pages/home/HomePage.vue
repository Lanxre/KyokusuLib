<script lang="ts" setup>
import FooterApp from "@/components/common/FooterApp.vue";
import HeaderApp from "@/components/common/HeaderApp.vue";
import MainContent from "@/layouts/home/MainContent.vue";
</script>

<template>
    <div class="min-h-screen flex flex-col">
        <HeaderApp/>
       
        <main class="grow">
            <MainContent/>
        </main>
        <FooterApp/>
    </div>
</template>