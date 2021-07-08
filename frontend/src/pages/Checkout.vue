<template>
  <div id="app">
    <CartListProducts />
    <CartPayment />
  </div>
</template>

<script>
import CartListProducts from "@/components/CartListProducts.vue";
import CartPayment from "@/components/CartPayment.vue";

export default {
  name: "Checkout",

  components: {
    CartListProducts,
    CartPayment,
  },

  data() {
    return {
      

      
      
    };
  },

  

  methods: {
    
    
    
  },
};
</script>

<style>
* {
  box-sizing: border-box;
}

body {
  margin: 20px 0;
  padding: 0;
  font-family: arial, sans-serif;
  overflow: scroll;
}

.container {
  width: 90%;
  margin: 0 auto;
  overflow: auto;
}

/* --- SMALL SCREEN --- */
@media all and (max-width: 599px) {
  .thumbnail img {
    display: none;
  }

  .quantity > input {
    width: 40px;
    height: 40px;
    left: calc(50% - 20px);
  }
}



/* --- LARGE SCREEN --- */
@media all and (min-width: 992px) {
  html {
    font-size: 16px;
  }
}
</style>
