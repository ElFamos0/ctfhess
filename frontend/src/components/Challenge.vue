<template>
    <v-card tonal @click="dialog = true">
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
    <v-dialog v-model="dialog" persistent max-width="50vw">
        <v-card tonal v-if="dialog" width="500px">
            <v-card-title>
                {{chall.pages[page].title}}
            </v-card-title>
            <v-card-text>
                {{chall.pages[page].description}}
                <v-container v-if="chall.pages[page].flag">
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
                <v-btn color="primary" @click="page++" :disabled="page == chall.pages.length-1">
                    Suivant
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
  name: 'chall-c',
  props: ['chall'],
  data() {
    return {
        dialog: false,
        page: 0,
        flag: undefined,
    }
  },
};
</script>