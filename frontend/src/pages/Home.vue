<template>
  <Banner :images="images"/>
  <FeaturedProduct />
  <Feedback />
</template>

<script>
import Banner from "../components/Banner.vue";
import FeaturedProduct from "../components/FeaturedProducts.vue";
import Feedback from "../components/Feedback.vue";

export default {
  name: "Home",
  components: {
    Banner,
    FeaturedProduct,
    Feedback,
  },
  data() {
    return {
      images: [
        {id:1, name:"slider-image-1-1920x600.jpg"},
        {id:2, name:"slider-image-2-1920x600.jpg"},
        {id:3, name:"slider-image-3-1920x600.jpg"},
      ],
    };
  },
};
</script>

<style></style>
