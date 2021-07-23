<template>
  <!--Main layout-->
  <main class="mt-5 pt-4 ">
    <div class="container wow fadeIn">
      <!-- Heading -->
      <div class="row">
        <h2 class="my-5 h2 text-center">Checkout form</h2>
      </div>
      <div class="changeInfo">
        <p>
          Nếu thông tin sai, click vào đây để thay đổi thông tin. Lưu ý, thông
          tin bạn nhập phải chính xác để phục vụ cho mục đích nhận hàng.
        </p>
        <span
          ><button class="btn btn-outline-primary" @click="changeInfoUser()">
            Thay đổi thông tin cá nhân
          </button></span
        >
      </div>

      <!--Grid row-->
      <div class="row position">
        <!--Grid column-->
        <div class="col-md-8 mb-4">
          <!--Card-->
          <div class="card">
            <!--Card content-->
            <form class="card-body">
              <!--Username-->
              <div class="md-form mb-5">
                <label for="name" class="">Họ và Tên: </label>
                <input
                  type="text"
                  id="name"
                  class="form-control"
                  placeholder="VD: Nguyễn Văn A"
                  v-model="user.name"
                  disabled
                />
              </div>

              <!--email-->
              <div class="md-form mb-5">
                <label for="email" class="">Email: </label>
                <input
                  type="text"
                  id="email"
                  class="form-control"
                  placeholder="VD: abcxyz@gmail.com"
                  v-model="user.email"
                  disabled
                />
              </div>

              <!--address-->
              <div class="md-form mb-5">
                <label for="address" class="">Địa chỉ: </label>
                <input
                  type="text"
                  id="address"
                  class="form-control"
                  placeholder="VD: 123 đường ABC quận XYZ thành phố MNP"
                  v-model="user.address"
                  disabled
                />
              </div>

              <!--Phone Number-->
              <div class="md-form mb-5">
                <label for="phone-number" class="">Điện thoại: </label>
                <input
                  type="text"
                  id="phone-number"
                  class="form-control"
                  placeholder="VD: 0987654321"
                  v-model="user.phone_number"
                  disabled
                />
              </div>
              <hr />

              <div class="d-block my-3">
                <div class="custom-control custom-radio">
                  <input
                    id="credit"
                    name="paymentMethod"
                    type="radio"
                    class="custom-control-input"
                    checked
                    required
                    @change="onChange(false)"
                  />
                  <label class="custom-control-label" for="credit"
                    >Credit card</label
                  >
                </div>
                <div class="custom-control custom-radio">
                  <input
                    id="debit"
                    name="paymentMethod"
                    type="radio"
                    class="custom-control-input"
                    required
                    @change="onChange(true)"
                  />
                  <label class="custom-control-label" for="debit">COD</label>
                </div>
              </div>

              <Stripe v-if="!COD" ref="gateway" />
              <hr class="mb-4" />
              <button
                class="btn btn-primary btn-lg btn-block"
                type="submit"
                @click.prevent="payment()"
              >
                Continue to checkout
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <CheckoutSuccess v-if="isShow" @close="Close(check)" :paymented="isPaied" />
  </main>
</template>

<script>
import Stripe from "@/components/Stripe";
import CheckoutSuccess from "@/components/CheckoutSuccess";
import { mapState, mapMutations } from "vuex";
import { PostData, AuthPutData } from "../utils/callapi";

export default {
  name: "CheckoutInfo",
  components: {
    Stripe,
    CheckoutSuccess,
  },
  data() {
    return {
      COD: false,
      gateway: "stripe",
      isShow: false,
      isPaied: false,
      isDisplayPayment: false,
    };
  },
  created() {
    console.log(this.user);
  },
  computed: {
    ...mapState("users", ["user"]),
    ...mapState("carts", ["order"]),
  },
  methods: {
    ...mapMutations("carts", ["emptyListProducts"]),

    onChange(bool) {
      this.COD = bool;
    },
    payment() {
      if (this.COD) {
        this.$router.push({ name: "Home" });
      } else {
        let order = this.order;
        this.$refs.gateway.createPaymentMethod().then((res) => {
          if (!res.paymentMethod) return;
          fetch("http://localhost:3000/api/payment", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
              payment_method_id: res.paymentMethod.id,
            }),
          }).then(async (res) => {
            console.log(res.status);
            if (res.status == 200) {
              this.isPaied = true;
              this.isShow = true;
              await PostData("email", order);
              await AuthPutData("orders", order);
            } else {
              this.isPaied = false;
              this.isShow = true;
            }
          });
        });
        // this.$router.push({ name: 'CheckoutInfomation' });
      }
    },
    Close(check) {
      if (this.isPaied) {
        this.emptyListProducts();
        this.isShow = check;
        this.$router.push({ name: "Home" });
      } else {
        this.isShow = check;
      }
    },
    changeInfoUser() {
      this.$router.push({ name: "UserInfor" });
    },
  },
};
</script>

<style scoped>
.position {
  display: flex;
  justify-content: center;
}
form > button {
  background-color: black;
  justify-content: center;
}
.changeInfo {
  margin-bottom: 20px;
  text-align: center;
  color: red;
  display: flex;
  justify-content: space-around;
}
.changeInfo p {
  display: inline;
  max-width: 550px;
}
</style>
