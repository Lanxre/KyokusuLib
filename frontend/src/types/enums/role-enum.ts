export enum KyokusuAppRole {
	ADMIN = "admin",
	MODERATOR = "moderator",
	USER = "user",
}

export const ROLE_WEIGHTS: Record<KyokusuAppRole, number> = {
    [KyokusuAppRole.USER]: 10,
    [KyokusuAppRole.MODERATOR]: 50,
    [KyokusuAppRole.ADMIN]: 100,
};

