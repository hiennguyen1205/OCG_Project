<template>
  <!-- Site footer -->
  <footer class="site-footer">
    <div class="container">
      <div class="row">
        <div class="col-sm-12 col-md-6">
          <h6>About</h6>
          <p class="text-justify">
            Đóng góp cho cộng đồng nằm trong sứ mệnh của chúng tôi. Để trở thành
            thương hiệu tiên phong và văn minh, chúng tôi đã thực hiện nhiều các
            hoạt động công đồng như hỗ trợ trẻ em nghèo, tham gia các hoạt động
            vì môi trường như thiết kế xanh, đi bộ từ thiện…Những bước chân,
            hành động không ngừng nghỉ để góp phần cho cuộc sống tốt hơn.
          </p>
        </div>

        <div class="col-xs-6 col-md-3 category-footer">
          <h6>Categories</h6>
          <ul class="footer-links">
            <li><a href="#">Sofa</a></li>
            <li><a href="#">Armchair</a></li>
            <li><a href="#">Tủ Tivi</a></li>
            <li><a href="#">Bàn ăn</a></li>
          </ul>
        </div>

        <div class="col-xs-6 col-md-3 quick-link-footer">
          <h6>CẢM HỨNG #NHAXINH</h6>
          <ul class="footer-links">
            <li><a href="#">Sản phẩm mới</a></li>
            <li><a href="#">Ý tưởng và cảm hứng</a></li>
            <li><a href="#">Phong cách nội thất</a></li>
            <li><a href="#">Câu chuyện thiết kế</a></li>
            <li><a href="#">Tạp chí Spaace</a></li>
          </ul>
        </div>
      </div>
      <hr />
    </div>
    <div class="container">
      <div class="row">
        <div class="col-md-8 col-sm-6 col-xs-12">
          <p class="copyright-text">
            Bản quyền thương hiệu được đăng ký tại Cục sở hữu trí tuệ.
          </p>
        </div>

        <div class="col-md-4 col-sm-6 col-xs-12">
          <ul class="social-icons">
            <li>
              <a class="facebook" href="#"><i class="fab fa-facebook"></i></a>
            </li>
            <li>
              <a class="twitter" href="#"><i class="fab fa-twitter"></i></a>
            </li>
            <li>
              <a class="pinteres" href="#"><i class="fab fa-pinterest"></i></a>
            </li>
            <li>
              <a class="linkedin" href="#"><i class="fab fa-linkedin"></i></a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "Footer",
};
</script>

<style scoped>
.site-footer {
  background-color: #26272b;
  /* background-color: #365ae9; */
  padding: 45px 0 20px;
  font-size: 15px;
  line-height: 24px;
  color: #737373;
}
.site-footer hr {
  border-top-color: #bbb;
  opacity: 0.5;
}
.site-footer hr.small {
  margin: 20px 0;
}
.site-footer h6 {
  color: #fff;
  font-size: 16px;
  text-transform: uppercase;
  margin-top: 5px;
  letter-spacing: 2px;
}
.site-footer a {
  color: #737373;
}
.site-footer a:hover {
  color: #3366cc;
  text-decoration: none;
}
.footer-links {
  padding-left: 0;
  list-style: none;
}
.footer-links li {
  display: block;
}
.footer-links a {
  color: #737373;
}
.footer-links a:active,
.footer-links a:focus,
.footer-links a:hover {
  color: #3366cc;
  text-decoration: none;
}
.footer-links.inline li {
  display: inline-block;
}
.site-footer .social-icons {
  text-align: right;
}
.site-footer .social-icons a {
  width: 40px;
  height: 40px;
  line-height: 40px;
  margin-left: 6px;
  margin-right: 0;
  border-radius: 100%;
  background-color: #33353d;
}
.copyright-text {
  margin: 0;
}

.social-icons {
  padding-left: 0;
  margin-bottom: 0;
  list-style: none;
}
.social-icons li {
  display: inline-block;
  margin-bottom: 4px;
}
.social-icons li.title {
  margin-right: 15px;
  text-transform: uppercase;
  color: #96a2b2;
  font-weight: 700;
  font-size: 13px;
}
.social-icons a {
  background-color: #eceeef;
  color: #818a91;
  font-size: 16px;
  display: inline-block;
  line-height: 44px;
  width: 44px;
  height: 44px;
  text-align: center;
  margin-right: 8px;
  border-radius: 100%;
  -webkit-transition: all 0.2s linear;
  -o-transition: all 0.2s linear;
  transition: all 0.2s linear;
}
.social-icons a:active,
.social-icons a:focus,
.social-icons a:hover {
  color: #fff;
  background-color: #29aafe;
}
.social-icons.size-sm a {
  line-height: 34px;
  height: 34px;
  width: 34px;
  font-size: 14px;
}
.social-icons a.facebook:hover {
  background-color: #3b5998;
}
.social-icons a.twitter:hover {
  background-color: #00aced;
}
.social-icons a.linkedin:hover {
  background-color: #007bb6;
}
.social-icons a.pinteres:hover {
  background-color: #e21a1a;
}
.footer-links a{
   text-decoration: none;
 }

@media all and (max-width: 1200px) {
 .category-footer{
   text-align: center;
 }
 .quick-link-footer{
  text-align: center;
 }

 
}

@media (max-width: 991px) {
  .site-footer [class^="col-"] {
    margin-bottom: 30px;
  }
}
@media (max-width: 768px) {
  .site-footer {
    padding-bottom: 0;
  }
  .site-footer .copyright-text,
  .site-footer .social-icons {
    text-align: center;
  }
  .social-icons li.title {
    display: block;
    margin-right: 0;
    font-weight: 600;
  }
  .category-footer{
   text-align: center;
 }
 .quick-link-footer{
  text-align: center;
 }
}

</style>
