<script setup lang="ts">
import { useTeam } from '~/composables/api/teams/useTeams';
import type { TeamMember } from '~/types/frontend/teams';
import TeamMemberCard from './TeamMemberCard.vue';

const props = defineProps<{
    slug: string;
}>();

const { getTeamMembers } = useTeam();

const { data: members, pending: isLoading } = useAsyncData<TeamMember[]>(
    `team-members-${props.slug}-preview`,
    async () => {
        try {
            return await getTeamMembers(props.slug, 6, 0);
        } catch (e) {
            console.error('Failed to fetch team members preview', e);
            return [];
        }
    },
    { default: () => [] }
);
</script>

<template>
    <div v-if="isLoading || members.length > 0" class="bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-800 rounded-3xl p-6 sm:p-8 shadow-sm w-full mt-6">
        <div class="flex items-end justify-between mb-6">
            <h2 class="flex items-center gap-3 text-xl md:text-2xl font-black tracking-tight text-zinc-900 dark:text-white uppercase italic">
                <div class="flex flex-col items-start gap-1">
                    Участники команды
                    <div class="h-1 w-1/2 bg-yellow-500 rounded-full"></div>
                </div>
            </h2>
        </div>

        <div class="mt-2">
            <div v-if="isLoading" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
                <div v-for="i in 6" :key="i" class="h-20 rounded-2xl bg-zinc-100 dark:bg-zinc-800/50 animate-pulse"></div>
            </div>

            <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
                <TeamMemberCard 
                    v-for="member in members" 
                    :key="member.user.id" 
                    :member="member" 
                />
            </div>
            
            <div class="mt-6 flex justify-center sm:justify-end">
                <NuxtLink 
                    :to="`/team/${slug}/members`"
                    class="px-6 py-2 bg-zinc-100 dark:bg-zinc-800 hover:bg-zinc-200 dark:hover:bg-zinc-700 text-zinc-700 dark:text-zinc-300 font-bold rounded-full transition-colors duration-300 flex items-center gap-2"
                >
                    Все участники
                    <Icon name="ph:arrow-right-bold" />
                </NuxtLink>
            </div>
        </div>
    </div>
</template>