'use strict';
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

var isLoggedInTechPoint = function ($q) {
    var errorObject ={};
    var role = getToken("organization");
    if (!role) {
        errorObject = { code: 'NOT_AUTHENTICATED_REVIEWER' };
        $q.reject(errorObject);
        return $q.reject(errorObject);
    }
    else if(role =="admin"){
        return null;
    }
    else if (role == "Techpoint") {
        return null;
    }
    else if(role == "Reviewers"){
        errorObject = { code: 'ALREADY_AUTHENTICATED_REVIEWER' };
        return $q.reject(errorObject);
    }
    else if(role == "Company"){
        errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);
    }
    else{
        errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);
    }
};

var isLoggedInReviewer = function ($q) {
    var role = getToken("organization");
    var errorObject ={};
    if (!role) {
        errorObject = { code: 'NOT_AUTHENTICATED_REVIEWER' };
        return $q.reject(errorObject);
    }
    else if(role =="admin"){
        return null;
    }
    else if (role == "TechPoint") {
        errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);      
    }
    else if(role == "Company"){
        errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);
    }
    else if(role == "Reviewers"){
        return null;
    }
    else{
        errorObject = { code: 'NOT_AUTHENTICATED_INSTRUCTOR' };
        return $q.reject(errorObject);
    }
};

var isLoggedInCompany = function ($q) {
    var role = getToken("organization");
    var errorObject ={};
    if (!role) {
        errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);

    }
    else if(role =="admin"){
        return null;
    }
    else if (role == "Techpoint") {
        errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);
    }
    else if(role == "Reviewers"){
        errorObject = { code: 'ALREADY_AUTHENTICATED_REVIEWER' };
        return $q.reject(errorObject);
    }
    else if(role == "Company"){
        return null;
    }
    else{
        errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);
    }
};

var isLoggedIn = function ($q, code) {
    var role = getToken("organization");
    var errorObject ={};
    if (!role) {
        return;
    }
    else if(role =="admin"){
        errorObject = {code: code};
        return $q.reject(errorObject);
    }
    else if (role == "TechPoint") {
        errorObject = { code: 'ALREADY_AUTHENTICATED_TECHPOINT' };
        return $q.reject(errorObject);   
    }
    else if(role == "Company"){
        errorObject = { code: 'ALREADY_AUTHENTICATED_COMPANY' };
        return $q.reject(errorObject);        
    }
    else if(role == "Reviewers"){
        errorObject = { code: 'ALREADY_AUTHENTICATED_REVIEWER' };
        return $q.reject(errorObject);        
    }
};