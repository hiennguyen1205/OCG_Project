<template>

<div id="carouselExampleIndicators" class="carousel slide top-buffer" data-bs-ride="carousel">
  <div class="carousel-indicators">
    <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
    <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="1" aria-label="Slide 2"></button>
    <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="2" aria-label="Slide 3"></button>
  </div>
  <div class="carousel-inner">
    <div class="carousel-item active">
      <img src="@/assets/images/slides_banner/slider-image-1-1920x600.jpg" class="d-block w-100" alt="">
      <div class="carousel-caption  d-md-block">
          <h2>Bàn phòng khách</h2>
          <h4>Sản phẩm đang được giảm giá với nhiều ưu đãi</h4>
        </div>
    </div>
    <div class="carousel-item">
      <img src="@/assets/images/slides_banner/slider-image-2-1920x600.jpg" class="d-block w-100" alt="">
      <div class="carousel-caption  d-md-block">
          <h2>Sofa phòng khách</h2>
          <h4>Nhanh tay sắm nhiều sản phẩm được giảm giá</h4>
        </div>
    </div>
    <div class="carousel-item">
      <img src="@/assets/images/slides_banner/slider-image-3-1920x600.jpg" class="d-block w-100" alt="">
      <div class="carousel-caption  d-md-block">
          <h2>Ghế dùng để decor phòng</h2>
          <h4>Sản phẩm đang được giảm giá với nhiều ưu đãi</h4>
        </div>
    </div>
  </div>
  <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="prev">
    <span class="carousel-control-prev-icon" aria-hidden="true"></span>
    <span class="visually-hidden">Previous</span>
  </button>
  <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="next">
    <span class="carousel-control-next-icon" aria-hidden="true"></span>
    <span class="visually-hidden">Next</span>
  </button>
</div>

</template>

<script>
export default {
  name: "Banner",
};
</script>

<style scoped>
.top-buffer{
  margin-top:70px;
}
.carousel-caption h2 {
  font-size: 50px;
  color: #f33f3f;
  display: block;
}
.carousel-caption h4 {
  font-size: 25px;
  color: #8b0202;
  display: block;
}


</style>
