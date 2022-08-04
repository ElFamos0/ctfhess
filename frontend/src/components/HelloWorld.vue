<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>user : {{ user }}</h2>
  </div>
</template>

<script>
import { getRequest } from "@/requests/getRequest";
export default {
  name: "HelloWorld",
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
