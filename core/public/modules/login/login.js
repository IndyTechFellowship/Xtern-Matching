
function setToken(token, type) {
    //localStorage.setItem('token', token);
    sessionStorage.setItem(type, token);
}

function getToken(tokenName) {
    return sessionStorage.getItem(tokenName);
}

function logout() {
    let tokens = ['auth', 'role', 'organization', 'company', 'jwtToken'];
    tokens.forEach(function(token){
        removeToken(token);
    });
    sessionStorage.removeItem("userKey");
}

function removeToken(token){
    sessionStorage.removeItem(token)
}

let isLoggedInTechPoint = function ($q) {
    //TODO change this to type
    var role = getToken("organization");
    if (!role) {
        var errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
        $q.reject(errorObject);
        return $q.reject(errorObject);
    }
    else if(role =="admin"){
        return;
    }
    else if (role == "Techpoint") {
        return;
    }
    else if(role == "Instructor"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_INSTRUCTOR' }
        return $q.reject(errorObject);
    }
    else if(role == "Salesforce"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);
    }
    else{
        var errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);
    }
};

var isLoggedInInstructor = function ($q) {
    var role = getToken("role");
    if (!role) {
        var errorObject = { code: 'NOT_AUTHENTICATED_INSTRUCTOR' };
        return $q.reject(errorObject);
    }
    else if(role =="admin"){
        return;
    }
    else if (role == "TechPoint") {
        var errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);      
    }
    else if(role == "Company"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);        
    }
    else if(role == "Instructor"){
        return;      
    }
    else{
        var errorObject = { code: 'NOT_AUTHENTICATED_INSTRUCTOR' };
        return $q.reject(errorObject);
    }
};

var isLoggedInCompany = function ($q) {
    //TODO change this to kind
    var role = getToken("organization");
    if (!role) {
        var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);

    }
    else if(role =="admin"){
        return;
    }
    else if (role == "Techpoint") {
        var errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);
    }
    else if(role == "Instructor"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_INSTRUCTOR' };
        return $q.reject(errorObject);
    }
    else if(role == "Salesforce"){
        return;
    }
    else{
        var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);
    }
};

var isLoggedIn = function ($q, code) {
    var role = getToken("organization");
    if (!role) {
        return;
    }
    else if(role =="admin"){
        var errorObject = {code: code};
        return $q.reject(errorObject);
    }
    else if (role == "TechPoint") {
        var errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);   
    }
    else if(role == "Company"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);        
    }
    else if(role == "Instructor"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_INSTRUCTOR' };
        return $q.reject(errorObject);        
    }
};