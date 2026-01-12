import type { ACTIVITY_TYPES } from "@/constants/user-activity";

export type BookmarkMetadata = {
  novela_title: string;
  desc: string;
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

export type UserActivity = BookmarkActivity | AchievementActivity | RanobeActivity | RemoveBookmarkActivity;
export type CreateUserActivity = 
	| CreateBookmarkActivity 
	| CreateAchievementActivity 
	| CreateRanobeActivity
  | CreateRemoveBookmarkActivity;

export interface GetActivityResponse {
	message: UserActivity[];
}
