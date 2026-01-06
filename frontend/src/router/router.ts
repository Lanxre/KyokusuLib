import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { KyokusuAppRole, ROLE_WEIGHTS } from "@/types/enums/role-enum";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "main",
    component: () => import("@/pages/home/HomePage.vue"),
  },
  {
    path: "/login",
    name: "login",
    component: () => import("@/pages/auth/LoginPage.vue"),
    meta: { requiresGuest: true },
  },
  {
    path: "/register",
    name: "register",
    component: () => import("@/pages/auth/RegisterPage.vue"),
    meta: { requiresGuest: true },
  },
  {
    path: "/settings",
    name: "settings",
    component: () => import("@/pages/settings/SettingsPage.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/profile",
    name: "current-profile",
    meta: { requiresAuth: true },
    redirect: () => {
      const authStore = useAuthStore();
      if (authStore.user?.id) {
        return { name: "profile", params: { id: authStore.user.id } };
      }
      return { name: "login" };
    },
  },
  {
    path: "/profile/:id",
    name: "profile",
    component: () => import("@/pages/profile/ProfilePage.vue"),
    meta: { requiresAuth: true },
    beforeEnter: (to, _, next) => {
          const id = to.params.id;
          if (typeof id === 'string' && /^\d+$/.test(id)) {
            next();
          } else {
            next("/");
          }
    },
  },
  {
    path: "/author/add",
    name: "add-author",
    component: () => import("@/pages/forms/AddAuthorPage.vue"),
    meta: { requiresAuth: true, minRole: KyokusuAppRole.MODERATOR },
  },
  {
    path: "/novela/add",
    name: "add-novela",
    component: () => import("@/pages/forms/AddNovelaPage.vue"),
    meta: { requiresAuth: true, minRole: KyokusuAppRole.MODERATOR },
  },
  {
    path: "/:pathMatch(.*)*",
    name: "notFound",
    component: () => import("@/pages/not-found/NotFoundPage.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _, next) => {
  const authStore = useAuthStore();
  const hasTokenCookie = document.cookie.includes("KYOKUSU_API_TOKEN");

  if (hasTokenCookie && !authStore.isAuthenticated) {
    try {
      await authStore.initAuth();
    } catch {
      return next("/login");
    }
  }

  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    return next("/");
  }

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return next("/login");
  }

  if (to.meta.minRole) {
    const userRole = (authStore.user?.role as KyokusuAppRole) ?? KyokusuAppRole.USER;
    const userWeight = ROLE_WEIGHTS[userRole] ?? 0;
    const requiredWeight = ROLE_WEIGHTS[to.meta.minRole as KyokusuAppRole] ?? 0;

    if (userWeight < requiredWeight) {
      return next("/");
    }
  }

  next();
});

export default router;