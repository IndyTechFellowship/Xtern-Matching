/**
 * Created by Steven on 3/28/2016.
 */
 (function (){
    var app = angular.module('DataManager',[]);
	//var host = "http://xtern-matching.appspot.com/"
	//var host = "http://xtern-matching-143216.appspot.com/" //DEV Server
	var host = "http://localhost:8080/";
    app.service('ProfileService', ['$http', function ($http){
        var self = this;
        self.profile = null;

        self.getStudentDataForId = function(id, callback){
            console.log(id);
            if(!self.profile || self.profile._id !== id) {
                $http({
                    method: 'GET',
                    url: host + "student/" +id,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    console.log('SUCCESS: get student data', data.data);
                    self.profile = cleanStudents(data.data);
                    callback(self.profile);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('', 'err')
                });
            } else {
                 callback(self.profile);
            }
        };

    }]).service('TechPointDashboardService',['$http', function ($http){
        var self = this;
        self.userSummaryData = null;

        self.queryUserSummaryData = function(callback){
            $http({
                method: 'GET',
                url: host + "student",
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
                callback('','err')
            });
        };
    }]).service('AuthService',['$http', function ($http) {
        var self = this;
        self.jwtToken = null;

        self.login = function(email,password,callback) {
            $http({
                method: 'POST',
                url: host + "auth/login",
                data: {
                    "email": email,
                    "password": password
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json"
                }
            }).then(function (data) {
                self.jwtToken = data.data['token'];
                callback(self.jwtToken);
            }, function errorCallback(response) {
                console.log('error occured: '+response);
                callback('','err')
            });
        };

        self.renderTokens = function(){
            $http({
                method: 'GET',
                url: host + "/getUser",
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + self.jwtToken
                }
            }).then(function (data) {
                setToken(data.role, "role");
                setToken(data.organization, "organization");
            }, function errorCallback(response) {
            
            });
            
        };

        self.logout = function (callback) {
            $http({
                method: 'POST',
                url: host + "auth/logout",
                data: {},
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function () {
                self.jwtToken = null;
                logout();
                callback();
            }, function errorCallback(response) {
                console.log('error occured: '+response);
                callback('err')
            });
        };

    }]).service('ResumeService',['$http', function ($http) {
        var self = this;
        self.jwtToken = null;
		self.uploadResume = function(id){
			var fd = new FormData();
			fd.append('file', document.getElementById("file").files[0]);
			$http.post(host + "student/resume/" + id, fd,{
                headers: {
					'Content-Type': undefined,
					'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            })
            .success(function () {
				console.log("Upload successful")
            }).error(function(response) {
                console.log('error occured: '+response);
                console.log('Here: '+getToken('auth'));
            });
        };
    }]);
})();
