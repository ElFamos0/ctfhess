import { createWebHashHistory, createRouter } from "vue-router";

import HomeView from "../views/HomeView.vue";
import ProfilView from "../views/ProfilView.vue";
import ChallView from "../views/ChallView.vue";
import EditChall from "../views/EditChall.vue";

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
    {
        path: "/edit",
        name: "Edit",
        component: EditChall,
        meta: {
            title: 'Edit',
        },
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});


export default router;