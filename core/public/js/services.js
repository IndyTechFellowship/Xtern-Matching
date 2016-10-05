/**
 * Created by Steven on 3/28/2016.
 */
 (function (){
    var app = angular.module('DataManager',[]);

    app.service('ProfileService', ['$http', function ($http){
        var self = this;
        self.profile = null;

        self.getStudentDataForId = function(id, callback){
            console.log(id);
            if(!self.profile || self.profile._id !== id) {
                $http({
                    method: 'GET',
                    url: "http://localhost:8080/student/" + id,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    console.log('get student data:');
                    console.log(data.data);
                    self.profile = data.data;
                    callback(self.profile);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('', 'err');
                });
            } else {
                 callback(self.profile);
            }
        };

    }]).service('CompanyService', ['$http', function ($http){
        var self = this;
        self.company = null;

        self.getCompanyDataForId = function(id, callback){
            // console.log(id);
            if(!self.company || self.company._id !== id) {
                $http({
                    method: 'GET',
                    // TODO: replace this id when company login is done
                    url: "http://localhost:8080/company/" + 5047308127305728,
                    // url: "http://localhost:8080/company/" + id,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    console.log('get company data:');
                    console.log(data.data);
                    self.company = data.data;
                    callback(self.company);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('', 'err');
                });
            } else {
                 callback(self.company);
            }
        };
    }]).service('TechPointDashboardService',['$http', function ($http){
        var self = this;
        self.userSummaryData = null;

        self.queryUserSummaryData = function(callback){
            $http({
                method: 'GET',
                url: "http://localhost:8080/student",
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function (data) {
                self.userSummaryData = data.data;
                callback(self.userSummaryData);
            }, function errorCallback(response) {
                console.log('error occured: '+response);
                console.log('Here: '+getToken('auth'));
                callback('','err');
            });
            // if(self.userSummaryData){
            //     $http({
            //         method: 'GET',
            //         url: "localhost:8080/student",
            //         headers: {
            //             'Content-Type': "application/json",
            //             'Accept': "application/json"
            //         }
            //     }).then(function (data) {
            //         self.userSummaryData = data.data;
            //         callback(self.userSummaryData);
            //     });
            // }
            // else{
            //     callback(self.userSummaryData);
            // }
        };
    }]).service('AuthService',['$http', function ($http) {
        var self = this;
        self.jwtToken = null;

        self.login = function(email,password,callback){
            $http.post("http://localhost:8080/auth/login",{"email":email, "password": password}).then(function(data) {
                self.jwtToken = data.data['token'];
                //console.log('Here: '+self.jwtToken);
                callback(self.jwtToken);
            }, function errorCallback(response) {
                console.log('error occured: '+response);
                callback('','err');
            });
        };
    }]);
})();
