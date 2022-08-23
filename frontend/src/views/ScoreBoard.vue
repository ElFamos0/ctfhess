<template>
    <v-container>
        <h1>Tableau des scores</h1>
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
    </v-container>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
export default {
  name: "ScoreBoard",
  props: ["conn"],
  data() {
    return {
        users :[],
    }
  },
  async created () {
    await getRequest('/users/scoreboard','json').then((res) =>{
      this.users = res.data
    })
  },
};
</script>
