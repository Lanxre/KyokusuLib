<script setup lang="ts">
import { staticImage } from "@/utils/str";
import type { TeamMember } from "~/types/frontend/teams";

defineProps<{
    member: TeamMember;
}>();
</script>

<template>
    <div class="flex items-center gap-4 bg-zinc-50 dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-800 rounded-2xl p-4 transition-colors hover:bg-zinc-100 dark:hover:bg-zinc-800/40 cursor-pointer">
        <NuxtLink :to="`/profile/${member.user.id}`" class="flex items-center gap-4 w-full">
            <img 
                v-if="member.user.picture"
                :src="staticImage(member.user.picture)" 
                alt="Avatar" 
                class="w-12 h-12 rounded-full object-cover"
            />
            <Icon v-else name="ph:user-bold" size="20" class="text-zinc-700/40" />
            
            <div class="flex-1 min-w-0 flex flex-col justify-center">
                <span class="font-bold text-sm text-zinc-900 dark:text-zinc-100 truncate transition-colors">
                    {{ member.user.name }}
                </span>
                <div class="flex items-center gap-2 mt-1">
                    <span class="text-xs font-semibold text-yellow-600 dark:text-yellow-500">
                        {{ member.role_name }}
                    </span>
                    <span v-if="member.user.user_level" class="text-[10px] font-bold px-2 py-0.5 rounded-full bg-zinc-200 dark:bg-zinc-800 text-zinc-600 dark:text-zinc-400">
                        Ур. {{ member.user.user_level.level }}
                    </span>
                </div>
            </div>
        </NuxtLink>
    </div>
</template>