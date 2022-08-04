<template>
  <div v-if="loggedIn">
  <v-img src="./assets/logo.png"
  max-height="200"
  max-width="250"
  class="mx-auto">
  </v-img>
  
  </div>
  <div v-else>
    <LoginView />
  </div>
</template>

<script>
import { authStore } from "./store/authStore";
import LoginView from "./views/LoginView.vue";

export default {
  name: "App",
  components: {
    LoginView,
  },
  data() {
    return {
      loggedIn: authStore.getters.loggedIn,
      accessToken: authStore.getters.accessToken
    }; 
    },
  }
</script>

<style lang="scss">
:root {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #362c50;
  margin-top: 60px;
  margin-left: auto;
  margin-right: auto;
}
</style>
