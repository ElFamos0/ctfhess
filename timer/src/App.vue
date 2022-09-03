<template>
  <v-app>
    <v-main>
      <video-bg class="video" :sources="[require('./assets/video.mp4')]">
        <v-container class="centered" v-if="secondes > 0">
          <v-row justify="center">
            <v-col cols="12">
              <h1 class="title text-center">
                Rendez-vous dans :
              </h1>
            </v-col>
          </v-row>
          <v-row justify="center" class="margin">
            <v-col cols="3">
              <v-card variant="plain" style="background: #0000;">
                <v-card-text class="text-center card">
                  {{jours}}
                </v-card-text>
                <v-card-text class="text-center subtext">
                  jours
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="3">
              <v-card variant="plain" style="background: #0000;">
                <v-card-text class="text-center card">
                  {{heures}}
                </v-card-text>
                <v-card-text class="text-center subtext">
                  heures
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="3">
              <v-card variant="plain" style="background: #0000;">
                <v-card-text class="text-center card">
                  {{minutes}}
                </v-card-text>
                <v-card-text class="text-center subtext">
                  minutes
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="3">
              <v-card variant="plain" style="background: #0000;">
                <v-card-text class="text-center card">
                  {{secondes}}
                </v-card-text>
                <v-card-text class="text-center subtext">
                  secondes
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
        <v-container class="centered" v-else>
          <v-row justify="center">
            <v-col cols="12">
              <h1 class="title text-center">
                Encore un peu de patience...
              </h1>
            </v-col>
          </v-row>
        </v-container>
      </video-bg>
    </v-main>
  </v-app>
</template>

<script>
import VideoBg from 'vue-videobg'

export default {
  name: 'App',

  components: {
    VideoBg
  },

  data: () => ({
    fin: new Date(process.env.VUE_APP_DATE),
    jours: 5,
    heures: 0,
    minutes: 0,
    secondes: 5,
  }),

  mounted() {
    this.start();
  },

  methods: {
    start() {
      setInterval(() => {
        let now = new Date();
        this.jours = Math.floor((this.fin - now) / (1000 * 60 * 60 * 24));
        this.heures = Math.floor((this.fin - now) / (1000 * 60 * 60)) % 24;
        this.minutes = Math.floor((this.fin - now) / (1000 * 60)) % 60;
        this.secondes = Math.floor((this.fin - now) / 1000) % 60;
      }, 1000);
    }
  }
}
</script>

<style>
video {
  /* blur */
  filter: blur(5px);
}

.centered {
  position: fixed;
  top: 45%;
  left: 50%;
  /* bring your own prefixes */
  transform: translate(-50%, -50%);
  width: 50vw;
}

.card {
  font-size: 7vh;
  color:white;
  text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000;


  transition: opacity 1s;
}

.title {
  font-size: 5vh;
  color:white;
  text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000;
}

.subtext {
  color:white;
  text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000;
}

.margin {
  margin-top:50px;
}
</style>