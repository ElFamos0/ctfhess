import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import Notifications from '@kyvg/vue3-notification'

loadFonts()

createApp(App)
  .use(vuetify)
  .use(router)
  .use(Notifications)
  .mount('#app')
