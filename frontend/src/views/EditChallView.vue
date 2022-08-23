<template>
  <v-container>
    <h2 class="mb-5">Edition de chall oklm</h2>
    <v-row justify="center">
      <v-col cols="6" class="text-center">
          <v-btn color="secondary" @click="delChall()">Supression du challenge</v-btn>
      </v-col>
    </v-row>
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
          v-model.number="form.day_open"
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
          v-model.number="form.points"
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

            <template v-for="p in form.pages.length+1" :key="p">
              <v-file-input v-model="files[p-1]" v-if="p == page+1" class="mb-2" chips multiple></v-file-input>
            </template>

            <v-checkbox
              v-model="form.pages[page].flag"
              label="Demande le flag ?"
            ></v-checkbox>

            <v-row justify="center">
              <template v-for="file in form.pages[page].files" :key="file">
                <v-col cols="2">
                    <v-tooltip location="top">
                        <template v-slot:activator="{ props }">
                            <div v-bind="props"  style="position:relative">
                              <v-btn v-bind="props" variant="text" class="delete" @click="deleteFile(file.uri)">
                                <v-icon large color="red darken-2">
                                    mdi-delete
                                </v-icon>
                              </v-btn>
                              <template v-if="file.name.split('.').pop() == 'jpeg' || file.name.split('.').pop() == 'jpg' || file.name.split('.').pop() == 'png'">
                                <v-img height="50px" :src="fileURI+file.uri" class="image" @click="image_dialog = true; image = file.uri"/>
                              </template>
                              <template v-else>
                                <v-spacer></v-spacer>
                                <v-icon large color="orange darken-2">
                                    mdi-file
                                </v-icon>
                              </template>
                            </div>
                        </template>
                        <span>{{file.name}}</span>
                    </v-tooltip>
                </v-col>
              </template>
            </v-row>
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
            Modifier le challenge
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
import { postRequest } from "@/requests/postRequest";
export default {
  name: 'ChallView',
  props: ['chall'],
  data() {
    return {
      form: {
        id: '',
        title: '',
        subtitle: '',
        short_description: '',
        day_open: 0,
        flag: '',
        points: 0,
        pages: [
          {
            title: '',
            description: '',
            flag: false,
          },
        ],
      },
      files : [[]],
      page: 0,
    }
  },
  mounted() {
    this.form = this.chall;
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
    delChall() {
      let data = {
        id: this.form.id,
      }
      postRequest(data, '/chall/delete', 'json').then(() => {
        this.$emit('close')
        this.$router.push('/')
      })
    },
    async validate() {
      await postRequest(this.form, '/chall/edit', 'json');

      let files = new FormData()
      for (let i = 0; i < this.files.length; i++) {
        let name = `files[${i}]`
        for (let file of this.files[i]) {
          files.append(name, file)
        }
      }

      await postRequest(files, '/file/reupload/'+this.chall.id, 'file')
      this.$router.push('/')
    },
    deleteFile(name) {
      getRequest('/file/delete/'+name, 'json').then(() => {
        window.location.reload()
      })
    }
  },
};
</script>