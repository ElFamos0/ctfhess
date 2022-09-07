<template>
    <v-container>
            <v-card>
                <v-card-title>
                    <h2>Timer du backend (c'est pour vérifier que le timer interne est bon)</h2>
                </v-card-title>
                <v-card-text>
                    <v-row>
                        <v-col cols="6">
                            <h3 class="mb-5">Temps écoulé depuis le début :</h3>
                            <h4 class="mt-5">{{timer}}</h4>
                        </v-col>
                        <v-col cols="6">
                            <h3 class="mb-5">Temps restant avant le jour suivant :</h3>
                            <h4 class="mt-5">{{timeleft}}</h4>
                        </v-col>
                    </v-row>
                </v-card-text>
            </v-card>
    </v-container>
</template>
    
    
    
<script>
import { getRequest } from "@/requests/getRequest";
export default {
    name: 'AdminView',
    data() {
        return {
            timerS: 0,
            timer: "",
            timerleftS: 0,
            timeleft: "",
            polling: null,
        }
    },
    mounted() {
        getRequest('/timer/get').then((res) => {
            this.timerS = res.data.timer;
            this.timeleftS = res.data.timeleft;

            this.update();
            this.polling = setInterval(this.update, 1000);
        });
    },
    methods: {
        update() {
            this.timeleftS -= 1
            this.timerS += 1

            this.timer = this.formatTime(this.timerS)
            this.timeleft = this.formatTime(this.timeleftS)
        },
        formatTime(t) {
            let days = Math.floor(t / 86400);
            let hours = Math.floor(t / 3600)%24;
            let minutes = Math.floor(t%3600 / 60);
            let seconds = t%3600 % 60;
            return `${days} jours ${hours} heures ${minutes} minutes ${seconds} secondes`
        }
    },
    beforeUnmount () {
        clearInterval(this.polling)
    }
}
</script>