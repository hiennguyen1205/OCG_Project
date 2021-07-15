<template>
  <div>
    <!-- Header -->
    <header class="container">
      <h1 class="text-header">Giỏ Hàng</h1>
      <h5 class="count">{{ totalProducts }} sản phẩm hiện có trong giỏ</h5>
    </header>
    <!-- End Header -->

    <!-- List Product -->
    <section class="container">
      <div v-if="productsInChild.length > 0">
        <ul class="products">
          <li class="row" v-for="product in productsInChild" :key="product.id">
            <!-- {{product.id}},{{index}},{{key}} -->
            <div class="col left ">
              <div class="thumbnail">
                <a href="#">
                  <img
                    :src="`http://localhost:3000/${product.image}`"
                    :alt="product.name"
                  />
                </a>
              </div>
              <div class="detail">
                <div class="name">
                  <router-link to="/detail_product"
                    ><p>{{ product.name }}</p></router-link
                  >
                </div>
                <div class="description">{{ product.description }}</div>
                <div class="price">{{ formatCurrency(product.price) }}</div>
              </div>
            </div>
            <div class="col right">
              <div class="quantity">
                <input
                  type="number"
                  class="quantity"
                  v-model="product.quantity"
                  @input="changeQuantity(product.id, $event)"
                />
              </div>

              <div class="remove">
                <i
                  class="fas fa-trash"
                  id="delete-product"
                  @click="removeItem(product.id)"
                ></i>
              </div>
            </div>
          </li>
        </ul>

        <!-- <div class="dropdown">
       
            <div class="row">
              <h4 class="col-md-4">Phương thức thanh toán</h4>
              <div class="col-md-4"></div>
              <button
                class="btn btn-secondary dropdown-toggle col-md-4"
                type="button"
                id="dropdownMenuButton1"
                data-bs-toggle="dropdown"
                aria-expanded="false"
              >
                Dropdown button
              </button>
              <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
                <li><a class="dropdown-item" href="#">Action</a></li>
                <li><a class="dropdown-item" href="#">Another action</a></li>
                <li>
                  <a class="dropdown-item" href="#">Something else here</a>
                </li>
              </ul>
            </div>
        
        </div> -->
      </div>

      <div v-else class="empty-product">
        <h3>Không có sản phẩm trong giỏ hàng.</h3>
        <router-link to="/products"
          ><button class="btn">Tới trang sản phẩm</button></router-link
        >
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "CartListProducts",
  data() {
    return {
      // productsInChild: [],
    };
  },
  created() {
    // this.productsInChild = [...this.$store.state.order.products] ;
  },
  methods: {
    changeQuantity(productId, event) {
      this.$store.commit("changeQuantity", {
        productId: productId,
        number: event.target.value,
      });
    },
    removeItem(productId) {
      this.$store.commit("removeItem", productId);
    },
    undoProduct() {
      this.$store.commit("undoProduct");
    },
    formatCurrency(money) {
      return money.toLocaleString("vi", { style: "currency", currency: "VND" });
    },
  },

  computed: {
    productsInChild() {
      if (this.$store.state.order.products === null) {
        return [];
      } else {
        return this.$store.state.order.products;
      }
    },
    totalProducts() {
      if (this.$store.state.order.products === null || this.$store.state.order.products === undefined) {
        return 0;
      } else {
        return this.$store.state.order.products.reduce(
          (sumProducts, product) => (sumProducts += parseInt(product.quantity)),
          0
        );
      }
    },
  },
};
</script>

<style scope>
.text-header {
  margin-top: 70px;
}

img {
  max-width: 100%;
  vertical-align: middle;
  border-radius: 4px;
}
a {
  text-decoration: none;
  color: #333333;
}

a:hover {
  color: #f58551;
}
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  margin: 0;
}

header .container {
  margin-bottom: 1.5rem;
  border-bottom: 1px solid #ddd;
}

/*  */

header .count {
  float: left;
  margin-top: 10px;
  color: #333333;
  height: 20px;
  line-height: 20px;
}

/* --- PRODUCT LIST --- */
/* .products {
  border-top: 1px solid #ddd;
} */

.products > li {
  padding: 1rem 0;
  /* border-bottom: 1px solid #ddd; */
}

.row {
  position: relative;
  overflow: auto;
  width: 100%;
}

.col,
.quantity {
  float: left;
}

.col.left {
  width: 70%;
  display: flex;
  align-items: center;
}

.col.right {
  width: 30%;
  position: absolute;
  right: 0;
  top: calc(50% - 30px);
  display: flex;
  justify-content: center;
  align-items: center;
}

.detail {
  padding: 0 0.5rem;
  line-height: 2.2rem;
  margin-left: 30px;
}

.detail .name {
  font-size: 1.2rem;
}

.detail .description {
  color: #7d7d7d;
  font-size: 1rem;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  width: 250px;
}

.detail .price {
  font-size: 1.5rem;
}

.quantity {
  width: 50%;
  text-align: center;
}

.quantity > input {
  display: inline-block;
  width: 60px;
  height: 60px;
  position: relative;
  left: calc(50% - 30px);
  background: #fff;
  border: 2px solid #ddd;
  color: #7f7f7f;
  text-align: center;
  font: 600 1.5rem Helvetica, Arial, sans-serif;
}

.quantity > input:hover,
.quantity > input:focus {
  border-color: #f58551;
}

.close:hover {
  fill: #f58551;
}
</style>
