import { createWebHashHistory, createRouter } from "vue-router";


import HomeView from "../views/HomeView.vue";
import LoginView from "../views/LoginView.vue";
import ProfilView from "../views/ProfilView.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: HomeView,
        meta: {
          title: 'Big Maison',
        },
      },
    {
        path: "/login",
        name: "Login",
        component: LoginView,
        meta: {
            title: 'Login',
        },
        },
    {
        path: "/profil",
        name: "Profil",
        component: ProfilView,
        meta: {
            title: 'Profil',
        },
        },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
  });


export default router;