<template>
  <div id="myModal" class="modal">
    <!-- Modal content -->
    <div class="modal-content">
      <label for="name"><b>Tên sản phẩm</b></label>
      <input
        type="text"
        placeholder="Enter Name"
        name="name"
        id="name"
        v-model="productUpdate.name"
        required
      />
      <br />
      <label for="name"><b>Giá</b></label>
      <input
        type="number"
        placeholder="Enter Price"
        name="price"
        id="price"
         v-model="productUpdate.price"
        required
      />
      <br />
       <label for="name"><b>Giảm giá (%)</b></label>
      <input
        type="number"
        placeholder="Enter Sale"
        name="sale"
        id="sale"
         v-model="productUpdate.sale"
        required
      />
      <br />
      <div class="button">
        <button @click="yesConfirm">YES</button>
        <button @click="noConfirm">NO</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ModalUpdate",
  props: ["productModal"],
  data(){
    return{
      productUpdate: {},
    };
  },
  mounted() {
    this.productUpdate = this.productModal;
  },
  methods: {
    noConfirm: function() {
      this.$emit("confirmUpdateModal", null);
      
    },
    yesConfirm: function() {
      this.productUpdate.sale = Number(this.productUpdate.sale)
      this.productUpdate.price = Number(this.productUpdate.price)
      this.$emit("confirmUpdateModal", this.productUpdate);
    },
  },
};
</script>

<style scoped>
.button {
  display: flex;
  justify-content: space-around;
}
[type=button]:not(:disabled), [type=reset]:not(:disabled), [type=submit]:not(:disabled), button:not(:disabled) {
  border-radius: 10px;
  cursor: pointer;
  height: 40px;
  width: 400px;
}
.modal {
  text-align: center;
  font-size: 25px;
  display: block;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
}

/* The Close Button */
.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
</style>

