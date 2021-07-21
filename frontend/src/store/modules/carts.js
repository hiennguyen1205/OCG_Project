

const state = () => ({
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
    order: {
        user: {},
        products: [],
    }
});

const getters = {
    getLengthListProducts(state) {
        if (state.order.products === null || state.order.products===undefined) {return 0}
        return state.order.products.length
    },
    calcSubTotal(state) {
        if (state.order.products != null) {
            return state.order.products.reduce(
                (totalPrice, product) => totalPrice + product.price * product.quantity,
                0
            );
        } else {
            return 0;
        }
    },
    calcTax(state) {
        if (state.order.products != null) {
            return (
                state.order.products.reduce(
                    (totalPrice, product) =>
                        totalPrice + product.price * product.quantity,
                    0
                ) / 10
            );
        } else {
            return 0;
        }

    },

    // calcTotalPrice(state) {
    //     if (state.order.products === null || state.order.products == undefined) return state.totalPrice=0
    //     return state.totalPrice = state.order.products.reduce((totalPrice, product) => totalPrice + product.price * product.quantity,
    //         0)
    // },
    emptyListProducts(state) {
        state.order = {};
    },
};



const actions = {
    async submitOrder(_,order) {
        await fetch('http://localhost:3000/auth/api/orders', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(order),
        })
            .then(() => {
                // commit("submitedOrder")  
                // console.log(state.order);
                console.log("OK");
            })
            .catch(err => console.log(err))
    },
    async getCartByUserId({ commit }, userId) {
        
        await fetch('http://localhost:3000/api/orders/' + userId, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
        })
            .then(async (response) => {
               
                const data = await response.json()
                commit("GET_CART", data)
                // console.log(data);
            })
            .catch(err => console.log(err))
    },
};

const mutations = {
    calcDiscount(state) {
        let s = state.promotions.filter(
            (promotion) => promotion.code === state.promoCode
        );

        if (s.length > 0) {
            state.discount =
                (parseInt(s[0].discount.substr(0, 2)) / 100) * state.order.products.reduce(
                    (totalPrice, product) => totalPrice + product.price * product.quantity,
                    0
                );
        } else state.discount = 0;
        // if (state.order.products === null || state.order.products===undefined){
        //     state.discount = 0 
        //     return
        // }
        // state.discount = state.order.products.reduce((salePrice, product) => salePrice += product.quantity * product.price * product.sale/100, 0)
        // console.log(1);
        // console.log(state.products);
        // console.log(state.discount);
    },

    changeQuantity(state, { productId, number }) {
        state.order.products = state.order.products.map((product) => {
            if (product.id === productId) {
                if (parseInt(number) > 0 && parseInt(number) < 300) {
                    product.quantity = parseInt(number);
                    return product;
                } else {
                    product.quantity = 1;
                    return product;
                }
            }
            return product;
        });
    },

    increaseQuantity(state, productId) {
        state.order.products = state.order.products.map((product) => {
            if (product.id === productId) {
                product.quantity++;
            }
            return product;
        });
    },

    decreaseQuantity(state, productId) {
        state.order.products = state.order.products.map((product) => {
            if (product.id === productId) {
                product.quantity--;
                if (product.quantity < 1) {
                    product.quantity = 1;
                }
            }
            return product;
        });
    },

    removeItem(state, productId) {
        let confirmDelete = confirm("Do you want to delete state product " + productId + "??");
        if (confirmDelete) {
            state.order.products = state.order.products.filter(
                (product) => product.id != productId
            );
        }
    },

    changeDiscountCode(state, value) {
        state.promoCode = value;
    },

    addProductToCart(state, product) {
        if (state.order.products != null) {
            let checkProduct = state.order.products.filter((productInStore) => productInStore.id === product.id);
            if (checkProduct.length === 0) {
                state.order.products.push(product);
            } else {
                state.order.products.forEach(element => {
                    if (element.id === product.id) {
                        element["quantity"] += product["quantity"];
                    }
                });
            }
        } else {
            state.order.products = [];
            state.order.products.push(product);
        }
    },

    saveOrder(state) {
        state.order.products.forEach(product => {
            product["active"] = 0;
        });
    },

    GET_CART(state, data) {
        state.order = data;
        // console.log(state.order);
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}