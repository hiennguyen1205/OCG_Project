<template>
  <nav
    class="navbar navbar-expand-lg navbar-light bg-light fixed-top"
    style="padding: 0px;"
  >
    <div class="container-fluid">
      <router-link to="/"><img :src="logoShop" alt=""/></router-link>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0 text-menu">
          <li class="nav-item">
            <router-link to="/">Home</router-link>
            <!-- <a class="nav-link " aria-current="page" href="#">Home</a> -->
          </li>

          <li class="nav-item dropdown">
            <router-link to="/products">Product</router-link>
          </li>

          <!-- <li class="nav-item">
            <a class="nav-link" href="#">Contact Us</a>
          </li> -->

          <li class="nav-item">
            <router-link to="/about">About</router-link>
            <!-- <a class="nav-link " href="#">About</a> -->
          </li>
        </ul>

        <div class="left-nav">
          <router-link to="/checkout">
            <i class="fas fa-shopping-cart"
              ><sup>{{ totalProductsInCart }}</sup></i
            ></router-link
          >
          <!-- <i class="fas fa-search"></i> -->
          <a v-if="isAuthenticated" @click.prevent="logOut">
            <i class="fas fa-sign-out-alt"></i>
          </a>
          <router-link v-else to="/login"
            ><i class="fas fa-user-alt"></i
          ></router-link>
        </div>
      </div>
    </div>
  </nav>
  <router-view></router-view>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      logoShop: require('@/assets/images/logos/nha-xinh-logo.jpg'),
    };
  },
  methods: {
    async logOut() {
      await fetch('http://localhost:3000/api/logout', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      })
        .then(async () => {
          await this.$store.dispatch('setAuth', false);
          await this.$store.dispatch('emptyListProducts');
          this.$router.push({ name: 'Home' });
        })
        .catch(() => {
          console.log('server failed');
        });
    },
  },

  computed: {
    totalProductsInCart() {
      return this.$store.state.order.products != null
        ? this.$store.state.order.products.length
        : '';
    },
    isAuthenticated() {
      return this.$store.state.authenticated;
    },
  },
};
</script>

<style scoped>
.name-shop {
  margin-left: 20px;
  font-size: 25px;
  color: rgb(5, 5, 5);
}
.container-fluid img {
  width: 91px;
  height: 37px;
}
.name-shop:hover {
  color: rgb(0, 158, 137);
}
.text-menu {
  width: 100%;
  justify-content: center;
}

a {
  text-decoration: none;
  padding: 20px 30px;
  border-bottom: 3px solid transparent;
  transition: 0.4s;
  font-size: 20px;
  color: rgb(87, 87, 87);
}
a::before {
  transform: scale(1.1);
}
nav .nav-item {
  padding: 10px 20px;
  border-bottom: 3px solid transparent;
  transition: 0.4s;
  font-size: 20px;
}
nav .nav-item.active,
nav .nav-item:hover {
  background-color: rgb(177, 175, 175);
  border-bottom-color: rgb(112, 255, 207);
}

nav div {
  background-color: rgb(196, 196, 196);
}

.left-nav {
  display: flex;
  justify-content: space-between;
}
i {
  font-size: 30px;
  margin-left: 20px;
}
i:hover {
  color: rgb(0, 158, 137);
}

.router-link-active {
  color: #f33f3f;
}
</style>
