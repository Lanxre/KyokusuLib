import { $api } from "~/composables/api/useApi";
import type { BackendNotification, NotificationStats } from "@/types/backend/notification";
import { useAuthStore } from "@/stores/auth";

export function useNotifications() {
	const list = useState<BackendNotification[]>("notifications-list", () => []);
	const unreadCount = useState<number>("notifications-unread", () => 0);
	const stream = useState<EventSource | null>("notifications-stream", () => null);

	let reconnectAttempts = 0;
	const MAX_RECONNECT_DELAY = 30000;

	function getLastEventId(): string {
		if (import.meta.server) return "";
		return localStorage.getItem("notifications-last-id") ?? "";
	}

	function setLastEventId(id: number) {
		if (import.meta.server) return;
		localStorage.setItem("notifications-last-id", String(id));
	}

	const authStore = useAuthStore();

	function connect() {
		if (import.meta.server || !authStore.isAuthenticated) return;

		disconnect();

		const url = getLastEventId()
			? `/api/notifications/stream?lastEventId=${getLastEventId()}`
			: "/api/notifications/stream";

		const es = new EventSource(url);

		es.addEventListener("notification", (event: MessageEvent) => {
			try {
				const notification: BackendNotification = JSON.parse(event.data);
				list.value = [notification, ...list.value];
				if (!notification.isRead) {
					unreadCount.value++;
				}
				setLastEventId(notification.id);
			} catch {}
		});

		es.onerror = () => {
			es.close();
			scheduleReconnect();
		};

		stream.value = es;
		reconnectAttempts = 0;
	}

	function scheduleReconnect() {
		const delay = Math.min(
			1000 * 2 ** reconnectAttempts,
			MAX_RECONNECT_DELAY,
		);
		reconnectAttempts++;
		setTimeout(() => connect(), delay);
	}

	function disconnect() {
		if (stream.value) {
			stream.value.close();
			stream.value = null;
		}
	}

	async function fetchList(limit?: number, offset?: number) {
		try {
			const data = await $api<BackendNotification[]>("/api/notifications", {
				query: { limit, offset },
			});
			list.value = data;
			return data;
		} catch {
			return [];
		}
	}

	async function fetchStats() {
		try {
			const data = await $api<NotificationStats>("/api/notifications/stats");
			unreadCount.value = data.unreadCount ?? 0;
			return data;
		} catch {
			return { unreadCount: 0 };
		}
	}

	async function markRead(id: number) {
		try {
			await $api(`/api/notifications/${id}/read`, { method: "PATCH" });
			const item = list.value.find((n) => n.id === id);
			if (item && !item.isRead) {
				item.isRead = true;
				unreadCount.value = Math.max(0, unreadCount.value - 1);
			}
		} catch {}
	}

	async function markAllRead() {
		try {
			await $api("/api/notifications/read-all", { method: "PATCH" });
			list.value.forEach((n) => { n.isRead = true; });
			unreadCount.value = 0;
		} catch {}
	}

	return {
		list,
		unreadCount,
		connect,
		disconnect,
		fetchList,
		fetchStats,
		markRead,
		markAllRead,
	};
}
