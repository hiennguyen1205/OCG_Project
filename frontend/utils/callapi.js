const rootPath = "http://localhost:3000/api"
const authenticatedPath = "http://localhost:3000/auth/api"

async function fetchProducts() {
    return await fetch(`${rootPath}/products?limit=${this.limit}&cursor=${this.cursor}&categoryId=${this.categoryId}&search=${this.searchProducts}&sort=${sort}&isFeature=${isfeatured}`, {
        method: httpMethod,
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
    })
}
async function fetchUser(httpMethod) {
    if (httpMethod == "GET") {
        return await fetch(`${rootPath}/users`, {
            method: httpMethod,
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
        })
    } else {
        //Post request
    }
}