export interface TeamRoleNames {
	owner: string;
	moderator: string;
	member: string;
}

export interface Team {
	id: number;
	owner_id: number;
	name: string;
	slug: string;
	description: string;
	avatar_url: string;
	banner_url?: string;
	role_names: TeamRoleNames;
	created_at: string;
	member_count: number;
	subscribers_count: number;
	is_member: boolean;
	stats: {
		member_count: number;
		subscribers_count: number;
	};
}

export interface TeamMember {
	user: {
		id: number;
		name: string;
		picture: string;
		active_tag: string;
		user_level: {
			level: number;
			level_title: string;
		}
	};
	role: string;
	role_name: string;
	joined_at: string;
}
