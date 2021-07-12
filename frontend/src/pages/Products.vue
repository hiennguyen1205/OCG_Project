<template>
  <!-- <Banner :images="images" /> -->
  <div class="container">
    <div class="row"></div>
  </div>
  <!-- Category -->
  <div class="products-list buffer">
    <div class="container">
      <div class="row">
        <div class="col-md-3">
          <h2 style="border-bottom: 1px solid rgb(189, 189, 189);">
            Danh mục
          </h2>
          <ul>
            <!-- <button @click="getAmrchairs">Armchair</button> -->
            <li>Armchair</li>
            <li>Sofa</li>
            <li>Table</li>
            <li>Tủ tivi</li>
            <li>Kệ trưng bày</li>
            <li>Ghế dài và đôn</li>
            <li>Ghế thư giãn</li>
            <li>Bàn ăn</li>
          </ul>
        </div>

        <!-- Products List -->
        <div class="col-md-9">
          <div class="row">
            <div class="filter">
              <div class="col-md-4">
                <select
                  name=""
                  @change="selectOption($event)"
                  style="margin: 30px 20px 0px 0px"
                >
                  <option value="" id="">--Sắp xếp--</option>
                  <option value="price-desc" id="op2">Giá cao đến thấp</option>
                  <option value="price-asc" id="op3">Giá thấp đến cao</option>
                </select>
              </div>
              <div class="col-md-4">
                <div class="input-group rounded">
                  <input
                    type="search"
                    class="form-control rounded"
                    placeholder="Search"
                    aria-label="Search"
                    aria-describedby="search-addon"
                    v-model="searchProducts"
                  />
                  <!-- <span class="input-group-text border-0" id="search-addon">
                    <i class="fas fa-search"></i>
                  </span> -->
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div
              class=" col-md-4"
              v-for="product in searchProductsByName"
              :key="product.id"
            >
              <div class="product-item">
                <div :class="{ discount: product.sale }">
                  {{ product.sale > 0 ? product.sale + "%" : "" }}
                </div>
                <img
                  :src="`http://localhost:3000/${product.image}`"
                  alt="hihi"
                />
                <div class="decription">
                  <router-link
                    :to="{
                      name: 'ProductDetail',
                      params: { id: product.id },
                    }"
                    ><h4>{{ product.name }}</h4></router-link
                  >
                  <p>
                    {{ product.description }}
                  </p>
                  <div class="price">
                    <div
                      id="style-price"
                      style="text-align:center; font-size:30px"
                    >
                      {{ formatCurrency(product.price) }}
                    </div>
                    <button @click="addOneToCart(product)">
                      <i
                        id="icon-cart"
                        class="fas fa-shopping-cart"
                        style="padding: 5px"
                      ></i
                      >Buy
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Pagination -->
          <!-- <div class="paging">
            <nav aria-label="Page navigation example">
              <ul class="pagination">
                <li class="page-item">
                  <a class="page-link" href="#" aria-label="Previous">
                    <span aria-hidden="true">&laquo;</span>
                    <span class="sr-only">Previous</span>
                  </a>
                </li>

                <li
                  class="page-item"
                  v-for="(page, index) in totalPage"
                  :key="index"
                >
                  <a class="page-link" href="#" @click="paging(index)">{{
                    index + 1
                  }}</a>
                </li>

                <li class="page-item">
                  <a class="page-link" href="#" aria-label="Next">
                    <span aria-hidden="true">&raquo;</span>
                    <span class="sr-only">Next</span>
                  </a>
                </li>
              </ul>
            </nav>
          </div> -->

          <!-- test xem thêm -->
          <div>
            <button @click="paging()">Xem thêm nào!!</button>
          </div>




        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import Banner from "../components/Banner.vue";

