<template>
  <header>
    <Navbar />
  </header>
  <footer>
    <Footer />
  </footer>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import Footer from "./components/Footer.vue";

export default {
  name: "App",
  components: {
    Navbar,
    Footer,
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
