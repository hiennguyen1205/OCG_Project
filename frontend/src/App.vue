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
import { mapActions, mapMutations } from "vuex";

export default {
  name: "App",
  components: {
    Navbar,
    Footer,
  },
  methods: {
    ...mapActions("carts", ["getCartByUserId"]),
    ...mapMutations("users", ["setAuth"]),
  },

  mounted() {
    console.log("get cookie");
    if (document.cookie) {
      console.log("cookie",document.cookie);
      let userId = parseInt(document.cookie.slice(3));
      this.getCartByUserId(userId);
      this.setAuth(true);
    } else {
      this.setAuth(false);
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
