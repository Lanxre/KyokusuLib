export interface NovelaChapterImage {
	id: number;
	image_url: string;
	caption: string;
}

export interface NovelaChapter {
	id: number;
	title: string;
	number: number;
	content: string;
	images: NovelaChapterImage[];
}

export interface NovelaVolume {
	id: number;
	title: string;
	number: number;
	chapters: NovelaChapter[];
}

export interface NovelaAuthor {
	id: number;
	name: string;
	role: string;
}

export interface NovelaDetails {
	id: number;
	title: string;
	alternative_titles: string[];
	description: string;
	type: string;
	age_rating: string;
	release_date: string;
	status: string;
	poster_url?: string;
	translation_status: string;
	rating: number;
	views: number;
	authors?: NovelaAuthor[];
	genres: string[];
	categories: string[];
	volumes: NovelaVolume[];
	bookmark?: string;
}
