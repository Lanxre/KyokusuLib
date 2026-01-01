import type { UserActivity } from '@/types/backend/user_activity';
import { UserActivityEnum } from '@/types/enums/activity-enum';
import type { ActivityUIConfig } from '@/types/frontend/user-activity';

type ConfigGenerator = (activity: UserActivity) => ActivityUIConfig;

const DEFAULT_CONFIG: ActivityUIConfig = {
    icon: 'activity',
    color: 'bg-zinc-100 text-zinc-600 dark:bg-zinc-800 dark:text-zinc-400',
    title: 'Активность',
    description: 'Действие пользователя'
};

const strategies: Partial<Record<UserActivityEnum, ConfigGenerator>> = {
    [UserActivityEnum.RANOBE_ADD]: (activity) => ({
        icon: 'book',
        color: 'bg-blue-100 text-blue-600 dark:bg-blue-900/30 dark:text-blue-400',
        title: 'Добавлено ранобэ',
        description: `"${activity.metadata.ranobe_title || 'Без названия'}" от ${activity.metadata.author || 'Неизвестный'}`
    }),

    [UserActivityEnum.ACHEIVEMENT_EARNED]: (activity) => ({
        icon: 'trophy',
        color: 'bg-yellow-100 text-yellow-600 dark:bg-yellow-900/30 dark:text-yellow-400',
        title: 'Получено достижение',
        description: `${activity.metadata.name || 'Достижение'} — ${activity.metadata.desc || ''}`
    }),
};

export function useActivityConfig() {
    const getActivityConfig = (activity: UserActivity): ActivityUIConfig => {
        const strategy = strategies[activity.activity_type as UserActivityEnum];
        
        if (!strategy) {
            return DEFAULT_CONFIG;
        } else {
            return strategy(activity);
        }
        
    };

    return {
        getActivityConfig
    };
}