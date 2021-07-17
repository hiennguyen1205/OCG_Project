<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <div class="section-heading">
          <h2>Sản phẩm được mua nhiều nhất</h2>
          <router-link to="/products/armchair">Xem thêm<i class="fa fa-angle-right"></i></router-link>
        </div>
      </div>
    </div>

    <div class="row">
      <div
        class="product-item col-md-4 "
        v-for="product in products"
        :key="product.id"
      >
        <div :class="{ discount: product.sale }">
          {{ product.Sale > 0 ? product.sale + "%" : "" }}
        </div>
        <img :src="`http://localhost:3000/${product.image}`" alt="hihi" />
        <div>
          <h4>{{ product.name }}</h4>
          <p>
            {{ product.description }}
          </p>
          <div style="height: 40px"></div>
          <div class="price">
            <p id="style-price">
              {{ formatCurrency(product.price) }}
              <button>
                <i
                  id="icon-cart"
                  class="fas fa-shopping-cart"
                  style="padding: 5px"
                ></i
                >Mua Hàng
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "FeaturedProduct",
  data() {
    return {
      products: [],
    };
  },

  async created() {
    const response = await fetch("http://localhost:3000/api/products?limit=6&cursor=0&isFeature=1");
    this.products = await response.json();
    // console.log(this.products);
  },
  methods: {
    formatCurrency(money) {
      return money.toLocaleString("vi", { style: "currency", currency: "VND" });
    },
  },
};
</script>

<style scoped>
.row {
  justify-content: space-between;
}
.section-heading {
  display: flex;
  justify-content: space-between;
  margin: 70px 0px;
  border-bottom: 1px solid rgb(230, 229, 229);
}
.section-heading a {
  font-size: 25px;
  text-decoration: none;
  color: #f33f3f;
}

.product-item {
  position: relative;
  border: 1px solid #eee;
  margin-bottom: 30px;
  width: 32%;
}

.col-md-4 {
  padding: 0px;
}

.product-item img {
  width: 100%;
  height: auto;
}

.product-item p {
  padding: 5px;
  text-align: justify;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}
.product-item h4 {
  padding: 5px;
}
.product-item {
  transition: transform 0.3s ease-out;
}
.product-item:hover {
  transform: scale(1.01);
  border: 1px solid rgb(107, 240, 206);
}

.price button {
  font-family: lato, sans-serif;
  font-weight: bold;
  font-size: 1em;
  letter-spacing: 0.1em;
  text-decoration: none;
  color: #000;
  display: inline-block;
  padding: 10px 30px;
  position: relative;
  border: 3px solid #9e9a9a;
  border-radius: 20px;
  right: 0px;
  background-color: rgb(255, 255, 255);
  border: 1px solid #e3e3e3;
  border-radius: 24px;
  bottom: 7px;
  right: 8px;
  padding-top: 10px;
  position: absolute;
}

.price button:hover {
  opacity: 0.8;
  background-color: #b61c36cc;
  text-decoration: none;
  color: #fff;
  transition: all 0.4s;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);
  transform: scale(1.05);
}
.discount {
  color: #fff;
  font-size: 13px;
  font-weight: 800;
  text-align: center;
  line-height: 40px;
  width: 40px;
  border-radius: 50%;
  position: absolute;
  transition: all 0.3s;
  right: 14px;
  top: 14px;
  background-color: #d7292a;
}


</style>
