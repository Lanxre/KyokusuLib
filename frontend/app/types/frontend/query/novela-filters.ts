export interface NovelaFilterState {
	search: string;
	genres: string[];
	categories: string[];
	status: string;
	translationStatus: string;
	type: string;
	chaptersFrom: number | null;
	chaptersTo: number | null;
}

export const DEFAULT_FILTER_STATE: NovelaFilterState = {
	search: '',
	genres: [],
	categories: [],
	status: '',
	translationStatus: '',
	type: '',
	chaptersFrom: null,
	chaptersTo: null,
};

export interface NovelaFilterOption {
	value: string;
	label: string;
}

export const STATUS_OPTIONS: NovelaFilterOption[] = [
	{ value: 'ongoing', label: 'Продолжается' },
	{ value: 'completed', label: 'Завершён' },
	{ value: 'hiatus', label: 'Заморожен' },
	{ value: 'cancelled', label: 'Отменён' },
];

export const TRANSLATION_STATUS_OPTIONS: NovelaFilterOption[] = [
	{ value: 'ongoing', label: 'Переводится' },
	{ value: 'completed', label: 'Переведён' },
	{ value: 'stopped', label: 'Остановлен' },
	{ value: 'licensed', label: 'Лицензирован' },
];

export const NOVELA_TYPE_OPTIONS: NovelaFilterOption[] = [
	{ value: 'ranobe', label: 'Ранобэ' },
	{ value: 'manga', label: 'Манга' },
	{ value: 'manhwa', label: 'Манхва' },
	{ value: 'manhua', label: 'Маньхуа' },
];
