<template>
    <div>
        <h1>Profil</h1>
        <h2>{{ user }}</h2>
    </div>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
export default {
  name: "ProfilView",
  data() {
    return {
      user: [],
    };
  },
  async created () {
    await getRequest('/users/data','json').then((res) =>{
      this.user = res.data
    })
  },
};
</script>
