<template>
    <v-card tonal>
        <v-overlay :model-value="blur" scroll-strategy="none" scrim="#000" contained class="align-center justify-center">
            <template v-if="timer">
                <h1>Ouvre dans :</h1>
                <h2>{{timer}}</h2>
            </template>
        </v-overlay>
        <v-container class="blur" @click="dialog = true">
            <v-card-title>
                HAHA BIEN ESSAYE
            </v-card-title>
            <v-card-subtitle>
                JE T'AI EU
            </v-card-subtitle>
            <v-card-text>
                ET ICI JE MET DU FAUX TEXTE COMME CA TA BIEN LE SEUM. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eu nunc egestas nisi placerat efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget nunc nec nibh consequat porta sit amet vitae orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ac tincidunt metus, in tempus sapien. Quisque molestie elementum mi nec commodo.
            </v-card-text>
        </v-container>
    </v-card>
    <v-dialog v-model="dialog" persistent max-width="50vw">
        <v-card tonal v-if="dialog" width="500px">
            <v-card-title>
                {{pages[page].title}}
            </v-card-title>
            <v-card-text>
                {{pages[page].text}}   
                
                <v-container v-if="pages[page].flag">
                    <v-text-field v-model="flag" label="Flag" :rules="[(v) => !!v || 'Flag is required']">
                    </v-text-field>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-btn color="secondary" @click="dialog = false">
                    Fermer
                </v-btn>
                <v-btn color="primary" @click="page--" :disabled="page <= 0">
                    Précédent
                </v-btn>
                <v-btn color="primary" @click="page++" :disabled="page >= pages.length-1">
                    Suivant
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
  name: 'chall-c',
  props: ["timer"],
  data() {
    return {
        dialog: false,
        page: 0,
        flag: undefined,
        blur: true,
        pages: [{
            title: 'Forceur',
            text: 'Je pensais pas que t\'allais le faire'
        }]
    }
  },
  mounted() {
    for (let i = 0; i < 50; i++) {
        this.pages.push({
            title: 'Forceur',
            text: 'Je pensais pas que t\'allais le faire'
        })
    }
    this.pages.push({
        title: 'Forceur',
        text: 'Je pensais pas que t\'allais le faire // ae0524c6-8bbb-44d3-aee0-fb7b1e4a5217'
    })
  },
};
</script>

<style scoped>
.blur {
    filter: blur(7px);
    user-select: none;
}
</style>