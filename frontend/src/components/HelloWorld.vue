<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <h1>Is initialized : {{ Vue3GoogleOauth.isInit }}</h1>
    <h1>Is authorized : {{ Vue3GoogleOauth.isAuthorized }}</h1>
    <h2>user : {{user}}</h2>

  </div>
</template>

<script>
import { inject } from 'vue'
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },

  data () {
    return {
      user: '',
    }
  },

  setup(){
    const Vue3GoogleOauth = inject('Vue3GoogleOauth')
    console.log(Vue3GoogleOauth)
    return {
      Vue3GoogleOauth
    };
  }
}
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
