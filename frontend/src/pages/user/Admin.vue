<template>
  <div class="table-wrapper">
    <div class="table-title">
      <div class="row">
        <div class="col-sm-8">
          <h2>Chỉnh sửa danh sách sản phẩm</h2>
        </div>
        <div class="col-sm-4">
          <!-- Search -->
          <div>
            <input
              type="text"
              placeholder="Enter Name Group"
              name="nameGroup"
              id="nameGroup"
              required
            />
            <button class="fa fa-search" onclick="search()"></button>
          </div>
          <button
            type="button"
            class="btn btn-info add-new"
            data-toggle="modal"
            onclick="openAddModal()"
          >
            <i class="fa fa-plus"></i> Add New
          </button>
        </div>
      </div>
    </div>
    <table class="table table-bordered">
      <thead>
        <tr>
          <th>Tên</th>
          <th>Ảnh</th>
          <th>Giảm giá (%)</th>
          <th>Giá</th>
          <th>Mua nhiều</th>
          <th>Sửa</th>
        </tr>
        <tr v-for="product in listProducts" :key="product.id">
          <td>{{ product.name }}</td>
          <td>{{ product.image }}</td>
          <td>{{ product.sale }}</td>
          <td>{{ product.price }}</td>
          <td>{{ product.isFeature == true ? "Đang hot" : "Không" }}</td>
          <td>
            <button><i class="fas fa-pen"></i></button>
            <button @click="openConfirmDelete(product.id)">
              <i class="fas fa-trash"></i>
            </button>
          </td>
        </tr>
      </thead>
      <tbody></tbody>
    </table>
    <ul class="pagination"></ul>
  </div>

 
  <div>
    <ModalUpdate />
  </div>
  <div v-show="confirm">
    <AlertSuccess />
  </div>
</template>

<script>
import AlertSuccess from "../../components/edit-product/AlertSuccess.vue";
import ModalUpdate from "../../components/edit-product/ModalUpdate.vue";
import { DeleteData, GetData } from "../../utils/callapi.js";

export default {
  name: "Admin",

  components: {
    AlertSuccess,
    ModalUpdate,
  },

  data() {
    return {
      listProducts: [],
      confirm: false,
    };
  },

  methods: {
    async getAllProducts() {
      this.listProducts = await GetData(`products?limit=10&cursor=0&all=1`);
    },

    openConfirmDelete(productId) {
      let index = this.listProducts.findIndex(
        (product) => product.id === productId
      );

      var result = confirm(
        "Want to delete " + this.listProducts[index].name + "?"
      );
      this.confirm = true;
      console.log(this.confirm);
      if (result) {
        this.deleteGroup(this.listProducts[index].id);
      }
    },

    async deleteGroup(productId) {
      let response = await DeleteData(`products/${productId}`);
      if (response.status) {
        alert("Xóa sản phẩm ok :>");
        this.getAllProducts();
      } else {
        alert("Xóa sản phẩm bị lỗi!!!");
      }
    },
  },

  created() {
    this.getAllProducts();
  },
};
</script>
