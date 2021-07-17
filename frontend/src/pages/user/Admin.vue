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
          <th>Sửa</th>
        </tr>
        <tr v-for="product in listProducts" :key="product.id">
        
          <td>{{ product.name }}</td>
          <td>{{ product.image }}</td>
          <td>{{ product.sale }}</td>
          <td>{{ product.price }}</td>
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

  <!-- Modal HTML -->
  <div id="myModal" class="modal fade" tabindex="-1">
    <div class="modal-dialog">
      <div class="modal-content">
        <!-- header -->
        <div class="modal-header">
          <h5 class="modal-title">Function Group</h5>
          <button type="button" class="close" data-dismiss="modal">
            &times;
          </button>
        </div>

        <!-- body -->
        <div class="modal-body">
          <div class="modal-container">
            <input type="hidden" id="id" name="id" />

            <label for="name"><b>Name</b></label>
            <input
              type="text"
              placeholder="Enter Name"
              name="name"
              id="name"
              required
            />
            <br />
            <label for="name"><b>Total Member</b></label>
            <input
              type="text"
              placeholder="Enter Total Member"
              name="totalMember"
              id="totalMember"
              required
            />
            <br />
            <!-- <label for="name"><b>Creator</b></label>
                        <input type="text" placeholder="Enter Creator Name" name="creator" id="creator" required> -->
          </div>
        </div>

        <!-- footer -->
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">
            Cancel
          </button>
          <button type="button" class="btn btn-primary" onclick="save()">
            Save
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- success alert -->

  <div class="modal" tabindex="-1" v-if="confirm">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Modal title</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <p>Modal body text goes here.</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
        <button type="button" class="btn btn-primary">Save changes</button>
      </div>
    </div>
  </div>
</div>

</template>

<script>
import { DeleteData, GetData } from "../../utils/callapi.js";
export default {
  name: "Admin",

  data() {
    return {
      listProducts: [],
      confirm: false,
    };
  },

  methods: {
    async getAllProducts() {
      this.listProducts = await GetData(`products?limit=10&cursor=0&all=1`);
      console.log(this.listProducts);
    },

    openConfirmDelete(productId) {
      let index = this.listProducts.findIndex(
        (product) => product.id === productId
      );
      var result = confirm("Want to delete " + this.listProducts[index].name + "?");
      if (result) {
        this.deleteGroup(this.listProducts[index].id);
      }
    },

    async deleteGroup(productId){
        let response = await DeleteData(`products/${productId}`);
        if(response.status){
           alert("Xóa sản phẩm ok :>")
        }else{
            alert("Xóa sản phẩm bị lỗi!!!")
        }
        console.log(response.status);
    },

  },

  created() {
    this.getAllProducts();
  },
};
</script>
