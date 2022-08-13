<template>
    <v-card tonal @click="dialog = true">
        <v-card-title>
            BIENVENUE SUR LE HAKING
        </v-card-title>
        <v-card-subtitle>
            WEB
        </v-card-subtitle>
        <v-card-text>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eu nunc egestas nisi placerat efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget nunc nec nibh consequat porta sit amet vitae orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ac tincidunt metus, in tempus sapien. Quisque molestie elementum mi nec commodo. Phasellus consectetur justo felis, at cursus risus tempus vitae. Phasellus eu metus turpis. Vestibulum mi felis, aliquam nec elit ut, rhoncus tempus erat. Maecenas a tincidunt magna. Nulla massa sapien, convallis vel nibh at, luctus vulputate erat. Interdum et malesuada fames ac ante ipsum primis in faucibus.
        </v-card-text>
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
  props: ['data'],
  data() {
    return {
        dialog: false,
        page: 0,
        flag: undefined,
        pages: [{
            title: 'Page 1 test',
            text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eu nunc egestas nisi placerat efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget nunc nec nibh consequat porta sit amet vitae orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ac tincidunt metus, in tempus sapien. Quisque molestie elementum mi nec commodo. Phasellus consectetur justo felis, at cursus risus tempus vitae. Phasellus eu metus turpis. Vestibulum mi felis, aliquam nec elit ut, rhoncus tempus erat. Maecenas a tincidunt magna. Nulla massa sapien, convallis vel nibh at, luctus vulputate'
        },
        {
            title: 'Page 2 test',
            text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eu nunc egestas nisi placerat efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget nunc nec nibh consequat porta sit amet vitae orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ac tincidunt metus, in tempus sapien. Quisque molestie elementum mi nec commodo. Phasellus consectetur justo felis, at cursus risus tempus vitae. Phasellus eu metus turpis. Vestibulum mi felis, aliquam nec elit ut, rhoncus tempus erat. Maecenas a tincidunt magna. Nulla massa sapien, convallis vel nibh at, luctus vulputate'
        },
        {
            title: 'Page 3 test',
            text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eu nunc egestas nisi placerat efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget nunc nec nibh consequat porta sit amet vitae orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ac tincidunt metus, in tempus sapien. Quisque molestie elementum mi nec commodo. Phasellus consectetur justo felis, at cursus risus tempus vitae. Phasellus eu metus turpis. Vestibulum mi felis, aliquam nec elit ut, rhoncus tempus erat. Maecenas a tincidunt magna. Nulla massa sapien, convallis vel nibh at, luctus vulputate'
        },
        {
            title: 'Page 4 test',
            text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eu nunc egestas nisi placerat efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget nunc nec nibh consequat porta sit amet vitae orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ac tincidunt metus, in tempus sapien. Quisque molestie elementum mi nec commodo. Phasellus consectetur justo felis, at cursus risus tempus vitae. Phasellus eu metus turpis. Vestibulum mi felis, aliquam nec elit ut, rhoncus tempus erat. Maecenas a tincidunt magna. Nulla massa sapien, convallis vel nibh at, luctus vulputate',
            flag: true,
        }]
    }
  },
};
</script>