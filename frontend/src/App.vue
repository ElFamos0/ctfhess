<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">
          <div style="display:inline-block;vertical-align:top;" class="mr-4">
            <img :src="require('@/assets/logo.png')" height="40"/>
          </div>
          <div style="display:inline-block;">
            {{navText}}
          </div>
        </v-tab>
      </template>

      <template v-slot:append>
        <template v-if="conn.user">
          <v-btn to="/admin" prepend-icon="mdi-badge-account-horizontal" v-if="conn.user.type==1">Admins</v-btn>
          <v-btn to="/edit" prepend-icon="mdi-pen" v-if="conn.user.type==1">Edition</v-btn>
          <v-btn to="/chall" prepend-icon="mdi-plus" v-if="conn.user.type==1">Creation</v-btn>
        </template>

        <template v-if="!conn.logged">
          <v-btn :href="logInURI" target="_self" prepend-icon="mdi-google">Connexion</v-btn>
        </template>
        <template v-else>
          <v-btn to="/scoreboard" prepend-icon="mdi-counter">Tableau des scores</v-btn>
          <v-btn to="/profil" prepend-icon="mdi-account">Profil</v-btn>
          <v-btn :href="logOutURI" target="_self" prepend-icon="mdi-google">Déconnexion</v-btn>
        </template>
      </template>
    </v-app-bar>

    <v-main>
      <router-view :conn="conn"></router-view>
    </v-main>

    <v-footer app fixed height="50pw">
    <v-col class="text-center" cols="12">
      <v-btn :href="gitlabURL" target="_self" icon variant="text">
        <v-icon>mdi-gitlab</v-icon>
      </v-btn>

      <strong>Fait avec ❤️ par Malo et Tristan</strong> 

      <v-tooltip location="top">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" href="https://discord.gg/zmP6p2xzGs" target="_self" icon variant="text">
            <v-icon>mdi-discord</v-icon>
          </v-btn>
        </template>
        <span>Hackin'TN</span>
      </v-tooltip>
      <v-tooltip location="top">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" href="https://discord.gg/MYKJgcCZXZ" target="_self" icon variant="text">
            <v-icon>mdi-discord</v-icon>
          </v-btn>
        </template>
        <span>Algo</span>
      </v-tooltip>
    </v-col>
  </v-footer>

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
      logInURI: `${process.env.VUE_APP_BACKEND_URI}/api/login`,
      logOutURI: `${process.env.VUE_APP_BACKEND_URI}/api/logout`,
      gitlabURL: "",
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

html{
    overflow-y: overlay;
}
</style>
