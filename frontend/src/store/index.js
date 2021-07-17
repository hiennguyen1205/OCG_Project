import { createStore } from "vuex";
import carts from "./modules/carts";
import users from "./modules/users";
import products from "./modules/products";
// Tạo store
const store = createStore({
    modules: {
        carts,
        users,
        products
    }
});
export default store;
