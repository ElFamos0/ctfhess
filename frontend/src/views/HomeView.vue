<template>
  <NavBarVue />

</template>

<script>
import { defineComponent } from 'vue';

// Components
import NavBarVue from '@/components/NavBar.vue';
export default defineComponent({
  name: 'HomeView',

  components: {
    NavBarVue,
  },
});
</script>
