/**
 * Created by Steven on 3/28/2016.
 */
 (function (){
    var app = angular.module('DataManager',[]);
	//var host = "http://xtern-matching.appspot.com/"
	//var host = "http://xtern-matching-143216.appspot.com/" //DEV Server
	var host = "http://localhost:8080/"
    app.service('ProfileService', ['$http', function ($http){
        var self = this;
        self.profile = null;

        self.getStudentDataForId = function(id, callback){
            console.log(id);
            if(!self.profile || self.profile._id !== id) {
                $http({
                    method: 'GET',
<<<<<<< HEAD
                    url: "http://xtern-matching.appspot.com/student/" + id,
=======
                    url: host + "student/" +id,
>>>>>>> master
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
<<<<<<< HEAD
                url: "http://xtern-matching.appspot.com/student",
=======
                url: host + "student",
>>>>>>> master
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
                callback('','err')
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
<<<<<<< HEAD
            $http.post("http://xtern-matching.appspot.com/auth/login",{"email":email, "password": password}).then(function(data) {
=======
            $http.post(host + "auth/login",{"email":email, "password": password}).then(function(data) {
>>>>>>> master
                self.jwtToken = data.data['token'];
                //console.log('Here: '+self.jwtToken);
                callback(self.jwtToken);
            }, function errorCallback(response) {
                console.log('error occured: '+response);
                callback('','err')
            });
        }
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
                },
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
