<script lang="ts" setup>
import FooterApp from "@/components/common/FooterApp.vue";
import HeaderApp from "@/components/common/HeaderApp.vue";
import NotFound from "@/layouts/not-found/NotFound.vue";
</script>

<template>
    <div class="min-h-screen flex flex-col">
        <HeaderApp/>
       
        <main class="flex-grow">
            <NotFound/>
        </main>
        <FooterApp/>
    </div>
</template>