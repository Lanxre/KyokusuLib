import { GenderSetting } from "../enums/gender-enum";

export interface GetUserDto {
	id: number;
	name: string;
	picture: string;
	banner: string;
	role: string;
	status: string;
	about: string;
	gender: GenderSetting.MALE | GenderSetting.FEMALE | GenderSetting.HIDDEN;
	birthday: string;
	is_public: boolean;
	last_login: string;
	create_at: string;
	active_tag: string;
	tags: UserTagDTO[];
	settings: PublicUserSettingsDTO;
	user_level: UserLevel;
}

export interface UserTagDTO {
	id: number;
	tag: string;
}

export interface PublicUserSettingsDTO {
	is_show_tag: boolean;
}

export interface UserLevel {
  level: number;
  experience: number;
  xp_needed_for_next: number;
  level_title: string;
}
