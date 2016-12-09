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
        self.studentKey = null;
        self.comments = null;
        self.commentKeys = null;

        self.getStudentData = function(key, callback){
            // console.log(id);
            if(!self.profile || self.studentKey != key) {
                $http({
                    method: 'GET',
                    url: host + "student/" + key,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    console.log('SUCCESS: get student data', data.data);
                    self.profile = cleanStudents(data.data);
                    self.studentKey = key;
                    callback(self.profile);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('', 'err');
                });
            } else {
                 callback(self.profile);
            }
        };

        self.getCommentData = function(callback) {
            // console.log(id);
            if(self.studentKey != key) {
                $http({
                    method: 'GET',
                    url: host + "comment",
                    data: {
                        "studentKey": self.studentKey
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    self.comments = data.data.comments;
                    self.commentKeys = data.data.keys;
                    callback(self.comments);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('', 'err');
                });
            } else {
                console.log('error occured: Called before student');
            }
        };

        // self.getStudentDataForIds = function(ids, callback){
        //     console.log("get student data for ids:");
        //     console.log(ids);
        //     // console.log(id);
        //         $http({
        //             method: 'POST',
        //             url: "http://localhost:8080/student/getstudents",
        //             data: {
        //                 "_ids": ids
        //             },
        //             headers: {
        //                 'Content-Type': "application/json",
        //                 'Accept': "application/json",
        //                 'Authorization': 'bearer ' + getToken('auth')
        //             }
        //         }).then(function (data) {
        //             console.log('get multiple students data:');
        //             console.log(data.data);
        //             callback(data.data);
        //         }, function errorCallback(response) {
        //             console.log('error occured: ');
        //             console.log(response);
        //             callback('', 'err');
        //         });
        // };

        self.addCommentToStudent = function(text, callback){
            $http({
                method: 'POST',
                url: host + "comment",
                data: {
                    "studentKey": self.studentKey,
                    "message": text
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

        self.removeCommentFromStudent = function(commentKey, callback){
            $http({
                method: 'DELETE',
                url: host + "comment/" + commentKey,
                data: {},
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

    }]).service('OrganizationService', ['$http', function ($http){
        var self = this;
        self.organization = null;
        self.organizationKey = null;

        self.getOrganizationData = function(key, callback){
            // console.log(id);
            if(!self.organization || self.organizationKey !== id) {
                $http({
                    method: 'GET',
                    url: host + "organization/" + id,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    console.log('get company data:');
                    console.log(data.data);
                    self.organization = data.data;
                    self.organizationKey = key;
                    callback(self.organization);
                }, function errorCallback(response) {
                    console.log('Company Services: error occured: ' + response);
                    console.log(response);
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

    }]).service('TechPointDashboardService',['$http', function ($http) {
        var self = this;
        self.studentSummaryData = null;
        self.studentKeys = null;

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
                self.studentSummaryData = data.data.students;
                self.studentKeys = data.data.keys;
                callback(self.studentSummaryData, self.studentKeys);
            }, function errorCallback(response) {
                console.log('error occured: '+response);
                console.log('Here: '+getToken('auth'));
                callback('','err');
            });
        };

    }]).service('AuthService',['$http', function ($http) {
        var self = this;
        self.userKey = null;
        
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
               // setToken(data.data.token);
               setToken(data.data['token'], "auth");
               setToken(data.data.organizationName, "organization");
               callback(data.data['token'],data.data.organizationName);
            }, function errorCallback(response) {
                console.log('error occured: ' + response);
                callback('','','err');
            });
        };

        // self.renderTokens = function (callback) {
        //     $http({
        //         method: 'GET',
        //         url: host + "admin/getUser",
        //         headers: {
        //             'Content-Type': "application/json",
        //             'Accept': "application/json",
        //             'Authorization': 'bearer ' + getToken('auth')
        //         }
        //     }).then(function (data) {
        //         setToken(data.data.organization, "organization");
        //         callback(data);
        //     }, function errorCallback(response) {
        //         callback('', response);
        //     });
        // };

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
                logout()
                callback();
            }, function errorCallback(response) {
                // console.log('error occured: '+response);
                callback('err')
            });
        };

    }]).service('ResumeService',['$http', function ($http) {
        var self = this;
        
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
