import PersonIcon from "@/assets/images/special/user.png";
import { BACKEND_URL } from "@/const";

export function correctProfileImageLink(link: string) {
	return `${BACKEND_URL}/${link}`;
}

export function correctProfileImage(picture: string) {
  if (
		picture !== undefined &&
		picture !== null &&
		picture.startsWith("/uploads")
	) {
		return `${BACKEND_URL}/${picture}`;
	} else if (picture !== undefined && picture !== null && picture !== "") {
		return picture;
	} else {
		return PersonIcon;
	}
}
