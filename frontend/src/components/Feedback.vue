<template>
  <div class="happy-clients">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <div class="section-heading">
            <h2>Phản hồi từ khách hàng</h2>
            <a href="testimonials.html"
              >Đọc thêm <i class="fa fa-angle-right"></i
            ></a>
          </div>
        </div>

        <!-- <div class="col-md-12">
          <div class="owl-clients owl-carousel text-center">
            <div class="service-item">
              <div class="icon">
                <i class="fa fa-user"></i>
              </div>
              <div class="down-content">
                <h4>John Doe</h4>
                <p class="n-m">
                  <em
                    >"Lorem ipsum dolor sit amet, consectetur an adipisicing
                    elit. Itaque, corporis nulla at quia quaerat."</em
                  >
                </p>
              </div>
            </div>
          </div>
        </div> -->

        <div
          class="feedback-item col-md-4"
          v-for="feedback in feedbacks"
          :key="feedback.id"
        >
          <img :src="feedback.image" alt="hihi" />
          <div>
            <h4>{{ feedback.nameCustomer }}</h4>
            <p>
              {{ feedback.content }}
            </p>
            <div style="height: 40px"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Feedback",
  data() {
    return {
      feedbacks: [
        {
          id: 1,
          idCustomer: 1,
          image: require("@/assets/images/feedback_ava/person_1.jpg"),
          nameCustomer: "Agatha Christie",
          content:
            "Tôi rất hài lòng với sản phẩm của cửa hàng. Mong muốn cửa hàng có nhiều đợt giảm giá hơn nữa!!",
          rate: 5,
        },
        {
          id: 2,
          idCustomer: 2,
          image: require("@/assets/images/feedback_ava/person_2.jpg"),
          nameCustomer: "John Doe",
          content:
            "Sản phẩm đẹp, shop phục vụ nhiệt tình, mình rất thích, chắc lại là fant của nhà xinh mất thôi. ",
          rate: 4,
        },
        {
          id: 3,
          idCustomer: 3,
          image: require("@/assets/images/feedback_ava/person_3.jpg"),
          nameCustomer: "Patricia Highsmith",
          content:
            "Hàng đẹp mẫu mã hợp mắt uy tín ủng hộ shop lâu dài hy vọng shop ra nhiều mẫu đẹp hơn nữa .",
          rate: 5,
        },
      ],
    };
  },
};
</script>

<style scoped>
.section-heading {
  text-align: left;
  margin: 60px 0px;
  border-bottom: 1px solid #eee;
}

.section-heading h2 {
  font-size: 28px;
  font-weight: 400;
  color: #1e1e1e;
  margin-bottom: 15px;
}

.section-heading a {
  float: right;
  margin-top: -35px;
  text-transform: uppercase;
  font-size: 13px;
  font-weight: 700;
  color: #f33f3f;
  text-decoration: none;
}

.feedback-item {
  text-align: center;
}

.feedback-item img {
  width: 100%;
  height: auto;
}
</style>
