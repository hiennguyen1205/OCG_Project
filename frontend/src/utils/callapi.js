var baseURL = `${process.env.VUE_APP_ROOT_API}/api/`
var baseAuthURL = `${process.env.VUE_APP_ROOT_API}/auth/api/`
async function PostData(url = '', data = {}) {
    console.log("postdata")
    // Default options are marked with *
    var Url = baseURL + url
    const response = await fetch(Url, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        credentials: 'include', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        // referrerPolicy: 'no-referrer',
        body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return response.json(); // parses JSON response into native JavaScript objects
}

async function GetData(url = '') {
    var Url = baseURL + url
    const response = await fetch(Url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
    });
    return response.json();
}

async function DeleteData(url = '') {
    var Url = baseURL + url
    const response = await fetch(Url, {
        method: 'DELETE',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
    });
    return response;
}

async function PutData(url = '', data = {}) {
    var Url = baseURL + url
    const response = await fetch(Url, {
        method: 'PUT',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        body: JSON.stringify(data)
    });
    return response;
}

async function AuthPutData(url = '', data = {}) {
    var Url = baseAuthURL + url
    const response = await fetch(Url, {
        method: 'PUT',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        body: JSON.stringify(data)
    });
    return response;
}

async function AuthGetData(url = '') {
    var Url = baseAuthURL + url
    const response = await fetch(Url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
    });
    return response.json();
}

export { PostData, GetData, DeleteData, PutData, AuthPutData, AuthGetData }