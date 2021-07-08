<template>
  <!-- Page Content -->
  <div class="page-heading about-heading">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <div class="text-content">
            <h4>Giới thiệu</h4>
            <h2>Công ty Nội thất Nhà Xinh</h2>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="best-features ">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <div class="section-heading">
            <h2>Giá trị và sự khác biệt</h2>
          </div>
        </div>

        <div class="col-md-6">
          <div class="right-image">
            <img src="../assets/images/about/about-1-570x350.jpg" alt="" />
          </div>
        </div>

        <div class="col-md-6">
          <div class="">
            <h4>Định hướng chiến lược</h4>
            <p>
              Với mong muốn phát triển thương hiệu Việt bằng nội lực, Nhà Xinh
              đã chú trọng vào thiết kế và sản xuất nội thất trong nước. Danh
              mục sản phẩm của Nhà Xinh thường xuyên được đổi mới và cập nhật,
              liên tục cung cấp cho khách hàng các dòng sản phẩm theo xu hướng
              mới nhất. Do chính người Việt thiết kế và sản xuất, nội thất
              thương hiệu Nhà Xinh luôn phù hợp với cuộc sống Á Đông, đem đến sự
              tiện nghi hoàn hảo trong mọi không gian sống. Hơn 70% sản phẩm của
              Nhà Xinh được thiết kế, sản xuất bởi đội ngũ nhân viên cùng công
              nhân ưu tú với nhà máy có cơ sở vật chất hiện đại bậc nhất tại
              Việt Nam. Sự khác biệt của Nhà Xinh chính là sáng tạo nội thất
              thành phong cách riêng, phù hợp với nhu cầu khách hàng.
            </p>
            <ul class="social-icons">
              <li>
                <a href="#"><i class="fab fa-facebook"></i></a>
              </li>
              <li>
                <a href="#"><i class="fab fa-twitter"></i></a>
              </li>
              <li>
                <a href="#"><i class="fab fa-linkedin"></i></a>
              </li>
              <li>
                <a href="#"><i class="fab fa-behance"></i></a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="team-members">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <div class="section-heading">
            <h2>Chất lượng và dịch vụ</h2>
          </div>
          <p>
            Chất lượng của nguyên vật liệu, phụ kiện và quy trình sản xuất đều
            được kiểm định và giám sát chặt chẽ bởi hệ thống quản lý chất lượng
            ISO 9001. Sản phẩm của Nhà Xinh được thiết kế theo định hướng công
            năng sử dụng, thẩm mỹ và chất lượng. Trong những năm gần đây, thương
            hiệu luôn hướng đến xu hướng thiết kế xanh nhằm đóng góp không chỉ
            một không gian sống tiện nghi mà còn là một môi trường sống trong
            lành cho người sử dụng và cộng đồng. Với nhiều cống hiến như vậy,
            Nhà Xinh vinh dự nhiều năm liền được trao tặng các danh hiệu “Hàng
            Việt Nam chất lượng cao”, “Trusted brand” và “Top 100 nhà cung cấp
            hàng đầu”. Bên cạnh đó, Nhà Xinh tự hào sở hữu đội ngũ tư vấn thiết
            kế và kỹ sư chuyên nghiệp, có kiến thức sâu rộng trong lĩnh vực đồ
            gỗ nội thất. Tập thể nhân viên tại Nhà Xinh cam kết nỗ lực tư vấn và
            trợ giúp khách hàng lựa chọn sản phẩm ưng ý nhất. Dịch vụ tư vấn
            thiết kế của Nhà Xinh sẽ giúp khách hàng kiến tạo một không gian
            sống như ý thông qua sự phối hợp điêu luyện giữa các sản phẩm nội
            thất và đồ trang trí.
          </p>
          <h5>Nhà máy</h5>
          <p>
            Phần lớn sản phẩm của Nhà Xinh được sản xuất tại nhà máy của Tổng
            công ty AA (AA Đức Hòa, Long An và AA Tây Ninh) có diện tích
            80.000m2, với trang thiết bị hiện đại nhất và theo tiêu chuẩn ISO
            9001, chuyên sản xuất các sản phẩm nội thất chất lượng cao cho các
            khách sạn, nhà hàng 5 sao, dự án nội thất cao cấp và xuất khẩu.
          </p>
        </div>
      </div>
    </div>
  </div>

  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <div class="signature">
          <h6>Nội thất Nhà Xinh</h6>
          <h6>www.nhaxinh.com</h6>
          <h6>Hotline: 1800 7200</h6>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AboutUs",
  components: {},
};
</script>

<style scoped>
.signature{
  float: right;
  margin-bottom: 10px;
}
.page-heading {
  padding: 210px 0px 130px 0px;
  text-align: center;
  background-position: center center;
  background-size: cover;
  position: relative;
  background-repeat: no-repeat;
  background-image: url("../assets/images/background/heading-1-1920x500.jpg");
}

.page-heading:before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.7);
}

.page-heading .text-content {
  position: relative;
}

.page-heading .text-content h4 {
  color: #f33f3f;
  font-size: 22px;
  text-transform: uppercase;
  font-weight: 700;
  margin-bottom: 15px;
}
.page-heading .text-content h2 {
  color: #fff;
  font-size: 62px;
  text-transform: uppercase;
  letter-spacing: 5px;
}

.section-heading h2 {
  font-size: 28px;
  font-weight: 400;
  color: #1e1e1e;
  margin-bottom: 55px;
  padding: 10px;
  border-bottom: 1px solid rgb(206, 206, 206);
}

.best-features {
  margin-top: 60px;
}
.best-features img {
  width: 100%;
  overflow: hidden;
}
.best-features h4 {
  font-size: 17px;
  color: #1a6692;
  margin-bottom: 20px;
}
.about-features p {
  border-bottom: 1px solid #eee;
  padding-bottom: 42px;
}

ul {
  display: flex;
  padding: 0;
  margin: 0;
  list-style: none;
}

ul.social-icons li {
  display: inline-block;
  margin-right: 5px;
}
ul.social-icons li a:hover {
  background-color: #f33f3f;
  color: #fff;
}

ul.social-icons li a {
  width: 50px;
  height: 45px;
  display: inline-block;
  line-height: 50px;
  background-color: #eee;
  color: #121212;
  font-size: 18px;
  text-align: center;
  transition: all 0.3s;
}
.fab {
  font-size: 30px;
}

.section-heading h2 {
  font-size: 28px;
  font-weight: 400;
  color: #1e1e1e;
  margin: 60px 0px;
  padding: 0px;
  padding-bottom: 20px;
}
</style>
