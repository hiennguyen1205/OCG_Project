<template>
  <div>
    <!-- Summary -->
    <section class="container">
      <div class="pay">
        <div class="promotion">
          <label for="promo-code">Nhập mã giảm giá</label>
          <input
            type="text"
            id="promo-code"
            :value="showPromoCode"
            @input="fowardDiscountCode"
          />
          <button
            type="button"
            @click="getDiscount"
            class="button-discount btn"
          ></button>

          <!-- <p>{{promoCode}}</p> -->
        </div>

        <div class="summary">
          <ul>
            <li>
              Subtotal <span> {{ formatCurrency(showSubTotal) }}</span>
            </li>
            <li>
              Discount <span>{{ formatCurrency(showDiscount) }}</span>
            </li>
            <li>
              Tax <span>{{ formatCurrency(showTax) }}</span>
            </li>
            <li class="total">
              Total
              <span>{{
                formatCurrency(showSubTotal + showTax - showDiscount)
              }}</span>
            </li>
          </ul>
        </div>
      </div>

      <div class="checkout">
        <router-link to="/checkout/infomation"><button type="button" class="btn" @click="checkout()">Check Out</button></router-link>
      </div>
    </section>
    <!-- End Summary -->
  </div>
</template>

<script>
import { mapActions, mapGetters, mapMutations, mapState } from "vuex";
import { formatCurrency } from "@/utils/currency.js";

export default {
  name: "CartPayment",

  data() {
    return {
      user: {},
      products: {},
      order: {},
    };
  },

  computed: {
    ...mapGetters('carts',["calcSubTotal", "calcTax"]),
    ...mapState('carts',["discount", "promoCode", "order"]),
    showSubTotal() {
      return this.calcSubTotal;
    },

    showTax() {
      return this.calcTax;
    },

    showDiscount() {
      return this.discount;
    },
    showPromoCode() {
      console.log(this.promoCode);
      return this.promoCode;
    },
  },

  methods: {
    ...mapMutations('carts',["calcDiscount", "changeDiscountCode"]),
    ...mapActions('carts',["submitOrder"]),

    formatCurrency,
    getDiscount() {
      this.calcDiscount();
    },
    fowardDiscountCode(event) {
      this.changeDiscountCode(event.target.value);
    },
    checkout() {
      this.submitOrder(this.order);
    },
  },
};
</script>

<style scope>
.pay {
  display: flex;
}

ul {
  padding: 0;
  margin: 0;
  list-style-type: none;
}

.btn:hover::after {
  right: -5px;
}
.btn:hover {
  background-color: #f58551;
  border-color: #f58551;
}
.btn::after {
  position: relative;
  right: 0;
  /* content: " \276f"; */
  transition: all 0.1s linear;
}
.btn {
  background-color: #16cc9b;
  border: 2px solid #16cc9b;
  color: #ffffff;
  transition: all 0.25s linear;
  cursor: pointer;
  margin-bottom: 10px;
}
/* --- SUMMARY --- */
.promotion,
.summary,
.checkout {
  float: left;
  width: 100%;
  margin-top: 1.5rem;
}

.promotion > label {
  float: left;
  width: 100%;
  margin-bottom: 1rem;
}

.promotion > input {
  float: left;
  width: 80%;
  font-size: 1rem;
  padding: 0.5rem 0 0.5rem 1.8rem;
  border: 2px solid #16cc9b;
  border-radius: 2rem 0 0 2rem;
  height: 2.8rem;
}

.promotion:hover > input {
  border-color: #f58551;
  outline: none;
}

.promotion > .btn {
  float: left;
  width: 20%;
  height: 2.8rem;
  border-radius: 0 2rem 2rem 0;
}

.promotion:hover > .btn {
  border-color: #f58551;
  background-color: #f58551;
}

.promotion > .btn::after {
  content: "\276f";
  font-size: 1rem;
}

.summary {
  font-size: 1.2rem;
  text-align: right;
}

.summary ul li {
  padding: 0.5rem 0;
}

.summary ul li span {
  display: inline-block;
  width: 30%;
}

.summary ul li.total {
  font-weight: bold;
}

.checkout {
  text-align: right;
}

.checkout > .btn {
  font-size: 1.2rem;
  padding: 0.8rem 2.8rem;
  border-radius: 1.5rem;
}

.empty-product {
  text-align: center;
}

.empty-product > .btn {
  font-size: 1.3rem;
  padding: 10px 30px;
  border-radius: 5px;
}
</style>
