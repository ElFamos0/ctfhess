import { createApp } from 'vue'
import App from './App.vue'
import {authStore} from './store/authStore'
import router from './router'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'

loadFonts()

createApp(App)
  .use(authStore)
  .use(vuetify)
  .use(router)
  .mount('#app')
