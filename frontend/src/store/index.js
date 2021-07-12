import { createStore } from "vuex";

// Tạo store
const store = createStore({
    state() {
        return {
            userId: -1,
            products: [],

            listBackupProducts: [],

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
        };
    },

    getters: {
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

        // calcTotal(getters) {
        //     return getters.calcSubTotal - getters.discount + getters.calcTax;
        // },
    },

    mutations: {
        updateUserId(state, id) {
            state.userId = id;
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
                console.log(1);
                state.products.push(product);
            } else {
                console.log(2);
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

    },

    // Giống mutations nhưng dùng cho hàm async
    actions: {

    },
});

export default store;
