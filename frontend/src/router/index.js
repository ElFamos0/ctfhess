import { createWebHashHistory, createRouter } from "vue-router";

import HomeView from "../views/HomeView.vue";
import ProfilView from "../views/ProfilView.vue";
import ChallView from "../views/ChallView.vue";
import EditChall from "../views/EditChall.vue";
import Scoreboard from "../views/ScoreBoard.vue";
import UploadView from "../views/UploadView.vue";

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
    {
        path: "/scoreboard",
        name: "Scoreboard",
        component: Scoreboard,
        meta: {
            title: 'Scoreboard',
        },
    },
    {
        path: "/upload",
        name: "Upload",
        component: UploadView,
        meta: {
            title: 'Upload',
        },
    },

];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});


export default router;