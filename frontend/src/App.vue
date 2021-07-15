<template>
  <header>
    <Navbar />
  </header>
  <footer>
    <Footer />
  </footer>
</template>

<script>
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';

export default {
  name: 'App',
  components: {
    Navbar,
    Footer,
  },
  mounted() {
    if (document.cookie) {
      let userId = parseInt(document.cookie.slice(3));
      this.$store.dispatch('getCartByUserId',userId);
      this.$store.dispatch('setAuth', true);
    } else {
      this.$store.dispatch('setAuth', false);
    }
  },
};
</script>

<style>
* {
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
}

/* body {
  font-family: 'montserrat', sans-serif;
} */
</style>
