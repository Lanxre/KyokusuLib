export function parseDateToISO(input: string): string {
	const date = new Date(input).toISOString().split("T")[0]!;
	return date;
}

export function parseDateToLocale(input: string): string {
	return new Date(input).toLocaleDateString("ru-RU", {
		day: "2-digit",
		month: "2-digit",
		year: "numeric",
	});
}
