import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import RedirectView from "../views/RedirectView.vue";
import DashboardView from "../views/DashboardView.vue";
import GuildView from "../views/GuildView.vue";

import { useUserStore } from "@/stores/user";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/callback",
      name: "redirect",
      component: RedirectView,
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: DashboardView,
      meta: { requiresAuth: true },
    },
    {
      path: "/guilds/:guildID",
      name: "guild",
      component: GuildView,
      meta: { requiresAuth: true },
    },
  ],
  scrollBehavior(_to, _from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    return { top: 0 };
  },
});

router.beforeEach((to, from) => {
  const userStore = useUserStore();
  if (to.meta.requiresAuth && !userStore.loggedIn) {
    return "/";
  }
});

export default router;
