import { createApp } from 'vue'
import App from './App.vue'
import gAuthPlugin from 'vue3-google-oauth2'
const app = createApp(App)

let gauthClientId = '130553832427-1oi6h98j0qkhlg3931b7bmcq65qho5jb.apps.googleusercontent.com';
app.use(gAuthPlugin, {
        clientId: gauthClientId,
        scope: 'email',
        prompt: 'consent',
})
app.mount('#app')