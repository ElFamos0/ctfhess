<template>
  <v-container>
    <h2 class="mb-5">Création de chall oklm</h2>
    <v-row justify="center">
      <v-col cols="6" class="text-center">
          <v-text-field
            v-model="form.title"
            label="Titre du challenge"
          ></v-text-field>
      </v-col>
      <v-col cols="6" class="text-center">
          <v-text-field
            v-model="form.subtitle"
            label="Sous-titre / Catégorie du challenge"
          ></v-text-field>
      </v-col>
      <v-col cols="6" class="text-center">
          <v-textarea
            v-model="form.short_description"
            label="Petite description"
          ></v-textarea>
      </v-col>
      <v-col cols="6" class="text-center">
        <v-text-field
          v-model="form.day_open"
          type="number"
          density="compact"
          hide-details
          variant="outlined"
          label="Jour d'ouverture"
        ></v-text-field>
      </v-col>
      <v-col cols="6" class="text-center">
          <v-text-field
            v-model="form.flag"
            label="Flag du challenge"
          ></v-text-field>
      </v-col>
      <v-col cols="6" class="text-center">
        <v-text-field
          v-model="form.points"
          type="number"
          density="compact"
          hide-details
          variant="outlined"
          label="Points"
        ></v-text-field>
      </v-col>
      <v-col cols="3" class="text-center">
      </v-col>
      <v-col cols="6" class="text-center">
        <v-card tonal>
          <v-card-title>
            Page {{page+1}}/{{form.pages.length}}
            <v-text-field
              v-model="form.pages[page].title"
              label="Titre de la page"
            ></v-text-field>
          </v-card-title>
          <v-card-text>
            <v-textarea
              v-model="form.pages[page].description"
              label="Description de la page"
            ></v-textarea>
            
            <v-checkbox
              v-model="form.pages[page].flag"
              label="Demande le flag ?"
            ></v-checkbox>
          </v-card-text>
          <v-card-actions>
              <v-btn color="primary" @click="delPage()" :disabled="form.pages.length <= 1">
                  Supprimer
              </v-btn>
              <v-btn color="primary" @click="page--" :disabled="page <= 0">
                  Précédent
              </v-btn>
              <v-btn color="primary" @click="addPage()">
                  Suivant
              </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col cols="3" class="text-center">
      </v-col>

      <v-col cols="6" class="text-center">
        <v-btn color="primary" @click="validate()">
            Créer le challenge
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { postRequest } from "@/requests/postRequest";
export default {
  name: 'ChallView',
  data() {
    return {
      form: {
        title: '',
        subtitle: '',
        short_description: '',
        day_open: '',
        flag: '',
        points: '',
        pages: [
          {
            title: '',
            description: '',
            flag: false,
          },
        ],
      },
      page: 0,
    }
  },
  methods: {
    addPage() {
      if (this.page == this.form.pages.length - 1) {
        this.form.pages.push({})
      }
      this.page++
    },
    delPage() {
      this.form.pages.splice(this.page, 1)
      if (this.page != 0) {
        this.page--
      }
    },
    validate() {
      postRequest(this.form, '/api/chall/create', 'json').then(() => {
        this.$router.push('/')
      })
    },
  },
};
</script>
