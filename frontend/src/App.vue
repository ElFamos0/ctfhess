<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">{{navText}}</v-tab>
      </template>

      <template v-slot:append>
        <v-btn to="/edit" prepend-icon="mdi-pen" v-if="conn.user.type==1">Edition</v-btn>
        <v-btn to="/chall" prepend-icon="mdi-plus" v-if="conn.user.type==1">Creation</v-btn>

        <v-btn :href="logInURI" target="_self" prepend-icon="mdi-google" v-if="!conn.logged">Connexion</v-btn>
        <v-btn :href="logOutURI" target="_self" prepend-icon="mdi-google" v-else>Déconnexion</v-btn>

        <v-btn to="/profil" prepend-icon="mdi-account" v-if="conn.logged">Profil</v-btn>
      </template>
    </v-app-bar>

    <v-main>
      <router-view :conn="conn"></router-view>
    </v-main>
  </v-app>
  <!-- <HomeView v-if="!loggedIn" /> -->
  <!-- <LoginView v-else /> -->
</template>

<script>
import { loggedIn } from "@/auth/loggedIn";
// import LoginView from "./views/LoginView.vue";
// import HomeView from "./views/HomeView.vue";

export default {
  name: "App",
  data() {
    return {
      conn : {
        logged: false,
        user: {},
      },
      logInURI: `http://${process.env.VUE_APP_BACKEND_DOMAIN}:${process.env.VUE_APP_BACKEND_PORT}/api/login`,
      logOutURI: `http://${process.env.VUE_APP_BACKEND_DOMAIN}:${process.env.VUE_APP_BACKEND_PORT}/api/logout`,
      navText: '',
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
      "La nuit de la compromission",
      "Base de données de compromission",
    ];

    this.navText = navTexts[Math.floor(Math.random() * navTexts.length)];

    loggedIn().then(data => {
      this.conn.logged = data.logged;
      this.conn.user = data.user;
    });
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
