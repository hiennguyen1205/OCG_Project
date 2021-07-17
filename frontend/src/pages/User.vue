<template>
  <div class="container-fluid buffer-top">
    <div class="row flex-nowrap">
      <div class="col-auto col-md-3 col-xl-2 px-sm-2 px-0 bg-dark">
        <div
          class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100"
        >
          <a
            href="/"
            class="d-flex align-items-center pb-3 mb-md-0 me-md-auto text-white text-decoration-none"
          >
            <span class="fs-5 d-none d-sm-inline">Hồ sơ cá nhân</span>
          </a>
          <ul
            class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start"
            id="menu"
          >
            <li class="nav-item">
              <router-link
                :to="{ name: 'UserInfor' }"
                class="nav-link align-middle px-0"
              >
                <i class="fs-4 bi-house"></i>
                <span class="ms-1 d-none d-sm-inline">Thông tin</span>
              </router-link>
            </li>
            <li class="nav-item">
              <router-link
                :to="{ name: 'UserPassword' }"
                class="nav-link align-middle px-0"
              >
                <i class="fs-4 bi-house"></i>
                <span class="ms-1 d-none d-sm-inline">Mật khẩu</span>
              </router-link>
            </li>

            <li class="nav-item">
              <router-link
                :to="{ name: 'Admin' }"
                class="nav-link align-middle px-0"
              >
                <i class="fs-4 bi-house"></i>
                <span class="ms-1 d-none d-sm-inline">Danh sách sản phẩm</span>
              </router-link>
            </li>
          </ul>
          <hr />
        </div>
      </div>
      <div class="col py-3">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'User',
};
</script>

<style scoped>
.buffer-top {
  margin-top: 80px;
}
@media only screen and (max-width: 768px) {
  .buffer-top {
    margin-top: 130px;
  }
}
</style>
