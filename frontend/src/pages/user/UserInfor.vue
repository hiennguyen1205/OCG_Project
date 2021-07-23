<template>
  <form @submit.prevent="changeInfor">
    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label">Ho Ten</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" :value="User.name" />
      </div>
    </div>
    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label">SDT</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" :value="User.phone_number" />
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
import { mapMutations, mapState } from 'vuex';

export default {
  name: 'UserInfor',

  data() {
    return {
      User: [],
    };
  },

  computed: {
    ...mapState('users', ['user']),
  },

  methods: {
    ...mapMutations('users', ['saveUser']),
    async changeInfor() {
      await fetch('http://localhost:3000/api/user', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          id: this.user.id,
          username: this.user.username,
          pass: this.user.password,
          name: this.user.name,
          phone_number: this.user.phone_number,
          email: this.User.email,
          address: this.User.address,
          role: this.user.role,
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
        this.saveUser(this.User);
        console.log(this.user);
      })
      .catch(async (error) => {
        console.log(error);
      });
  },
};
</script>

<style></style>
