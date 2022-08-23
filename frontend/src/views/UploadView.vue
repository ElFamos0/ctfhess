<template>
<v-container>
        <h1>Upload file</h1>
        <v-file-input class="mb-2" multiple @change="addFile"></v-file-input>
        <v-btn class="mb-2" color="grey lighten-3" @click="uploadFile">Upload</v-btn>

    </v-container>

</template>

<script>
import { postRequest } from "@/requests/postRequest";
export default {
    name: 'UploadView',
    data() {
        return {
            form: {
                file: null,
            },
        }
    },
    methods: {
        uploadFile() {
            console.log("uploadFile")
            console.log(this.form.file)
            postRequest(this.form, '/file/upload', 'file').then(() => {
                this.$router.push('/')
            })
        },
        addFile(e) {
            // add file to form
            this.form.file = e.target.files[0]
        },
    },
};

</script>