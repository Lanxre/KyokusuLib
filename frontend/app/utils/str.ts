import PersonIcon from "@/assets/images/special/user.png";
import { getApiBase } from "@/constants/urls";

export function correctProfileImageLink(link: string) {
	const backendUrl = getApiBase();
	return `${backendUrl}/${link}`;
}

export function correctProfileImage(picture: string | null | undefined) {
	if (picture && picture.startsWith("/uploads")) {
		const backendUrl = getApiBase();
		return `${backendUrl}${picture}`;
	} else if (picture) {
		return picture;
	} else {
		return PersonIcon;
	}
}

export const formatDateUserActivity = (dateString: string) => {
	const date = new Date(dateString);
	const now = new Date();
	const diff = now.getTime() - date.getTime();

	const minutes = Math.floor(diff / 60000);
	const hours = Math.floor(diff / 3600000);
	const days = Math.floor(diff / 86400000);

	if (minutes < 1) return "Только что";
	if (minutes < 60) return `${minutes} м. назад`;
	if (hours < 24) return `${hours} ч. назад`;
	if (days < 7) return `${days} д. назад`;

	return date.toLocaleDateString("ru-RU", {
		day: "numeric",
		month: "short",
		year: "numeric",
	});
};
