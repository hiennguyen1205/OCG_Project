<template>
  <!--Main layout-->
  <main class="mt-5 pt-4 position">
    <div class="container wow fadeIn">

      <!-- Heading -->
      <h2 class="my-5 h2 text-center">Checkout form</h2>
      <!--Grid row-->
      <div class="row">
        <!--Grid column-->
        <div class="col-md-8 mb-4">
          <!--Card-->
          <div class="card">
            <!--Card content-->
            <form class="card-body">
              <!--Username-->
              <div class="md-form mb-5">
                <label for="name" class="">Họ và Tên: </label>
                <input type="text" id="name" class="form-control" placeholder="FullName" >
              </div>

              <!--email-->
              <div class="md-form mb-5">
                  <label for="email" class="">Email (optional)</label>
                <input type="text" id="email" class="form-control" placeholder="youremail@example.com">
              </div>

              <!--address-->
              <div class="md-form mb-5">
                   <label for="address" class="">Địa chỉ</label>
                <input type="text" id="address" class="form-control" placeholder="1234 Main St">
              </div>

              <!--Grid row-->
              <div class="row">
                <!--Grid column-->
                <div class="col-lg-4 col-md-12 mb-4">
                  <label for="country">Country</label>
                  <select class="custom-select d-block w-100" id="country" required>
                    <option value="">Lựa chọn thành phố</option>
                    <option>Hà Nội</option>
                     <option>Hải Phòng</option>
                      <option>Đà Nẵng</option>
                       <option>Tp Hồ Chí Minh</option>
                  </select>
                  <div class="invalid-feedback">
                    Please select a valid country.
                  </div>
                </div>
                <!--Grid column-->
                <!--Grid column-->
                <div class="col-lg-4 col-md-6 mb-4">
                  <label for="state">Quận/Huyện</label>
                  <input type="text"  class="custom-select d-block w-100">
                  <div class="invalid-feedback">
                    Please provide a valid state.
                  </div>
                </div>
              </div>
              <!--Grid row-->

              <hr>

              <div class="custom-control custom-checkbox">
                <input type="checkbox" class="custom-control-input" id="same-address">
                <label class="custom-control-label" for="same-address">Giao Hàng đến đại chỉ trên</label>
              </div>
              <div class="custom-control custom-checkbox">
                <input type="checkbox" class="custom-control-input" id="save-info">
                <label class="custom-control-label" for="save-info">Lưu thông tin cho lần mua tiếp theo</label>
              </div>

              <hr>

              <div class="d-block my-3">
                <div class="custom-control custom-radio">
                  <input id="credit" name="paymentMethod" type="radio" class="custom-control-input" checked required>
                  <label class="custom-control-label" for="credit">Credit card</label>
                </div>
                <div class="custom-control custom-radio">
                  <input id="debit" name="paymentMethod" type="radio" class="custom-control-input" required>
                  <label class="custom-control-label" for="debit">Debit card</label>
                </div>
                <div class="custom-control custom-radio">
                  <input id="paypal" name="paymentMethod" type="radio" class="custom-control-input" required>
                  <label class="custom-control-label" for="paypal">Paypal</label>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6 mb-3">
                  <label for="cc-name">Name on card</label>
                  <input type="text" class="form-control" id="cc-name" placeholder="" required>
                  <small class="text-muted">Full name as displayed on card</small>
                  <div class="invalid-feedback">
                    Name on card is required
                  </div>
                </div>
                <div class="col-md-6 mb-3">
                  <label for="cc-number">Credit card number</label>
                  <input type="text" class="form-control" id="cc-number" placeholder="" required>
                  <div class="invalid-feedback">
                    Credit card number is required
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="col-md-3 mb-3">
                  <label for="cc-expiration">Expiration</label>
                  <input type="text" class="form-control" id="cc-expiration" placeholder="" required>
                  <div class="invalid-feedback">
                    Expiration date required
                  </div>
                </div>
                <div class="col-md-3 mb-3">
                  <label for="cc-expiration">CVV</label>
                  <input type="text" class="form-control" id="cc-cvv" placeholder="" required>
                  <div class="invalid-feedback">
                    Security code required
                  </div>
                </div>
              </div>
              <hr class="mb-4">
              <button class="btn btn-primary btn-lg btn-block" type="submit">Continue to checkout</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </main>


  
</template>

<script>
export default {
    name: "CheckoutInfo",
}
</script>

<style scoped>
.position{
    display: flex;
    justify-content: center;
}
</style>