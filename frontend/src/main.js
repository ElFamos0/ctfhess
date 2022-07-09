import { createApp } from 'vue'
import { createStore } from 'vuex'
import auth from '@/api/store'
import App from '@/App.vue'
import router from './router' 
import BootstrapVue3 from 'bootstrap-vue-3'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css'
import vuetify from './plugins/vuetify'

const store = createStore({
        modules: {
          auth
        }
      })
      
      createApp(App)
        .use(BootstrapVue3)
        .use(router)
        .use(store)
        .use(vuetify)
        .use(VueConfetti)
        .component('font-awesome-icon', FontAwesomeIcon)
        .mount('#app')
      