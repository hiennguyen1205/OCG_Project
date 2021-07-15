<template>
  <div class="products">
    <div class="container">
      <div class="text-content buffer-top-text">
        <h2>Thông tin chi tiết Sản Phẩm</h2>
      </div>
      <div class="row buffer-top">
        <div class="col-md-4 col-xs-12">
          <div>
            <img
              :src="
                product.image ? `http://localhost:3000/${product.image}` : ''
              "
              alt=""
              class="img-fluid wc-image"
            />
          </div>
        </div>

        <div class="col-md-8 col-xs-12">
          <div class="form">
            <h2>{{ product.name }}</h2>
            <br />
            <p class="lead">
              <!-- <small><del> $999.00</del></small><strong class="text-primary">$779.00</strong> -->
              <strong class="text-primary">{{ money }} </strong>
            </p>

            <br />

            <br />

            <div class="row">
              <!-- <div class="col-sm-4">
                <label class="control-label">Loại sản phẩm</label>
                <div class="form-group">
                  <select class="form-control">
                    <option value="0">18 gears</option>
                    <option value="1">21 gears</option>
                    <option value="2">27 gears</option>
                  </select>
                </div>
              </div> -->
              <div class="col-sm-12">
                <label class="control-label">Số lượng</label>

                <div class="row">
                  <div class="col-sm-6">
                    <div class="form-group">
                      <input
                        type="text"
                        class="form-control"
                        placeholder="1"
                        v-model.number="quantityOfProduct"
                      />
                    </div>
                  </div>

                  <div class="col-sm-6">
                    <button
                      class="btn btn-primary btn-block"
                      @click="addToCart(product)"
                    >
                      Add to Cart
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <h2>Mô tả sản phẩm</h2>
        <p class="lead">
          {{ product.description }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductDetail",
  props: ["id"],
  data() {
    return {
      product: {},
      quantityOfProduct: 1,
    };
  },

  async beforeCreate() {
    // console.log(this.id);
    const response = await fetch(
      `http://localhost:3000/api/products/${this.id}`
    );
    this.product = await response.json();
    // console.log(this.product);
  },

  methods: {
    addToCart(product) {
      product["quantity"] = this.quantityOfProduct;
      this.$store.commit("addProductToCart", product);
    },
  },

  computed: {
    money() {
      const formatter = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      });
      return formatter.format(this.product.price);
    },
  },
};
</script>

<style scoped>
.page-heading {
  padding: 210px 0px 90px 0px;
  text-align: center;
  background-position: center center;
  background-size: cover;
  position: relative;
  background-repeat: no-repeat;
  background-image: url("../assets/images/background/heading-1-1920x500.jpg");
}

.page-heading:before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.7);
}

.buffer-top-text {
  margin-top: 100px;
  border-bottom: 1px solid rgb(163, 163, 163);
}
.buffer-top {
  margin-top: 50px;
}
.row {
  margin-bottom: 20px;
}
</style>
