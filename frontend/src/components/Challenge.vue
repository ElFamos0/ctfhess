<template>    
    <v-card height="100%" tonal @click="click()" :color="color">
        <v-card-title>
            {{chall.title}}
        </v-card-title>
        <v-card-subtitle>
            <VueMarkdownIt :source="chall.subtitle"></VueMarkdownIt>{{chall.points}} Points
        </v-card-subtitle>
        <v-card-text>
            <VueMarkdownIt :source="chall.short_description"></VueMarkdownIt>
        </v-card-text>
    </v-card>
    <v-dialog v-model="dialog" persistent max-width="50vw">
        <v-card tonal v-if="dialog" width="500px">
            <v-card-title>
                {{chall.pages[page].title}}
            </v-card-title>
            <v-card-text>
                <VueMarkdownIt :source="chall.pages[page].description"></VueMarkdownIt>
                <v-container v-if="chall.pages[page].flag">
                    <v-text-field v-model="flag" label="Flag" :rules="[(v) => !!v || 'Flag is required']"></v-text-field>
                    <v-btn @click="submitFlag()" color="secondary">Envoyer</v-btn>
                </v-container>
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
        page: 0,
        flag: undefined,
        color: '',
        logInURI: `http://${process.env.VUE_APP_BACKEND_DOMAIN}:${process.env.VUE_APP_BACKEND_PORT}/api/login`,
    }
  },
  mounted() {
    if (this.conn.logged && this.conn.user.completions) {
        for (let i = 0; i < this.conn.user.completions.length; i++) {
            if (this.chall.id == this.conn.user.completions[i].chall_id) {
                this.color = 'green';
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
            }
        })
    },
  }
};
</script>

<style>
.notif {
  /* higher font size */
  font-size: 1.5em !important;
}
</style>