<template>
  <form @submit.prevent="changePassword">
    <div class="row mb-3">
      <label for="inputPassword3" class="col-sm-2 col-form-label"
        >Mật khẩu mới</label
      >
      <div class="col-sm-10">
        <input type="password" class="form-control" v-model="password" />
      </div>
    </div>
    <div class="row mb-3">
      <label for="inputPassword3" class="col-sm-2 col-form-label"
        >Nhập lại mật khẩu mới</label
      >
      <div class="col-sm-10">
        <input
          type="password"
          class="form-control"
          v-model="password_confirm"
        />
      </div>
    </div>

    <button type="submit" class="btn btn-primary">Thay đổi mật khẩu</button>
  </form>
</template>

<script>
export default {
  name: 'UserPassword',
  data() {
    return {
      password: '',
      password_confirm: '',
    };
  },
  methods: {
    async changePassword() {
      if (this.password != this.password_confirm) {
        alert('Passwords did not match');
      } else {
        await fetch(`${process.env.VUE_APP_ROOT_API}/api/change-password`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
          body: JSON.stringify({
            password: this.password,
          }),
        })
          .then(async () => {
            alert('Change password successfully');
          })
          .catch(async (error) => {
            alert('Something went wrong' + error);
          });
      }
    },
  },
};
</script>

<style></style>
