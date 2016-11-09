
function setToken(token, type) {
    //localStorage.setItem('token', token);
    localStorage.setItem(type, token);
}

function getToken(tokenName) {
    return localStorage.getItem(tokenName);
}

function logout() {
    localStorage.removeItem("auth");
    localStorage.removeItem("role");
    localStorage.removeItem("organization");
}

var isLoggedInTechPoint = function () {
    var role = getToken("role");
    if (!role) {
        var errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
        return false;
    }
    else if(role =="admin"){
        return true;
    }
    else if (role == "TechPoint") {
        return true;
    }
    else if(role == "Company"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return errorObject;
    }
    else if(role == "Instructor"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_INSTRUCTOR' };
        return errorObject;
    }
    else{
        var errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
        return errorObject;
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
    var role = getToken("role");
    if (!role) {
        var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
        return false;
    }
    else if(role =="admin"){
        return true;
    }
    else if (role == "TechPoint") {
        var errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return errorObject;
    }
    else if(role == "Company"){
        return true;
    }
    else if(role == "Instructor"){
        var errorObject = { code: 'ALREADY_AUTHENTICATED_INSTRUCTOR' };
        return errorObject;
    }
    else{
        var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
        return errorObject;
    }
};

var isLoggedIn = function ($q) {
    var role = getToken("role");
    if (!role) {
        return;
    }
    else if(role =="admin"){
        return;// implemented reroute
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
