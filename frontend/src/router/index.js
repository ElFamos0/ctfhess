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

router.beforeEach((to, from, next) => {
    const nearestWithTitle = to.matched.slice().reverse().find(r => r.meta && r.meta.title);
    const nearestWithMeta = to.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);
    const previousNearestWithMeta = from.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);
  
    // If a route with a title was found, set the document (page) title to that value.
    if(nearestWithTitle) {
      document.title = nearestWithTitle.meta.title;
    } else if(previousNearestWithMeta) {
      document.title = previousNearestWithMeta.meta.title;
    }
  
    // Remove any stale meta tags from the document using the key attribute we set below.
    Array.from(document.querySelectorAll('[data-vue-router-controlled]')).map(el => el.parentNode.removeChild(el));
  
    // Skip rendering meta tags if there are none.
    if(nearestWithMeta) {
      // Turn the meta tag definitions into actual elements in the head.
      nearestWithMeta.meta.metaTags.map(tagDef => {
        const tag = document.createElement('meta');
  
        Object.keys(tagDef).forEach(key => {
          tag.setAttribute(key, tagDef[key]);
        });
  
        // We use this to track which meta tags we create so we don't interfere with other ones.
        tag.setAttribute('data-vue-router-controlled', '');
  
        return tag;
      })
      // Add the meta tags to the document head.
      .forEach(tag => document.head.appendChild(tag));
    }
    next()
  });
  


router.beforeEach((to, from, next) => {
    const nearestWithTitle = to.matched.slice().reverse().find(r => r.meta && r.meta.title);
    const nearestWithMeta = to.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);
    const previousNearestWithMeta = from.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);
  
    // If a route with a title was found, set the document (page) title to that value.
    if(nearestWithTitle) {
      document.title = nearestWithTitle.meta.title;
    } else if(previousNearestWithMeta) {
      document.title = previousNearestWithMeta.meta.title;
    }
  
    // Remove any stale meta tags from the document using the key attribute we set below.
    Array.from(document.querySelectorAll('[data-vue-router-controlled]')).map(el => el.parentNode.removeChild(el));
  
    // Skip rendering meta tags if there are none.
    if(nearestWithMeta) {
      // Turn the meta tag definitions into actual elements in the head.
      nearestWithMeta.meta.metaTags.map(tagDef => {
        const tag = document.createElement('meta');
  
        Object.keys(tagDef).forEach(key => {
          tag.setAttribute(key, tagDef[key]);
        });
  
        // We use this to track which meta tags we create so we don't interfere with other ones.
        tag.setAttribute('data-vue-router-controlled', '');
  
        return tag;
      })
      // Add the meta tags to the document head.
      .forEach(tag => document.head.appendChild(tag));
    }
    next()
  });
  

export default router;