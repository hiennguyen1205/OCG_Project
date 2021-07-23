<template>
  <form>
    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label">Họ tên</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" v-model="User.name" />
      </div>
    </div>

    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label"
        >Số điện thoại</label
      >
      <div class="col-sm-10">
        <input type="text" class="form-control" v-model="User.phone_number" />
      </div>
    </div>

    <div class="row mb-3">
      <label for="inputEmail3" class="col-sm-2 col-form-label">Ho Ten</label>
      <div class="col-sm-10">

        <input
          type="text"
          class="form-control"
          v-model="User.username"
          :readonly="isReadOnly"
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
    <button
      type="submit"
      class=" col-sm-3 btn btn-primary"
      style="margin-left=30px"
      @click="changeInfor($event)"
    >
      Lưu Thông Tin
    </button>
  </form>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  name: "UserInfor",

  data() {
    return {
      User: [],
      name: "",
      phone_number: "",
      username: "",
      email: "",
      address: "",
      isReadOnly: true,
    };
  },

  methods: {
    ...mapMutations("users", ["saveUser"]),
    async changeInfor(event) {
      event.preventDefault();
      console.log(this.User);
      await fetch("http://localhost:3000/api/user", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          id: this.User.id,
          username: this.User.username,
          password: this.User.password,
          email: this.User.email,
          address: this.User.address,
          name: this.User.name,
          phone_number: this.User.phone_number,
          // role: this.user.role,
        }),
      })
        .then(async () => {
          alert("Change infor successfully");
        })
        .catch(async (error) => {
          alert("Something went wrong" + error);
        });
    },
  },

  async beforeMount() {
    await fetch("http://localhost:3000/api/user", {
      method: "GET",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    })
      .then(async (response) => {
        console.log();
        this.User = await response.json();
        this.saveUser(this.User);
      })
      .catch(async (error) => {
        console.log(error);
      });
  },
};
</script>

<style></style>
