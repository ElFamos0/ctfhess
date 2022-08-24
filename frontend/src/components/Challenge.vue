<template>    
    <v-card height="100%" tonal @click="click()" :class="'d-flex flex-column ' + color">
        <v-card-title>
            {{chall.title}}
        </v-card-title>
        <v-card-subtitle>
            <VueMarkdownIt breaks :source="chall.subtitle"></VueMarkdownIt>{{chall.points}} Points
        </v-card-subtitle>
        <v-card-subtitle>
            {{chall.completions}} personnes ont réussis ce challenge.
        </v-card-subtitle>
        <v-card-text>
            <VueMarkdownIt breaks :source="chall.short_description"></VueMarkdownIt>
        </v-card-text>
        <v-card-actions>
            <v-row justify="center">
                <v-col cols="6" class="text-center">
                    <p class="author">Proposé par : <b>{{chall.author}}</b></p>
                </v-col>
            </v-row>
        </v-card-actions>
    </v-card>
    <v-dialog v-model="dialog" persistent max-width="50vw">
        <v-card tonal v-if="dialog" width="50vw">
            <v-card-title>
                {{chall.pages[page].title}}
            </v-card-title>
            <v-card-text>
                <VueMarkdownIt breaks :source="chall.pages[page].description" class="mb-4"></VueMarkdownIt>
                
                <v-dialog v-model="image_dialog">
                    <v-card tonal v-if="dialog" width="500px">
                        <v-img :src="fileURI+image" class="image"/>
                    </v-card>
                </v-dialog>

                <v-container v-if="chall.pages[page].flag" class="mb-4">
                    <v-text-field v-model="flag" label="Flag" :rules="[(v) => !!v || 'Flag is required']"></v-text-field>
                    <v-btn @click="submitFlag()" color="secondary">Envoyer</v-btn>
                </v-container>

                <v-row justify="center" class="mb-1">
                    <template v-for="file in chall.pages[page].files" :key="file">
                        <v-col cols="2">
                            <v-tooltip location="top">
                                <template v-slot:activator="{ props }">
                                    <template v-if="file.name.split('.').pop() == 'jpeg' || file.name.split('.').pop() == 'jpg' || file.name.split('.').pop() == 'png'">
                                        <v-img height="50px" :src="fileURI+file.uri" class="image" @click="image_dialog = true; image = file.uri" v-bind="props"/>
                                    </template>
                                    <template v-else>
                                        <v-btn height="100%" :href="fileURI+file.uri" v-bind="props" variant="text">
                                            <v-icon large color="orange darken-2">
                                                mdi-file
                                            </v-icon>
                                        </v-btn>
                                    </template>
                                </template>
                                <span>{{file.name}}</span>
                            </v-tooltip>
                        </v-col>
                    </template>
                </v-row>

            </v-card-text>
            <v-card-actions>
                <v-btn color="secondary" @click="dialog = false">
                    Fermer
                </v-btn>
                <v-btn color="primary" @click="page--" :disabled="page <= 0">
                    Précédent
                </v-btn>
                <v-btn color="primary" @click="page++" :disabled="page == chall.pages.length-1">
                    Suivant
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import VueMarkdownIt from 'vue3-markdown-it';
import { postRequest } from "@/requests/postRequest";

export default {
  name: 'chall-c',
  components: {
    VueMarkdownIt,
  },
  props: ['chall', 'conn'],
  data() {
    return {
        dialog: false,
        image_dialog: false,
        image: '',
        page: 0,
        flag: undefined,
        color: '',
        logInURI: `${process.env.VUE_APP_BACKEND_URI}/api/login`,
        fileURI: process.env.VUE_APP_BACKEND_URI+'/api/file/get/',
    }
  },
  mounted() {
    if (this.conn.logged && this.conn.user.completions) {
        for (let i = 0; i < this.conn.user.completions.length; i++) {
            if (this.chall.id == this.conn.user.completions[i].chall_id) {
                this.color = 'bg-green-darken-4';
                break;
            }
        }
    }
  },
  methods: {
    click() {
        if (this.conn.logged) {
            this.dialog = true
        } else {
            window.open(this.logInURI, '_self')
        }
    },
    submitFlag() {
        let data = {
            id: this.chall.id,
            flag: this.flag,
        }
        postRequest(data, '/chall/complete', 'json').then((data) => {
            console.log(data.data)
            if (data.data.error) {
                this.$notify({
                    text:data.data.error,
                });
            } else {
                this.$notify({
                    text:`GG ! Tu as gagnés ${this.chall.points} points!`,
                });
                this.color = 'green';
                setTimeout(() => {
                    this.$router.go();
                }, 1000);
            }
        })
    },
  }
};
</script>

<style>
.author {
    color:blanchedalmond;
}

.image {
    background-color: beige;
    border-radius: 5px;
}

.notif {
  /* higher font size */
  font-size: 1.5em !important;
}
</style>