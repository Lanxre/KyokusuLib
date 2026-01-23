import { ref } from "vue";
import { useNotificationStore } from "~/stores/notification";
import { useApi, $api } from "../useApi";
import type { Team } from "~/types/frontend/teams";

export function useTeam() {
	const { notify } = useNotificationStore();
	const team = ref<Team | null>(null);
	const isLoading = ref(false);
	const error = ref<string | null>(null);

	const fetchTeam = async (slug: string) => {
		isLoading.value = true;
		error.value = null;
		try {
			const { data, error: fetchError } = await useApi<Team>(
				`/api/teams/${slug}`,
			);

			if (fetchError.value) throw new Error(fetchError.value.message);
			team.value = data.value!;
		} catch (e: any) {
			error.value = e.message;
			team.value = null;
		} finally {
			isLoading.value = false;
		}
	};

	const createTeam = async (form: {
		name: string;
		slug: string;
		description: string;
	}) => {
		isLoading.value = true;
		try {
			const data = await $api<Team>("/api/teams", {
				method: "POST",
				body: form,
			});

			notify({ type: "success", title: "Успешно", content: "Команда создана" });
			return data!;
		} catch (e: any) {
			console.log(e);
			notify({ type: "error", title: "Ошибка", content: e.message });
			return null;
		} finally {
			isLoading.value = false;
		}
	};

	const updateTeam = async (slug: string, form: FormData) => {
		isLoading.value = true;
		try {
			const { data, error: updateError } = await useApi<Team>(
				`/api/teams/${slug}`,
				{
					method: "PATCH",
					body: form,
				},
			);

			if (updateError.value)
				throw new Error(
					updateError.value.data?.error || "Failed to update team",
				);

			if (team.value && data.value) {
				Object.assign(team.value, data.value);
			}

			notify({
				type: "success",
				title: "Успешно",
				content: "Команда обновлена",
			});
			return data.value;
		} catch (e: any) {
			notify({ type: "error", title: "Ошибка", content: e.message });
			return null;
		} finally {
			isLoading.value = false;
		}
	};

	const joinTeam = async (slug: string) => {
		try {
			const { error: joinError } = await useApi(`/api/teams/${slug}/join`, {
				method: "POST",
			});

			if (joinError.value) throw new Error(joinError.value.data?.error);

			notify({
				type: "success",
				title: "Успешно",
				content: "Вы вступили в команду",
			});
			return true;
		} catch (e: any) {
			notify({ type: "error", title: "Ошибка", content: e.message });
			return false;
		}
	};

	return {
		team,
		isLoading,
		error,
		fetchTeam,
		createTeam,
		updateTeam,
		joinTeam,
	};
}
