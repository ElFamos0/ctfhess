<template>
  <div v-if="loggedIn">
    <img alt="Vue logo" src="./assets/logo.png" />
    <HelloWorld :msg="'Tema la taille du test'" />
  </div>
  <div v-else>
    <LoginView />
  </div>
</template>

<script>
import { authStore } from "./store/authStore";
import HelloWorld from "./components/HelloWorld.vue";
import LoginView from "./views/LoginView.vue";

export default {
  name: "App",
  components: {
    HelloWorld,
    LoginView,
  },
  data() {
    return {
      loggedIn: authStore.getters.loggedIn,
    };
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