export default {
  name: "ProductList",
  components: {
    // Banner,
  },

  data() {
    return {
      products: [],
      pageIndex: 1,
      limit: 6,
      sort: "ID",
      search: "",
      totalProducts: 0,
      cursor: 0,
      totalPage: 0,
      category: "",
      searchProducts: "",
    };
  },

  async created() {
    const response = await fetch(
      `http://localhost:3000/api/products?limit=${this.limit}&cursor=${this.cursor}&categoryId=1`
    );
    let result = await response.json();
    this.products = result;
    // console.log(this.products);
    this.totalPage = result.totalPages;
    this.totalProducts = result.totalProducts;
    // console.log(  this.totalPage );
    // console.log(  this.totalProducts );
  },

  computed: {
    searchProductsByName() {
      return this.products.filter((product) => {
        return product.name.toLowerCase().includes(this.searchProducts);
      });
    },
  },

  methods: {
    formatCurrency(money) {
      return money.toLocaleString("vi", { style: "currency", currency: "VND" });
    },

    addOneToCart(product) {
      product["quantity"] = 1;
      this.$store.commit("addProductToCart", product);
    },


    async paging() {
      this.limit += this.limit;
      const response = await fetch(
        `http://localhost:3000/api/products?limit=${this.limit}&cursor=${this.cursor}&categoryId=1`
      );
      let result = await response.json();
      console.log(result);
      this.products = result;
    },

    selectOption(event) {
      this.$store.state.products = this.products;
      this.$store.commit("selectOption", event.target.value);
    },
    updateSearchValue(event) {
      // console.log(event.target.value);
      this.searchProducts = event.target.value;
      // console.log(this.searchProducts);
    },
  },
};
</script>

<style scoped>
.filter {
  display: flex;
  justify-content: space-between;
  margin-bottom: 50px;
}
.buffer {
  margin: 120px 0px 60px 0px;
}
li {
  margin-left: 5px;
}
.decription a {
  text-decoration: none;
  font-size: 20px;
  font-weight: 700;
  text-align: center;
  color: black;
}
.decription a:hover {
  background-color: rgb(250, 133, 133);
  background-color: #f33f3f;
  border-color: #f33f3f;
  /* color: #fff; */
}

.rounded {
  margin-top: 10px;
}
.input-group-text {
  margin-top: 10px;
}

/* .active{
   text-decoration: underline;
   background-color: white
} */

ul {
  padding: 0px;
  list-style: none;
}
/* pagination */
.paging {
  display: flex;
  justify-content: center;
}
/* Banner */
.slide {
  margin-bottom: 30px;
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
  border: 1px solid rgb(94, 94, 94);
  margin-bottom: 30px;

  /* width: 32%; */
}

.product-item img {
  width: 100%;
  max-height: 155px;
  overflow: hidden;
}

.product-item p {
  padding: 5px 5px 0px 5px;
  text-align: justify;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  /* transition: all 0.2s; */
}
.product-item h4 {
  padding: 5px;
}

.style-price {
  font-size: 20px;
  text-align: center;
}

.product-item {
  transition: transform 0.6s ease;
}

.product-item:hover {
  transform: scale(1.1);
  border: 1px solid rgb(240, 90, 4);
}

.price button {
  font-family: lato, sans-serif;
  font-weight: bold;
  font-size: 1em;
  letter-spacing: 0.1em;
  text-decoration: none;
  color: #000;
  display: inline-block;
  width: 60%;
  padding: 10px 15px;
  margin: 0px auto;
  /* position: relative; */
  border: 3px solid #9e9a9a;
  border-radius: 20px;
  right: 0px;
  background-color: rgb(255, 255, 255);
  border: 1px solid #e3e3e3;
  border-radius: 24px;
  margin-bottom: 5px;
}
.price {
  display: flex;
  flex-direction: column;
}

.price button {
  transition: transform 0.6s linear;
}

.price button:hover {
  opacity: 0.8;
  background-color: #b61c36cc;
  color: #fff;
  transition: all 0.4s;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);
  transform: scale(1.02);
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
