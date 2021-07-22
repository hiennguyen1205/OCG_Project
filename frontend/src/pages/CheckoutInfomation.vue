<template>
  <!--Main layout-->
  <main class="mt-5 pt-4 ">
    <div class="container wow fadeIn">
      <!-- Heading -->
      <div class="row">
        <h2 class="my-5 h2 text-center">Checkout form</h2>
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
                  placeholder="FullName"
                />
              </div>

              <!--email-->
              <div class="md-form mb-5">
                <label for="email" class="">Email (optional)</label>
                <input
                  type="text"
                  id="email"
                  class="form-control"
                  placeholder="youremail@example.com"
                />
              </div>

              <!--address-->
              <div class="md-form mb-5">
                <label for="address" class="">Địa chỉ</label>
                <input
                  type="text"
                  id="address"
                  class="form-control"
                  placeholder="1234 Main St"
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
  </main>
</template>

<script>
import Stripe from '@/components/Stripe';
export default {
  name: 'CheckoutInfo',
  components: {
    Stripe,
  },
  data() {
    return {
      COD: false,
      gateway: 'stripe',
    };
  },
  methods: {
    onChange(bool) {
      this.COD = bool;
    },
    payment() {
      if (this.COD) {
        this.$router.push({ name: 'Home' });
      } else {
        this.$refs.gateway.createPaymentMethod().then(async (res) => {
          console.log(res.paymentMethod.id);
          const response = await fetch('http://localhost:3000/api/payment', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
              payment_method_id: res.paymentMethod.id,
            }),
          });
          console.log(response);
        });
        // this.$router.push({ name: 'CheckoutInfomation' });
      }
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
</style>
