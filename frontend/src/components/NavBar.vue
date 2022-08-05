<template>
<v-toolbar color="indigo" dark fixed app>
        <v-toolbar-title>Challenge SWAG</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items class="hidden-sm-and-down">
          <v-btn v-for="item in menu" @click="this.$router.push('/profil')" :key="item.icon" :to="item.link" flat>{{
            item.title
          }}</v-btn>
        </v-toolbar-items>
      </v-toolbar>
</template>

<script>
export default {
  data() {
    return {
      menu: [
        { icon: "Profil", title: "Profil" },
        { icon: "Scores", title: "Scores" },
        { icon: "Extra", title: "Extra" }
      ]
    };
  },

  methods: {
    menuItems() {
      return this.menu;
    }
  }
};
</script>