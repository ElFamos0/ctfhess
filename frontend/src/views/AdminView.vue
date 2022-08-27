<template>
<v-container>
        <h1>ADMIN PAGE</h1>
        <v-card v-if="conn.user.type==0">
            <v-card-title>
                <h2>Voici un petit flag HTN{j3_suis_4dm1n?}</h2>
            </v-card-title>
            </v-card>
        <v-card v-if="conn.user.type==1">
            <v-card-title>Liste admin </v-card-title>
                <v-row class="mb-2 mt-2" justify="center" v-for="a in adminMembers" :key="a">
                    <h1>{{a.surname}} {{a.name}}</h1>
                </v-row>
        </v-card>
        <v-card v-if="conn.user.type==1">
            <v-card-title>Ajout ADMIN</v-card-title>
            <v-row class="mb-2 mt-2" justify="center">
                <v-col cols="12" sm="6" md="4">
                    <!-- On change execute add with e -->
                    <v-autocomplete
                        v-model="selectedAdmin"
                        :items="bruteusers"
                        :item-title="item => item.surname + ' ' + item.name"
                        return-object
                        label="User"
                        single-line
                        clearable
                        @click:clear ="deleteselected"
                        @input="add"
                        >
                    </v-autocomplete>
                </v-col>
            </v-row>
        </v-card>

        <v-btn v-if="conn.user.type==1" @click="addAdmin">Ajouter</v-btn>

</v-container>

</template>



<script>
import { getRequest } from "@/requests/getRequest";
import { postRequest } from "@/requests/postRequest";
import { loggedIn } from "@/auth/loggedIn";
export default {
    name: 'AdminView',
    data() {
        return {
            conn : {
            logged: false,
            user: {},
            },
            adminMembers: null,
            users: [],
            bruteusers: [],
            selectedAdmin: null,
        }
    },
    mounted() {
        getRequest('/users/admins').then((res) => {
            this.adminMembers = res.data;
        });
        getRequest('/users/listusers').then((res) => {
            this.bruteusers = res.data;
            }
        );
        loggedIn().then(data => {
        this.conn.logged = data.logged;
        this.conn.user = data.user;
        });
    },


    methods: {
        add(e) {
            console.log("add");
            this.selectedAdmin = e;
            console.log(this.selectedAdmin);
        },
        deleteselected() {
            this.selectedAdmin = null;
        },
        addAdmin() {
            console.log(this.selectedAdmin.id);
            postRequest({id:this.selectedAdmin.id}, '/users/makeadmin','json').then((res) => {
                console.log(res);
                this.$router.push('/');
            });
        },
    },
}
</script>