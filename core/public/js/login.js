function isLoggedIn() {
    //return true;
    //console.log(localStorage.getItem('token'), localStorage.getItem(type));
    return localStorage.getItem('auth');
    //problems with multiple types logged in
};

function setToken(token, type) {
    //localStorage.setItem('token', token);
    localStorage.setItem(type, token);
}

function getToken(tokenName) {
    return localStorage.getItem(tokenName);
}

function logoutStorage(type) {
    localStorage.removeItem(type);
}
