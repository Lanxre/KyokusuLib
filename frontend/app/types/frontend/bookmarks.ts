export type BookmarkCategory = 'planned' | 'reading' | 'completed' | 'on_hold' | 'dropped';

export interface Bookmark {
  user_id: number;
  novela_id: number;
  category: BookmarkCategory;
  created_at: string;
  updated_at: string;
}