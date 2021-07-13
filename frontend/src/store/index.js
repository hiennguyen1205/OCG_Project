import { createStore } from "vuex";

// Tạo store
const store = createStore({
    state() {
        return {
            authenticated: false,
            promotions: [
                {
                    code: "SUMMER",
                    discount: "50%",
                },
                {
                    code: "AUTUMN",
                    discount: "40%",
                },
                {
                    code: "WINTER",
                    discount: "30%",
                },
            ],
            promoCode: "",
            discount: 0,
            tax: 0,
            sortType: "",

            products: [],
            user: {},
            order :{},
        };
    },

    getters: {
        getUserId() {
            return parseInt(document.cookie.slice(3,));
        },
        calcSubTotal(state) {
            return state.products.reduce(
                (totalPrice, product) => totalPrice + product.price * product.quantity,
                0
            );

        },
        calcTax(state) {
            return (
                state.products.reduce(
                    (totalPrice, product) =>
                        totalPrice + product.price * product.quantity,
                    0
                ) / 10
            );
        },

        calcTotalPrice(state) {
            state.totalPrice = state.products.reduce((totalPrice, product) => totalPrice + product.price * product.quantity,
                0)
        }
    },

    mutations: {
        updateUserId(state) {
            state.userId = parseInt(document.cookie.slice(3,));
        },
        calcDiscount(state) {
            let s = state.promotions.filter(
                (promotion) => promotion.code === state.promoCode
            );

            if (s.length > 0) {
                console.log(state.calcSubTotal);
                state.discount =
                    (parseInt(s[0].discount.substr(0, 2)) / 100) * state.products.reduce(
                        (totalPrice, product) => totalPrice + product.price * product.quantity,
                        0
                    );
            } else state.discount = 0;

        },

        changeQuantity(state, { productId, number }) {
            console.log(number, productId);
            state.products = state.products.map((product) => {
                if (product.id === productId) {
                    if (parseInt(number) > 0 && parseInt(number) < 300) {
                        product.quantity = parseInt(number);
                        return product;
                    } else {
                        console.log(1);
                        product.quantity = 0;
                        return product;
                    }
                }
                return product;
            });
        },

        removeItem(state, productId) {
            let confirmDelete = confirm("Do you want to delete state product " + productId + "??");
            if (confirmDelete) {
                state.products = state.products.filter(
                    (product) => product.id != productId
                );
            }
        },

        undoProduct(state) {
            state.products = [...state.listBackupProducts]
        },

        changeDiscountCode(state, value) {
            state.promoCode = value;
        },

        addProductToCart(state, product) {
            let checkProduct = state.products.filter((productInStore) => productInStore.id === product.id);
            // console.log(checkProduct.length===0);
            if (checkProduct.length === 0) {
                // console.log(1);
                state.products.push(product);
            } else {
                // console.log(2);
                state.products.forEach(element => {
                    if (element.id === product.id) {
                        element["quantity"] += product["quantity"];
                    }
                });
            }

        },

        selectOption(state, sort) {
            state.sortType = sort;
            console.log("sort " + state.sortType);
            if (state.sortType === "price-asc") {
                return state.products.sort(function (product1, product2) {
                    return product1.price - product2.price;
                });
            } else if (state.sortType === "price-desc") {
                return state.products.sort(function (product1, product2) {
                    return product2.price - product1.price;
                });
            } else {
                return state.products.sort(function (product1, product2) {
                    return product1.id - product2.id;
                });
            }
        },

        saveOrder(state) {
            state.order.products.forEach(product => {
                product["active"] = 0;
            });
        },
        submitedOrder(state) {
            state.products = []

        },
        GET_CART(state, data){
            state.order = data;
            console.log(state.products);
        },
        SET_AUTH(state, auth) {
            state.authenticated = auth;
        }
    },

    // Giống mutations nhưng dùng cho hàm async
    actions: {
        async submitOrder({ commit }, order) {
            await fetch('http://localhost:3000/api/orders', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify(order),
            })
                .then(() => {
                    commit("submitedOrder")
                    console.log("OK");
                })
                .catch(err => console.log(err))
        },
        async getCartByUserId({ commit }, userId) {
            await fetch('http://localhost:3000/api/orders/'+userId, {
                method: 'GET',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
            })
                .then(async (response) => {
                    const data = await response.json()
                    commit("GET_CART", data)
                    console.log("ĐƯỢC");
                })
                .catch(err => console.log(err))
        }
                    console.log(order);
                    console.log("OK");
                })
                .catch(err => console.log(err))

        },

        setAuth: ({ commit }, auth) => {
            commit("SET_AUTH", auth)
        },
    },
});
export default store;
