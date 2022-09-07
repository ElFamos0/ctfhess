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
    import { postRequest } from "@/requests/postRequest";
    export default {
        name: 'AdminView',
        data() {
            return {
                timer: "",
                timeleft: "",
            }
        },
        mounted() {
            setInterval(this.update, 1000);
        },
        methods: {
            update() {
                getRequest('/admin/timer/get').then((res) => {
                    this.timer = res.data.timer;
                    this.timeleft = res.data.timeleft;
                });
            },
        },
    }
    </script>