<template>
  <v-container>
    <v-row class="mb-5 mt-5">
      <v-col cols="4" class="text-center">
        <!-- empty -->
      </v-col>
      <v-col cols="4" class="text-center">
        <chall> </chall>
      </v-col>
      <v-col cols="4" class="text-center">
        <!-- empty -->
      </v-col>
    </v-row>
    <v-row class="mb-5 mt-5">
      <v-col cols="4" class="text-center">
        <fakechall> </fakechall>
      </v-col>
      <v-col cols="4" class="text-center">
        <fakechall> </fakechall>
      </v-col>
      <v-col cols="4" class="text-center">
        <fakechall> </fakechall>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import chall from "@/components/Challenge.vue";
import fakechall from "@/components/FakeChallenge.vue";

export default {
  name: 'HomeView',
  components: {
    chall,
    fakechall,
  },
  data() {
    return {}
  },
};
</script>
