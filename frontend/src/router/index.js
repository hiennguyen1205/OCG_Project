import { createRouter, createWebHistory } from "vue-router"


import Home from "../pages/Home.vue"
// import About from "./pages/AboutUs.vue"
// import Products from "./pages/Products.vue"
// import Checkout from "./pages/Checkout.vue"
// import Login from "./pages/Login.vue"
// import ProductDetail from "./pages/ProductDetail.vue"

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        // path: '/about', component: About  
        path: '/about',
        name: 'AboutUs',
        component: () => import(/* webpackChunkName: "about" */'../pages/AboutUs.vue')
    },
    {
        // path: '/products', component: Products
        path: '/products',
        name: 'Products',
        component: () => import(/* webpackChunkName: "products" */ '../pages/Products.vue')
    },
    {
        // path: '/checkout', component: Checkout
        path: '/checkout',
        name: 'Checkout',
        component: () => import(/* webpackChunkName: "checkout" */ '../pages/Checkout.vue')
    },
    {
        // path: '/login', component: Login 
        path: '/login',
        name: 'Login',
        component: () => import(/* webpackChunkName: "login" */ '../pages/Login.vue')
    },
    {
        path: '/detail_product/:id',
        name: "ProductDetail",
        component: () => import(/* webpackChunkName: "detail_product/:id" */ '../pages/ProductDetail.vue'),
        props: true
    },

]

const router = createRouter({
    history: createWebHistory(),
    scrollBehavior() {
        return { top: 0 };
    },
    routes, // short for `routes: routes`
})

export default router