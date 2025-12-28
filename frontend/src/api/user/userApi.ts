import { GetUserDto } from "@/types/backend/user";
import { useApi } from "../api";

export function useUserApi() {
    const getUser = async (userId: number) => {
        if (userId === null || userId === undefined) return null;

        try {
            const { data } = await useApi(`/api/user/${userId}`, {
                credentials: "include",
            }).get().json<GetUserDto>();

            return data.value;
        } catch (e) {
            console.error(`USER FETCHER: [${e}]`);
            return null;
        }
    };

    return {
        getUser,
    };
}