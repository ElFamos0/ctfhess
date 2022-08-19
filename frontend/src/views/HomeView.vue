<template>
  <v-container>
    <v-row class="mb-5 mt-5" justify="center" v-for="c in challenges" :key="c">
        <v-col cols="4" class="text-center" v-for="chall in c" :key="chall">
          <template v-if="!chall.fake" >
            <chall :chall="chall"> </chall>
          </template>
          <template v-else>
            <fakechall></fakechall>
          </template>
        </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
import chall from "@/components/Challenge.vue";
import fakechall from "@/components/FakeChallenge.vue";

export default {
  name: 'HomeView',
  components: {
    chall,
    fakechall,
  },
  data() {
    return {
      challenges: [],
    }
  },
  mounted() {
    getRequest('/chall/getall', 'json').then(data => {
      data = data.data
      if (data.success) {
        this.challenges = data.challenges
        console.log(this.challenges)
      }
    })
  },
};
</script>
