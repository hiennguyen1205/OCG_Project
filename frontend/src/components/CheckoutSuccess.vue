<template>
 <div id="myModal" class="modal parent">
    <!-- Modal content -->
    <div class="modal-content">
      <div class="modal-header">
     
      <h2>Thông báo</h2>
       <span class="close"  @click="close">&times;</span>
    </div>
      {{paymented ? "Bạn đã thanh toán thành công":"Bạn chưa thanh toán thành công!!!"}}
      <div class="button">
        <!-- <button >Quay lại trang sản phẩm</button> -->
        <button @click="close" :class="[paymented ? 'btn btn-success' : 'btn btn-danger']">Đóng</button>
      </div>
    </div>
  </div>

</template>

<script>
export default {
  name: "CheckoutSuccess",
 props:["paymented"],
  data() {
    return {
    };
  },
  methods: {
    close(){
      this.$emit("Close",false)
    }
  },
};
</script>

<style scoped>
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
.parent{
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 220px;
}
.modal-header > h2{
  margin: 0px auto;
}

</style>