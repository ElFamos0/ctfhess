<template>
  <div v-if="loggedIn">
  <HomeView />
  </div>
  <div v-else>
    <LoginView />
  </div>
  <router-view/>
</template>

<script>
import { authStore } from "./store/authStore";
import LoginView from "./views/LoginView.vue";
import HomeView from "./views/HomeView.vue";

export default {
  name: "App",
  components: {
    LoginView,
    HomeView,
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
  margin-left: auto;
  margin-right: auto;
}
</style>
