import { GenderSetting } from "../enums/gender-enum";

export interface User {
	id: number;
	email: string;
	name: string;
	picture: string;
	banner: string;
	role: string;
	status: string;
	last_login: string;
	create_at: string;
	about: string;
	gender: GenderSetting.MALE | GenderSetting.FEMALE | GenderSetting.HIDDEN;
	birthday: string;
	is_public: boolean;
	tag: string;
}

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
	tag: string;
}