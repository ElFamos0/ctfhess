<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">{{navText}}</v-tab>
      </template>

      <template v-slot:append>
        <v-btn to="/login" prepend-icon="mdi-google">Connexion</v-btn>
        <v-btn to="/profil" prepend-icon="mdi-account">Profil</v-btn>
      </template>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
  <!-- <HomeView v-if="!loggedIn" /> -->
  <!-- <LoginView v-else /> -->
</template>

<script>
import { authStore } from "./store/authStore";
// import LoginView from "./views/LoginView.vue";
// import HomeView from "./views/HomeView.vue";

export default {
  name: "App",
  data() {
    return {
      loggedIn: authStore.getters.loggedIn,
      accessToken: authStore.getters.accessToken,
      navText:"",
    }; 
  },

  mounted() {
    let navTexts = [
      "Essaie donc de me compromettre",
      "Compromet dans la boîte",
      "Enracine moi",
      "Sanctuaire de vulnérabilitées",
      "Compromet ma machine virtuelle",
      "L'heure de capturer le drapeau",
      "Prime à l'insecte",
      "La nuit de la compromission"
    ];

    this.navText = navTexts[Math.floor(Math.random() * navTexts.length)];
  }
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
