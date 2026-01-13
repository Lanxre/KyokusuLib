import { defineStore, storeToRefs } from "pinia";
import { computed, watch, ref } from "vue";
import {
	useIdle,
	useWindowFocus,
	useOnline,
	useIntervalFn,
} from "@vueuse/core";
import { useAuthStore } from "@/stores/auth";
import { OnlineStatusEnum } from "@/types/enums/online-status-enum";
import { $api } from "@/composables/api/useApi";

export const useActivityStore = defineStore("activity", () => {
	const authStore = useAuthStore();
	const { isAuthenticated } = storeToRefs(authStore);

	const { idle, lastActive } = useIdle(60 * 1000);
	const focused = useWindowFocus();
	const online = useOnline();

	const isUserActive = computed(() => {
		return online.value && focused.value && !idle.value;
	});

	const sendHeartbeat = async () => {
		if (!isAuthenticated.value || import.meta.server) return;

		try {
			await $api("/api/user/activity", {
				method: "POST",
				body: {
					status: isUserActive.value
						? OnlineStatusEnum.ONLINE
						: OnlineStatusEnum.OFFLINE,
					last_active: lastActive.value,
				},
			});
		} catch (e) {
			console.warn("Heartbeat failed");
		}
	};

	const { pause, resume } = useIntervalFn(() => {
		if (isUserActive.value) {
			sendHeartbeat();
		}
	}, 30000, { immediate: false });

	const initActivityTracking = () => {
		if (import.meta.server) return;

		watch(
			isAuthenticated,
			(isAuth) => {
				if (isAuth) {
					sendHeartbeat();
					resume();
				} else {
					pause();
				}
			},
			{ immediate: true },
		);

		watch(isUserActive, (active) => {
			if (active && isAuthenticated.value) {
				sendHeartbeat();
			}
		});
	};

	return {
		isUserActive,
		initActivityTracking,
	};
});