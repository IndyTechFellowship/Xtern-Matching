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
            // console.log(id);
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
                    callback('', 'err');
                });
            } else {
                 callback(self.profile);
            }
        };

        self.getStudentDataForIds = function(ids, callback){
            console.log("get student data for ids:");
            console.log(ids);
            // console.log(id);
                $http({
                    method: 'POST',
                    url: "http://localhost:8080/student/getstudents",
                    data: {
                        "_ids": ids
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    console.log('get multiple students data:');
                    console.log(data.data);
                    callback(data.data);
                }, function errorCallback(response) {
                    console.log('error occured: ');
                    console.log(response);
                    callback('', 'err');
                });
        };

        self.addCommentToStudent = function(student_id, author_name, group_name, text, callback){
            $http({
                method: 'POST',
                // TODO: replace this when comments are redone
                url: host + "student/addComment",
                data: {
                    "id": student_id,
                    "author_name": author_name,
                    "group_name": group_name,
                    "text": text
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function (data) {
                callback(data);
            }, function errorCallback(response) {
                console.log('error occured: ' );
                console.log(response);
                callback('', 'err');
            });
        };

        self.removeCommentFromStudent = function(student_id, author_name, group_name, text, callback){
            $http({
                method: 'POST',
                // TODO: replace this when comments are redone
                url: host + "student/deleteComment",
                data: {
                    "id": student_id,
                    "author_name": author_name,
                    "group_name": group_name,
                    "text": text
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function (data) {
                callback(data);
            }, function errorCallback(response) {
                console.log('error occured: ' );
                console.log(response);
                callback('', 'err');
            });
        };

    }]).service('CompanyService', ['$http', function ($http){
        var self = this;
        self.company = null;

        self.getCurrentCompany = function(callback){
            // console.log(id);
            // if(!self.company || self.company._id !== id) {
                $http({
                    method: 'GET',
                    url: host + "company/getCurrentCompany/0",
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    // console.log('get current company data:');
                    // console.log(data.data);
                    self.company = data.data;
                    callback(self.company);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('', 'err');
                });
            // } else {
            //      callback(self.company);
            // }
        };

        self.getCompanyDataForId = function(id, callback){
            // console.log(id);
            if(!self.company || self.company._id !== id) {
                $http({
                    method: 'GET',
                    url: host + "company/" + id,
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

        self.addStudentToWishList = function (studentId, callback) {
            $http({
                method: 'POST',
                // TODO: replace this id when company login is done
                url: host + "company/addStudent",
                // url: "http://localhost:8080/company/" + id,
                data: {
                    // "id": 5066549580791808,
                    "token": getToken('auth'),
                    "studentId": studentId
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function (data) {
                callback(data);
            }, function errorCallback(response) {
                console.log('error occured: ' );
                console.log(response);
                callback('', 'err');
            });
        };

        self.removeStudentFromWishList = function (studentId, callback) {
            $http({
                method: 'POST',
                // TODO: replace this id when company login is done
                url: host + "company/removeStudent",
                data: {
                    // "id": 5066549580791808,
                    "studentId": studentId
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function (data) {
                callback(data);
            }, function errorCallback(response) {
                console.log('error occured: ');
                console.log(response);
                callback('', 'err');
            });
        };

        self.switchStudentsInWishList = function (student1Id, student2Id, callback) {
            $http({
                method: 'POST',
                // TODO: replace this id when company login is done
                url: host + "company/switchStudents",
                data: {
                    "id": 5066549580791808,
                    "student1Id": student1Id,
                    "student2Id": student2Id
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function (data) {
                callback(data);
            }, function errorCallback(response) {
                console.log('error occured: ');
                console.log(response);
                callback('', 'err');
            });
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
                    "email":"xniccum@gmail.com",
                    "password": "admin1"
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
                callback('','err');
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
