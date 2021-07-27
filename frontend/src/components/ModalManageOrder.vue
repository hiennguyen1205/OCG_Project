<template>
  <div class="parent modal">
    <div class="flex-w mb-4 d-flex justify-content-end">
      <div></div>
    </div>
    <div class="table-btn">
<table class="table table-striped">
      <thead>
        <tr style="font-size: 18px">
          <th scope="col" style="width: 5%;" class="t-center col-1" />
          <th scope="col" class="t-center col-2">Ảnh</th>
          <th scope="col" class="t-center col-2">Tên sản Phẩm</th>
          <th scope="col" class="t-center col-1">Giá</th>
          <th scope="col" class="t-center col-3">Số lượng</th>
          <th scope="col" class="t-center col-2">Tổng tiền</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(order, index) in listOrderItems" :key="index">
          <th scope="row" class="t-center parentCenter">
            <div class="center">{{ index + 1 }}</div>
          </th>
          <td class="t-center parentCenter">
            <div class="center">
              <img :src="`http://localhost:3000/${order.image}`" alt="hihi" />
            </div>
          </td>
          <td class="t-center parentCenter">
            <div class="center">{{ order.name }}</div>
          </td>
          <td class="t-center parentCenter">
            <div class="center">{{ formatCurrency(order.price) }}</div>
          </td>
          <td class="t-center parentCenter">
            <div class="center">
              {{ order.quantity }}
            </div>
          </td>
          <td class="t-center parentCenter">
            <div class="center">
              {{ formatCurrency(Number(order.quantity) * Number(order.price)) }}
            </div>
          </td>
        </tr>

      </tbody>
      
    
    </table>

<button @click="close" class="btn btn-primary">Đóng</button>
    </div>
    
  </div>
</template>

<script>
import { formatCurrency } from '@/utils/currency.js';

export default {
  name: "ModalManageOrder",
  props: ["listOrderItems"],
  data() {
    return {
        total: 0,
    };
  },
  computed: {
  },
  created() {
    console.log(this.listOrderItems[0].quantity*this.listOrderItems[0].price );
  },
  methods: {
     formatCurrency,
     close(){
         this.$emit("close");
     }
  },
};
</script>

<style scoped>
.table-btn{
    background: white;
}
.parent {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 220px;
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
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);
}
.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
}
.table {
  background-color: white;
}
img {
    max-width:150px;
}
tbody{
    font-size: 18px;
}

</style>
