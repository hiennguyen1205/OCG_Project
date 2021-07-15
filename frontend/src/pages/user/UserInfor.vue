<template>
  <form @submit.prevent="changeInfor">
    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label">Tài khoản</label>
      <div class="col-sm-10">
        <input
          type="text"
          class="form-control"
          :value="User.username"
          readonly
        />
      </div>
    </div>
    <div class="row mb-3">
      <label for="inputPassword3" class="col-sm-2 col-form-label">Email</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" v-model="User.email" />
      </div>
    </div>
    <div class="row mb-3">
      <label for="inputPassword3" class="col-sm-2 col-form-label"
        >Địa chỉ</label
      >
      <div class="col-sm-10">
        <input type="text" class="form-control" v-model="User.address" />
      </div>
    </div>

    <button type="submit" class="btn btn-primary">Lưu Thông Tin</button>
  </form>
</template>

<script>
export default {
  name: 'UserInfor',

  data() {
    return {
      User: [],
    };
  },
  methods: {
    async changeInfor() {
      await fetch('http://localhost:3000/api/change-infor', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          email: this.User.email,
          address: this.User.address,
        }),
      })
        .then(async () => {
          alert('Change infor successfully');
        })
        .catch(async (error) => {
          alert('Something went wrong' + error);
        });
    },
  },

  async beforeMount() {
    await fetch('http://localhost:3000/api/user', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    })
      .then(async (response) => {
        this.User = await response.json();
        console.log(this.User);
      })
      .catch(async (error) => {
        console.log(error);
      });
  },
};
</script>

<style></style>
