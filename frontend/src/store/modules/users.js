import router from "../../router/index.js"
const state = () => ({
    authenticated: false,
});

const getters = {
    getUserId() {
        return parseInt(document.cookie.slice(3,));
    },
};



const actions = {
    login: async ({ commit }, user) => {
        console.log("user ", user);
        await fetch('http://localhost:3000/api/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(user),
        })
            .then((response) => {
                console.log(response);
                if (response.status ===200) {
                    commit('setAuth',true);
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

    register: ({ commit }, user) => {
        fetch("http://localhost:3000/api/register", {
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
};

const mutations = {
    // updateUserId(state) {
    //     state.userId = parseInt(document.cookie.slice(3,));
    // },
    setAuth(state, auth) {
        state.authenticated = auth;
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}