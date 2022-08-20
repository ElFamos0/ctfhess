<template>
    <v-container>
        <v-dialog v-model="dialog">
            <v-card v-if="dialog">
                <EditChallViewVue :chall="challenge" @close="dialog = false"></EditChallViewVue>
            </v-card>
        </v-dialog>
        <v-row class="mb-5 mt-5" justify="center" v-for="c in challenges" :key="c">
            <v-col cols="4" class="text-center" v-for="chall in c" :key="chall">
                <v-card tonal @click="dialog = true; challenge = chall;">
                    <v-card-title>
                        {{chall.title}}
                    </v-card-title>
                    <v-card-subtitle>
                        {{chall.subtitle}}
                    </v-card-subtitle>
                    <v-card-text>
                        {{chall.short_description}}
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
import EditChallViewVue from "./EditChallView.vue";

export default {
    name: 'HomeView',
    components: {
        EditChallViewVue,
    },
    data() {
        return {
            dialog: false,
            challenge: {},
            challenges: [],
        }
    },
    mounted() {
        getRequest('/chall/getalladmin', 'json').then(data => {
            data = data.data
            if (data.success) {
                this.challenges = data.challenges
                console.log(this.challenges)
            }
        })
    },
};
</script>
