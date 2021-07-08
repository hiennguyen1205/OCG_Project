// import { createApp } from "vue"
// import { createRouter, createWebHistory } from "vue-router"
// import App from "./App.vue"

// import Home from "./pages/Home.vue"
// import About from "./pages/AboutUs.vue"
// import Products from "./pages/Products.vue"
// import Checkout from "./pages/Checkout.vue"
// import Login from "./pages/Login.vue"
// import ProductDetail from "./pages/ProductDetail.vue"

// const routes = [
//     { path: '/', component: Home },
//     { path: '/about', component: About },
//     { path: '/products', component: Products  },
//     { path: '/checkout', component: Checkout },
//     { path: '/login', component: Login },
//     { path: '/detail_product/:id', name: "ProductDetail", component: ProductDetail, props:true },

// ]

// const router = createRouter({
//     history: createWebHistory(),
//     routes, // short for `routes: routes`
// })
// const app = createApp(App)
// app.use(router)
// app.mount('#app')



import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'
import store from './store/index'

createApp(App)
.use(router)
.use(store)
.mount('#app')
