import type { UserActivity } from '@/types/backend/user_activity';

export interface ActivityUIConfig {
    icon: string;
    color: string;
    title: string;
    description: string;
}

export function useActivityConfig() {
    const getActivityConfig = (activity: UserActivity): ActivityUIConfig => {
        switch (activity.activity_type) {
            case 'ranobe_added':
                return {
                    icon: 'book',
                    color: 'bg-blue-100 text-blue-600 dark:bg-blue-900/30 dark:text-blue-400',
                    title: 'Добавлено ранобэ',
                    description: `"${activity.metadata.ranobe_title}" от ${activity.metadata.author}`
                };
            case 'achievement_earned':
                return {
                    icon: 'trophy',
                    color: 'bg-yellow-100 text-yellow-600 dark:bg-yellow-900/30 dark:text-yellow-400',
                    title: 'Получено достижение',
                    description: `${activity.metadata.name} — ${activity.metadata.desc}`
                };
            default:
                return {
                    icon: 'activity',
                    color: 'bg-zinc-100 text-zinc-600 dark:bg-zinc-800 dark:text-zinc-400',
                    title: 'Активность',
                    description: 'Действие пользователя'
                };
        }
    };

    return {
        getActivityConfig
    };
}