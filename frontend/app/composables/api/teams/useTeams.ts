import { ref } from "vue";
import { useNotificationStore } from "~/stores/notification";
import { useApi, $api } from "../useApi";
import type { Team, TeamMember, TeamSubscriber } from "~/types/frontend/teams";

export function useTeam() {
	const { notify } = useNotificationStore();
	const team = ref<Team | null>(null);
	const isLoading = ref(false);
	const error = ref<string | null>(null);

	const getTeamBySlug = async (slug: string) => {
		isLoading.value = true;
		error.value = null;
		try {
			const data = await $api<Team>(`/api/teams/${slug}`);
			return data;
		} catch (e: any) {
			error.value = e.message;
			return null;
		} finally {
			isLoading.value = false;
		}
	};

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
		isLoading.value = true;
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
		} finally {
			isLoading.value = false;
		}
	};

	const leaveTeam = async (slug: string) => {
		isLoading.value = true;
		try {
			const { error: leaveError } = await useApi(`/api/teams/${slug}/leave`, {
				method: "POST",
			});

			if (leaveError.value) throw new Error(leaveError.value.data?.error);

			notify({
				type: "success",
				title: "Успешно",
				content: "Вы покинули команду",
			});
			return true;
		} catch (e: any) {
			notify({ type: "error", title: "Ошибка", content: e.message });
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const getUserTeams = async (userId: number) => {
		isLoading.value = true;
		error.value = null;
		try {
			const data = await $api<Team[]>(`/api/teams`, {
				query: { user_id: userId }
			});
			return data || [];
		} catch (e: any) {
			error.value = e.message;
			return [];
		} finally {
			isLoading.value = false;
		}
	};

	const getTeamMembers = async (slug: string, limit: number = 20, offset: number = 0) => {
		isLoading.value = true;
		error.value = null;
		try {
			const data = await $api<TeamMember[]>(`/api/teams/${slug}/members`, {
				query: { limit, offset }
			});
			return data || [];
		} catch (e: any) {
			error.value = e.message;
			return [];
		} finally {
			isLoading.value = false;
		}
  };

	const getTeamSubscribers = async (slug: string, limit: number = 20, offset: number = 0) => {
		isLoading.value = true;
		error.value = null;
		try {
			const data = await $api<TeamSubscriber[]>(`/api/teams/${slug}/subscribers`, {
				query: { limit, offset }
			});
			return data || [];
		} catch (e: any) {
			error.value = e.message;
			return [];
		} finally {
			isLoading.value = false;
		}
  };

	const subscribeTeam = async (teamId: number) => {
		isLoading.value = true;
		try {
			const { error: subscribeError } = await useApi(`/api/teams/subscribe`, {
				method: "POST",
				query: { team_id: teamId }
			});

			if (subscribeError.value) throw new Error(subscribeError.value.data?.error);

			notify({
				type: "success",
				title: "Успешно",
				content: "Вы подписались на команду",
			});
			return true;
		} catch (e: any) {
			notify({ type: "error", title: "Ошибка", content: e.message });
			return false;
		} finally {
			isLoading.value = false;
		}
  };

	const unsubscribeTeam = async (teamId: number) => {
		isLoading.value = true;
		try {
			const { error: unsubscribeError } = await useApi(`/api/teams/unsubscribe`, {
				method: "POST",
				query: { team_id: teamId }
			});

			if (unsubscribeError.value) throw new Error(unsubscribeError.value.data?.error);

			notify({
				type: "success",
				title: "Успешно",
				content: "Вы отписались от команды",
			});
			return true;
		} catch (e: any) {
			notify({ type: "error", title: "Ошибка", content: e.message });
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	return {
		team,
		isLoading,
		error,
		getTeamBySlug,
		fetchTeam,
		createTeam,
		updateTeam,
		joinTeam,
		leaveTeam,
		getUserTeams,
    getTeamMembers,
		getTeamSubscribers,
		subscribeTeam,
		unsubscribeTeam,
	};
}
