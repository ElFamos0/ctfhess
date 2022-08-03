import { createApp } from 'vue'
import App from './App.vue'
import {authStore} from './store/authStore'
import {router} from './router'


const app = createApp(App)
app.use(authStore).use(router).mount('#app')