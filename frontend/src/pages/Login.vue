<template>
  <div class="login-wrap">
    <div class="login-html">
      <input id="tab-1" type="radio" name="tab" class="sign-in" checked /><label
        for="tab-1"
        class="tab"
        >Sign In</label
      >
      <input id="tab-2" type="radio" name="tab" class="sign-up" /><label
        for="tab-2"
        class="tab"
        >Sign Up</label
      >
      <div class="login-form">
        <form @submit.prevent="Login">
          <div class="sign-in-htm">
            <div class="group">
              <label for="user" class="label">Username</label>
              <input
                id="user-logout"
                v-model="username"
                type="text"
                class="input"
              />
            </div>
            <div class="group">
              <label class="label">Password</label>
              <input
                id="pass-logout"
                v-model="password"
                type="password"
                class="input"
                data-type="password"
                required
              />
            </div>

            <div class="group">
              <input type="submit" class="button" value="Sign In" />
            </div>
            <div class="hr"></div>
            <div class="foot-lnk">
              <a href="#forgot" class="forgot">Forgot Password?</a>
            </div>
          </div>
        </form>
        <form @submit.prevent="Signup">
          <div class="sign-up-htm">
            <div class="group">
              <label for="user" class="label">Username</label>
              <input
                id="user"
                v-model="username"
                type="text"
                class="input"
                required
              />
            </div>
            <div class="group">
              <label class="label">Password</label>
              <input
                id="pass"
                v-model="password"
                type="password"
                class="input"
                data-type="password"
                required
              />
            </div>
            <div class="group">
              <label class="label">Repeat Password</label>
              <input
                id="re-pass"
                v-model="password_confirm"
                type="password"
                class="input"
                data-type="password"
                required
              />
            </div>
            <div class="group">
              <label class="label">Email Address</label>
              <input
                id="email"
                type="email"
                v-model="email"
                class="input"
                required
              />
            </div>
            <div class="group">
              <label class="label">Name</label>
              <input type="text" v-model="name" class="input" required />
            </div>
            <div class="group">
              <label class="label">Phone Number</label>
              <input type="text" v-model="phonenumber" class="input" required />
            </div>
            <div class="group">
              <label class="label">Address</label>
              <input id="addr" type="text" v-model="address" class="input" />
            </div>
            <div class="group">
              <input type="submit" class="button" value="Sign Up" />
            </div>
            <div class="hr"></div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from "vuex";
export default {
  name: "Loginout",
  data() {
    return {
      username: "",
      email: "",
      password: "",
      password_confirm: "",
      address: "",
      name: "",
      phonenumber: "",
    };
  },
  computed: {
    ...mapState("carts", ["order"]),
  },
  methods: {
    ...mapActions("carts", ["getCartByUserId"]),
    ...mapActions("users", ["login", "register","getUser"]),
    Login() {
      this.login({
        username: this.username,
        password: this.password,
      })
        .then(() => {
          this.getCartByUserId(parseInt(document.cookie.slice(3)));
          this.getUser();
        })
        .catch((e) => {
          console.log(e);
        });
    },
    Signup() {
      if (this.password != this.password_confirm) {
        alert("Passwords did not match");
      } else {
        this.register({
          username: this.username,
          password: this.password,
          name: this.name,
          phone_number: this.phonenumber,
          email: this.email,
          address: this.address,
        })
          .then(async () => {
            await this.Login();
            this.$router.push({ name: "Home" });
          })
          .catch((error) => {
            console.log(error);
          });
      }
    },
  },
};
</script>

<style scoped>
body {
  margin: 0;
  color: #6a6f8c;
  background: #c8c8c8;
  font: 600 16px/18px "Open Sans", sans-serif;
}
*,
:after,
:before {
  box-sizing: border-box;
}
.clearfix:after,
.clearfix:before {
  content: "";
  display: table;
}
.clearfix:after {
  clear: both;
  display: block;
}
a {
  color: inherit;
  text-decoration: none;
}

.login-wrap {
  width: 100%;
  margin: auto;
  max-width: 525px;
  min-height: 830px;
  position: relative;
  background: url(https://raw.githubusercontent.com/khadkamhn/day-01-login-form/master/img/bg.jpg)
    no-repeat center;
  box-shadow: 0 12px 15px 0 rgba(0, 0, 0, 0.24),
    0 17px 50px 0 rgba(0, 0, 0, 0.19);
  margin-top: 120px;
  margin-bottom: 20px;
}
.login-html {
  width: 100%;
  height: 100%;
  position: absolute;
  padding: 90px 70px 50px 70px;
  background: rgba(40, 57, 101, 0.9);
}
.login-html .sign-in-htm,
.login-html .sign-up-htm {
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  position: absolute;
  transform: rotateY(180deg);
  backface-visibility: hidden;
  transition: all 0.4s linear;
}
.login-html .sign-in,
.login-html .sign-up,
.login-form .group .check {
  display: none;
}
.login-html .tab,
.login-form .group .label,
.login-form .group .button {
  text-transform: uppercase;
}
.login-html .tab {
  font-size: 22px;
  margin-right: 15px;
  padding-bottom: 5px;
  margin: 0 15px 10px 0;
  display: inline-block;
  border-bottom: 2px solid transparent;
}
.login-html .sign-in:checked + .tab,
.login-html .sign-up:checked + .tab {
  color: #fff;
  border-color: #1161ee;
}
.login-form {
  min-height: 345px;
  position: relative;
  perspective: 1000px;
  transform-style: preserve-3d;
}
.login-form .group {
  margin-bottom: 15px;
}
.login-form .group .label,
.login-form .group .input,
.login-form .group .button {
  width: 100%;
  color: #fff;
  display: block;
}
.login-form .group .input,
.login-form .group .button {
  border: none;
  padding: 15px 20px;
  border-radius: 25px;
  background: rgba(255, 255, 255, 0.1);
}
.login-form .group input[data-type="password"] {
  text-security: circle;
  -webkit-text-security: circle;
}
.login-form .group .label {
  color: #aaa;
  font-size: 12px;
}
.login-form .group .button {
  background: #1161ee;
}
.login-form .group label .icon {
  width: 15px;
  height: 15px;
  border-radius: 2px;
  position: relative;
  display: inline-block;
  background: rgba(255, 255, 255, 0.1);
}
.login-form .group label .icon:before,
.login-form .group label .icon:after {
  content: "";
  width: 10px;
  height: 2px;
  background: #fff;
  position: absolute;
  transition: all 0.2s ease-in-out 0s;
}
.login-form .group label .icon:before {
  left: 3px;
  width: 5px;
  bottom: 6px;
  transform: scale(0) rotate(0);
}
.login-form .group label .icon:after {
  top: 6px;
  right: 0;
  transform: scale(0) rotate(0);
}
.login-form .group .check:checked + label {
  color: #fff;
}
.login-form .group .check:checked + label .icon {
  background: #1161ee;
}
.login-form .group .check:checked + label .icon:before {
  transform: scale(1) rotate(45deg);
}
.login-form .group .check:checked + label .icon:after {
  transform: scale(1) rotate(-45deg);
}
.login-html
  .sign-in:checked
  + .tab
  + .sign-up
  + .tab
  + .login-form
  .sign-in-htm {
  transform: rotate(0);
}
.login-html .sign-up:checked + .tab + .login-form .sign-up-htm {
  transform: rotate(0);
}

.hr {
  height: 2px;
  margin: 60px 0 50px 0;
  background: rgba(255, 255, 255, 0.2);
}
.foot-lnk {
  text-align: center;
}
.forgot {
  color: #fff;
}
</style>
