<template>
  <v-container>
    <notifications position="top center" classes="notif vue-notification"/>
    <v-row class="mb-5 mt-5" justify="center" v-for="c in challenges" :key="c">
        <v-col cols="3" class="text-center" v-for="chall in c" :key="chall">
          <template v-if="!chall.fake" >
            <chall :chall="chall" :conn="conn"> </chall>
          </template>
          <template v-else>
            <template v-if="chall.day_open == nextDay" >
              <fakechall :timer="timer"></fakechall>
            </template>
            <template v-else>
              <fakechall></fakechall>
            </template>
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
  props:["conn"],
  components: {
    chall,
    fakechall,
  },
  data() {
    return {
      challenges: [],
      polling: null,
      timer: "",
      timeleft: 0,
      nextDay:0,
    }
  },
  mounted() {
    getRequest('/chall/getall', 'json').then(data => {
      data = data.data
      if (data.success) {
        this.challenges = data.challenges
        console.log(data.challenges);
      }
      for (let i = 0; i < this.challenges.length; i++) {
        let broke = false
        for (let j = 0; j < this.challenges[i].length; j++) {
          if (this.challenges[i][j].fake) {
            this.nextDay = this.challenges[i][j].day_open
            broke = true
            break
          }
        }
        if (broke) {
          break
        }
      }
      console.log(this.nextDay)
    })
    getRequest('/timer/get').then((res) => {
      this.timeleft = res.data.timeleft;
      this.update();
      this.polling = setInterval(this.update, 1000);
    });
  },
  methods: {
      update() {
        this.timeleft -= 1
        let hours = Math.floor(this.timeleft / 3600);
        let minutes = Math.floor(this.timeleft%3600 / 60);
        let seconds = this.timeleft%3600 % 60;
        this.timer = `${hours}h ${minutes}m ${seconds}s`
      },
  },
  beforeUnmount () {
      clearInterval(this.polling)
  }
};
</script>
