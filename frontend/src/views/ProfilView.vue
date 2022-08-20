<template>
  <v-container>
    <h1>Profil de {{user.name}}</h1>
    <h4>Promotion : {{user.promo}}</h4>
    <h4>Points : {{user.points}}</h4>
  </v-container>
  <v-row justify="center">
    <v-col cols="6">
      <v-card tonal>
          <v-card-title>
              Challenges réalisé :
          </v-card-title>
          <v-card-text>
          <v-row justify="center" v-for="c in user.completions" :key="c" >
            <v-col cols="6">
              <v-card tonal class="bg-pink-lighten-3">
                  <v-card-title>
                      {{c.challenge.title}}
                  </v-card-title>
                  <v-card-subtitle class="mb-2">
                      <VueMarkdownIt :source="c.challenge.subtitle"></VueMarkdownIt>{{c.challenge.points}} Points
                  </v-card-subtitle>
              </v-card>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
import VueMarkdownIt from 'vue3-markdown-it';
export default {
  name: "ProfilView",
  props: ['conn'],
  components: {
    VueMarkdownIt,
  },
  data() {
    return {
      user: [],
    };
  },
  async created () {
    await getRequest('/users/data','json').then((res) =>{
      this.user = res.data
    })
  },
};
</script>
