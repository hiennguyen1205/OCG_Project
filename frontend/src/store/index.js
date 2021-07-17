import { createStore } from "vuex";
import carts from "./modules/carts";
import users from "./modules/users";
// Tạo store
const store = createStore({
    modules: {
        carts,
        users
    }
});
export default store;
