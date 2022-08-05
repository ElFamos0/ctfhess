<template>
    <h2>user : {{ user }}</h2>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
export default {
  name: "HelloWORLDView",
  props: {
    msg: String,
  },

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
