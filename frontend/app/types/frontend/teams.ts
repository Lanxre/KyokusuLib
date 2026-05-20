export interface Team {
	id: number;
	owner_id: number;
	name: string;
	slug: string;
	description: string;
	avatar_url: string;
	banner_url?: string;
  created_at: string;
  member_count: number;
	subscribers_count: number;
}
