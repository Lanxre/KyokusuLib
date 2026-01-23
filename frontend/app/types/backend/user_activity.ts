import type { ACTIVITY_TYPES } from "@/constants/user-activity";

export type BookmarkMetadata = {
	novela_title: string;
	desc: string;
};

export type UserNovelaLikeMetadata = {
	name: string;
	desc: string;
};

export type UserNovelaRatingMetadata = {
	name: string;
	desc: string;
	rating: number;
};

export type AchievementMetadata = {
	name: string;
	desc: string;
};

export type RanobeAddMetadata = {
	ranobe_title: string;
	author: string;
};

interface BaseActivity {
	id: number;
	target_id: number;
	timestamp: string;
}

interface BaseCreateActivity {
	user_id: number;
	target_id: number;
}

interface BookmarkActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.NOVELA_BOOKMARK;
	metadata: BookmarkMetadata;
}

interface RemoveBookmarkActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.NOVELA_BOOKMARK_REMOVE;
	metadata: BookmarkMetadata;
}

interface AchievementActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.ACHIEVEMENT_EARNED;
	metadata: AchievementMetadata;
}

interface RanobeActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.RANOBE_ADD;
	metadata: RanobeAddMetadata;
}

interface UserNovelaLikeActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.USER_NOVELA_LIKE;
	metadata: UserNovelaLikeMetadata;
}

interface UserNovelaRemoveLikeActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.USER_NOVELA_LIKE_REMOVE;
	metadata: UserNovelaLikeMetadata;
}

interface UserNovelaRatingActivity extends BaseActivity {
	activity_type: typeof ACTIVITY_TYPES.USER_NOVELA_RATING;
	metadata: UserNovelaRatingMetadata;
}

interface CreateBookmarkActivity extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.NOVELA_BOOKMARK;
	metadata: BookmarkMetadata;
}

interface CreateRemoveBookmarkActivity extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.NOVELA_BOOKMARK_REMOVE;
	metadata: BookmarkMetadata;
}

interface CreateAchievementActivity extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.ACHIEVEMENT_EARNED;
	metadata: AchievementMetadata;
}

interface CreateRanobeActivity extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.RANOBE_ADD;
	metadata: RanobeAddMetadata;
}

interface CreateUserNovelaLike extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.USER_NOVELA_LIKE;
	metadata: UserNovelaLikeMetadata;
}

interface CreateUserNovelaRemoveLike extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.USER_NOVELA_LIKE_REMOVE;
	metadata: UserNovelaLikeMetadata;
}

interface CreateUserNovelaRating extends BaseCreateActivity {
	activity_type: typeof ACTIVITY_TYPES.USER_NOVELA_RATING;
	metadata: UserNovelaRatingMetadata;
}

export type UserActivity =
	| BookmarkActivity
	| AchievementActivity
	| RanobeActivity
	| RemoveBookmarkActivity
	| UserNovelaLikeActivity
	| UserNovelaRemoveLikeActivity
	| UserNovelaRatingActivity;

export type CreateUserActivity =
	| CreateBookmarkActivity
	| CreateAchievementActivity
	| CreateRanobeActivity
	| CreateRemoveBookmarkActivity
	| CreateUserNovelaLike
	| CreateUserNovelaRemoveLike
	| CreateUserNovelaRating;

export interface GetActivityResponse {
	message: UserActivity[];
}
