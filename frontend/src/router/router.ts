import { createRouter, createWebHistory } from "vue-router";
import LoginPage from "@/pages/auth/LoginPage.vue";
import RegisterPage from "@/pages/auth/RegisterPage.vue";
import HomePage from "@/pages/home/HomePage.vue";
import NotFoundPage from "@/pages/not-found/NotFoundPage.vue";
import SettingsPage from "@/pages/settings/SettingsPage.vue";
import ProfilePage from "@/pages/profile/ProfilePage.vue";
import { useAuthStore } from "@/stores/auth";

const routes = [
	{
		path: "/",
		name: "main",
		component: HomePage,
	},
	{
		path: "/login",
		name: "login",
		component: LoginPage,
		meta: { requiresGuest: true },
	},
	{
		path: "/register",
		name: "register",
		component: RegisterPage,
		meta: { requiresGuest: true },
	},
	{
		path: "/settings",
		name: "settings",
		component: SettingsPage,
		meta: { requiresAuth: true },
	},
	{
		path: "/profile",
		name: "current-profile",
		meta: { requiresAuth: true },
		beforeEnter: (to, from, next) => {
			const authStore = useAuthStore();
			
			if (authStore.user && authStore.user.id) {
				next({ name: "profile", params: { id: authStore.user.id } });
			} else {
				next("/login");
			}
		},
	},
	{
		path: "/profile/:id",
		name: "profile",
		component: ProfilePage,
		meta: { requiresAuth: true },
	},
	{
		path: "/:pathMatch(.*)*",
		name: "notFound",
		component: NotFoundPage,
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

router.beforeEach(async (to, _, next) => {
	const authStore = useAuthStore();
	const hasCookie = document.cookie.includes("KYOKUSU_API_TOKEN");

	if (hasCookie && !authStore.isAuthenticated) {
		try {
			await authStore.initAuth(); 
		} catch (e) {
			console.error("Auth init failed in router", e);
		}
	}


	if (to.meta.requiresAuth && !authStore.isAuthenticated) {
		next("/login");
	} else if (to.meta.requiresGuest && authStore.isAuthenticated) {
		next("/"); 
	} else {
		next();
	}
});

export default router;