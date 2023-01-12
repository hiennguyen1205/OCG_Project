<template>
  <div>
    <h1>Chi tiết danh sách đơn hàng</h1>

    <table class="table table-bordered">
      <thead>
        <tr>
          <th>STT</th>
          <th>Mã Đơn</th>
          <th>Tên người mua</th>
          <th>Thanh Toán</th>
          <th>Vận Chuyển</th>
          <th>Chi Tiết</th>
        </tr>
        <tr v-for="(order, index) in listOrders" :key="index">
          <td>{{ index + 1 }}</td>
          <td>MDH{{ order.order_id }}</td>
          <td>{{ order.username }}</td>
          <td>
            {{ order.is_paied === 1 ? "Đã thanh toán" : "Chưa thanh toán" }}
          </td>
          <td>{{ checkStatus(order.status) }}</td>
          <td>
            <button
              class="btn btn-primary"
              @click="GetOrderItems(order.order_id)"
            >
              <i class="fas fa-info"></i>
            </button>

            <button class="btn btn-danger" @click="DeleteOrder(order.order_id)">
              <i class="fas fa-trash"></i>
            </button>
          </td>
        </tr>
      </thead>
    </table>
  </div>

  <ModalManageOrder v-if="isShow" :listOrderItems="listOrderItems" @click="close()"/>
</template>

<script>
import { AuthGetData, AuthPutData } from "../../utils/callapi";
import ModalManageOrder from "../../components/ModalManageOrder.vue";
export default {
  name: "AdminOrder",
  components: { ModalManageOrder },
  data() {
    return {
      listOrders: [],
      isShow: false,
      listOrderItems: [],
    };
  },

  created() {
    this.getAllOrder();
  },
  methods: {
    async getAllOrder() {
      let orderInfo = await AuthGetData(`orders/admin`);
      console.log(orderInfo);
      this.listOrders = orderInfo;
    },
    checkStatus(status) {
      if (status === "preparing") return "Đang chuẩn bị";
      else if (status === "shipping") return "Đang giao hàng";
      else return "Đã giao";
    },
    async GetOrderItems(orderId) {
      let response = await AuthGetData(`orders/details/${orderId}`);
      this.listOrderItems = await response;
      this.isShow = await true;
      //   console.log(response);
    },
    async DeleteOrder(orderId) {
      this.listOrders = this.listOrders.filter(
        (order) => order.order_id !== orderId
      );
      let response = await AuthPutData(`orders/${orderId}`, {});
      console.log(response);
    },
    close(){
        this.isShow = false;
    }
  },
};
</script>

<style scoped></style>
