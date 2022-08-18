import { createWebHashHistory, createRouter } from "vue-router";

import HomeView from "../views/HomeView.vue";
import LoginView from "../views/LoginView.vue";
import ProfilView from "../views/ProfilView.vue";
import ChallView from "../views/ChallView.vue";

const routes = [
    {
        path: "/",
        name: "home",
        component: HomeView,
        meta: {
            title: 'Maison',
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
    {
        path: "/chall",
        name: "Chall",
        component: ChallView,
        meta: {
            title: 'Chall',
        },
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});


export default router;