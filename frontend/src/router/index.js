import { createWebHashHistory, createRouter } from "vue-router";

import LoginView from "../views/LoginView.vue";
import ProfilView from "../views/ProfilView.vue";

const routes = [
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