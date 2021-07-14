<template>
  <form>
    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label">Tên</label>
      <div class="col-sm-10">
        <input type="email" class="form-control" id="inputEmail3" />
      </div>
    </div>
    <div class="row mb-3">
      <label for="inputPassword3" class="col-sm-2 col-form-label">Email</label>
      <div class="col-sm-10">
        <input type="password" class="form-control" id="inputPassword3" />
      </div>
    </div>
    <div class="row mb-3">
      <label for="inputPassword3" class="col-sm-2 col-form-label"
        >Địa chỉ</label
      >
      <div class="col-sm-10">
        <input type="password" class="form-control" id="inputPassword3" />
      </div>
    </div>

    <button type="submit" class="btn btn-primary">Lưu Thông Tin</button>
  </form>
</template>

<script>
export default {
  name: 'UserInfor',
};
</script>

<style></style>
