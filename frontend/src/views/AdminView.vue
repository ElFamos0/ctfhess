<template>
    <v-dialog v-model="dialog">
        <v-card tonal width="75vw">
            <SmallProfileVue :user="profileUser"></SmallProfileVue>
        </v-card>
    </v-dialog>
    <v-container>
        <v-card v-if="conn.user.type==0">
            <v-card-title>
                <h2>Gestions des administrateurs</h2>
            </v-card-title>
            <v-card-title>
                <h4>Voici un petit flag HTN{j3_suis_4dm1n?}</h4>
            </v-card-title>
        </v-card>
        <v-card>
            <v-card-title>
                <h2>Gestions des administrateurs</h2>
            </v-card-title>
            <v-card-text>
                <v-row>
                    <v-col cols="6">
                        <h3 class="mb-5">Admins</h3>
                        <v-row class="mb-2 mt-2" justify="center" v-for="a in adminMembers" :key="a">
                            <h4 @click="open(a)" class="hover-click">{{a.surname}} {{a.name}}</h4>
                            <v-icon large color="red darken-2" @click="remAdmin(a)">
                                mdi-delete
                            </v-icon>
                        </v-row>
                    </v-col>
                    <v-col cols="6">
                        <h3 class="mb-5">Utilisateurs</h3>
                        <v-row class="mb-2 mt-2" justify="center" v-for="u in bruteusers" :key="u">
                            <v-icon large color="green darken-2" @click="addAdmin(u)">
                                mdi-arrow-left
                            </v-icon>
                            <h4 @click="open(u)" class="hover-click">{{u.surname}} {{u.name}}</h4>
                        </v-row>
                    </v-col>
                </v-row>
            </v-card-text>
        </v-card>
    </v-container>
</template>



<script>
import { getRequest } from "@/requests/getRequest";
import { postRequest } from "@/requests/postRequest";
import SmallProfileVue from "@/components/SmallProfile.vue";
export default {
    name: 'AdminView',
    components: {
        SmallProfileVue,
    },
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
            profileUser: null,
            dialog: false,
        }
    },
    mounted() {
        this.update();
    },
    methods: {
        update() {
            getRequest('/admin/get').then((res) => {
                this.adminMembers = res.data;
            });
            getRequest('/users/list').then((res) => {
                this.bruteusers = res.data;
                }
            );
        },
        addAdmin(user) {
            postRequest({id:user.id}, '/admin/add','json').then(() => {
                this.update();
            });
        },
        remAdmin(user) {
            postRequest({id:user.id}, '/admin/remove','json').then(() => {
                this.update();
            });
        },
        open(u) {
            this.dialog = true
            this.profileUser = u
        }
    },
}
</script>

<style scoped>
.hover-click:hover {
    cursor: pointer;
}
</style>