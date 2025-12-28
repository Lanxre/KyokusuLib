<script setup lang="ts">
import DiscordIcon from "@/assets/images/social/discord.png";
import GoogleIcon from "@/assets/images/social/google.png";
import LogoutIcon from "@/assets/images/special/log-out.png";
import { useSocialNetworkSettings } from "@/composables/api/settings/useSocialNetworkSettings";

const { 
    socialNetworkData,
    handleDiscord,
    handleGoogle,
} = useSocialNetworkSettings();



</script>

<template>
    <div class="space-y-8 cursor-default transition-colors duration-300">
        <div>
            <h2 class="text-xl font-medium mb-1 text-zinc-900 dark:text-white">Социальные сети</h2>
        </div>
        <div class="space-y-6 text-left">
            <div class="pb-6 border-b space-y-4 border-zinc-200 dark:border-zinc-800">
                <h3 class="text-base leading-xl font-semibold text-zinc-800 dark:text-zinc-200">Подключение</h3>  
            </div>
            
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6">
                <div class="group flex flex-col items-start justify-between p-4 rounded-lg border transition-all duration-300 bg-white border-zinc-200 hover:border-zinc-300 hover:shadow-md dark:bg-[#5865f2]/20 dark:hover:bg-[#5865f2]/30 dark:border-zinc-800 dark:hover:border-[#5865f2]/50 cursor-pointer">
                    <div class="space-y-1 flex flex-col w-full" v-on:click="handleDiscord">
                        <div class="flex items-center justify-between w-full">
                            <span class="text-lg font-medium block text-zinc-700 dark:text-zinc-200">Discord</span>
                            <img :src="DiscordIcon" class="w-8 h-8 transition-transform group-hover:scale-110 dark:invert"/>
                        </div>

                        <div v-if="socialNetworkData.is_discord_connected" class="flex items-center gap-4">
                            <img :src="LogoutIcon" class="w-6 h-6 transition-transform group-hover:scale-110 dark:invert"/>
                            <p class="text-xs text-zinc-500 dark:text-stone-300">
                                Уже подключен
                            </p>
                        </div>
                        <div v-else>
                            <p class="text-xs text-zinc-500 dark:text-stone-300">
                                Подключить аккаунт Discord
                            </p>
                        </div>
                    </div>
                </div>

                <div class="group flex flex-col items-start justify-between p-4 rounded-lg border transition-all duration-300 bg-white border-zinc-200 hover:border-zinc-300 hover:shadow-md dark:bg-zinc-900 dark:hover:bg-zinc-800 dark:border-zinc-800 cursor-pointer overflow-hidden relative">
                    <div class="absolute inset-0 opacity-0 dark:opacity-10 bg-linear-to-r from-red-500 via-green-500 to-blue-500 transition-opacity group-hover:opacity-20 pointer-events-none"></div>
                    
                    <div class="space-y-1 w-full relative z-10" v-on:click="handleGoogle">
                        <div class="flex items-center justify-between w-full">
                            <span class="text-lg font-medium block text-zinc-700 dark:text-zinc-200">Google</span>
                            <img :src="GoogleIcon" class="w-8 h-8 transition-transform group-hover:scale-110 invert-0 dark:invert"/>
                        </div>
                        <div v-if="socialNetworkData.is_google_connected" class="flex items-center gap-4">
                            <img :src="LogoutIcon" class="w-6 h-6 transition-transform group-hover:scale-110 dark:invert"/>
                            <p class="text-xs text-zinc-500 dark:text-stone-300">
                                Уже подключен
                            </p>
                        </div>
                        <div v-else>
                            <p class="text-xs text-zinc-500 dark:text-stone-300">
                                Подключить аккаунт Google
                            </p>
                        </div>
                    </div>
                </div>
               
                
            </div>
        </div>
    </div>
</template>