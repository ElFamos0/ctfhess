<template>
    <v-container>
        <h1>Tableau des scores</h1>
    </v-container> 
    <v-row justify="center"> 
        <v-col cols="4">
            <Suspense v-if="graphActivated">
                <graph :promo="promo" class="chart"></graph>
                <template #fallback>
                    Pas de données.
                </template>
            </Suspense>
        </v-col>
    </v-row> 
    <v-container>

    <v-row justify="center"> 
        <v-col cols="6">
            <v-btn :disabled="inte" @click="swap()">Classement 1A</v-btn>
        </v-col>
        <v-col cols="6">
            <v-btn :disabled="general" @click="swap()">Classement Général</v-btn>
        </v-col>
    </v-row> 
    </v-container>
    <v-container>
        <v-table>
            <thead>
                <tr>
                    <th class="text-center">
                        Prénom
                    </th>
                    <th class="text-center">
                        Points
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in users" :key="user" :class="user.id == conn.user.id ? 'bg-orange-lighten-4' : ''">
                    <td>{{ user.name }}</td>
                    <td>{{ user.points }}</td>
                </tr>
            </tbody>
        </v-table>
        <p v-if="users.length == 0">Aucun utilisateurs..</p>
    </v-container>
</template>

<script>
import Graph from "@/components/Graph.vue";
import { getRequest } from "@/requests/getRequest";
export default {
  name: "ScoreBoard",
  props: ["conn"],
  components: {
    Graph,
  },
  data() {
    return {
        users :[],
        general: false,
        inte: true,
        graphActivated: true,
        promo: "1a",
    }
  },
  async created () {
    await getRequest('/users/scoreboard/'+this.promo,'json').then((res) =>{
        if (!res.data) {
            this.users = []
        } else {
            this.users = res.data
        }
    })
  },
  methods: {
    async swap() {
        this.graphActivated = false
        this.general = !this.general;
        this.inte = !this.inte;
        if (this.general) {
            this.promo = "all";
            await getRequest('/users/scoreboard/'+this.promo,'json').then((res) =>{
                if (!res.data) {
                    this.users = []
                } else {
                    this.users = res.data
                }
            })
        } else {
            this.promo = "1a";
            await getRequest('/users/scoreboard/'+this.promo,'json').then((res) =>{
                if (!res.data) {
                    this.users = []
                } else {
                    this.users = res.data
                }
            })
        }
        this.graphActivated = true
    },
  },
};
</script>

<style scoped>
.chart {
    max-height: 30vh;
    max-width: 40vw;
}
</style>