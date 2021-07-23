import router from "../../router/index.js"
import { GetData } from "../../utils/callapi.js";
const state = () => ({
    authenticated: false,
    user: {},
});

const getters = {
    getUserId() {
        return parseInt(document.cookie.slice(3,));
    },
};



const actions = {
    login: async ({ commit }, user) => {
        await fetch('http://localhost:3000/api/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(user),
        })
            .then((response) => {
                if (response.status === 200) {
                    //đăng nhập xong lưu user
                    
                    commit('setAuth', true);
                    router.push({ name: "Home" });
                } else {
                    alert("Tài khoản hoặc mật khẩu sai!!!")
                    return
                }
            })
            .catch((e) => {
                console.log(e);
            })
    },

    register: async ({ commit }, user) => {
        await fetch("http://localhost:3000/api/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(user),
        })
            .then(() => {
                commit('setAuth', true);
                alert("Password created successfully");
            })
            .catch((error) => {
                commit('setAuth', false);
                console.log(error);
            });
    },

    logout: ({ commit }) => {
        fetch("http://localhost:3000/api/logout", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
        })
            .then(() => {
                commit('setAuth', false);
                router.push({ name: "Home" });
            })
            .catch(() => {
                console.log("server failed");
            });
    },

    getUser: async ({commit}) => {
        let user = await GetData(`user`);
        commit('saveUser', user);
    }
};

const mutations = {
    emptyUser(state) {
        state.user = {};
    },
    // updateUserId(state) {
    //     state.userId = parseInt(document.cookie.slice(3,));
    // },
    setAuth(state, auth) {
        state.authenticated = auth;
    },
    saveUser(state, userInput) {
        state.user = userInput;
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}