import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import RedirectView from "../views/RedirectView.vue";
import DashboardView from "../views/DashboardView.vue";

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
    },
  ],
  scrollBehavior(_to, _from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    return { top: 0 };
  },
});

export default router;
