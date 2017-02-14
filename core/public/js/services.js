'use strict';
    var app = angular.module('DataManager',[]);
	var host = "http://localhost:8080";
    app
        .service('ProfileService', ['$http', function ($http){
            this.getStudent = function(key, callback){
                $http({
                    method: 'GET',
                    url: "student/" + key,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var student = data.data;
                    student.key = key;
                    callback(student);
                }, function errorCallback(err) {
                    callback(null, err);
                });
            };

            this.getStudentDataForIds = function(keys, callback){
                $http({
                    method: 'GET',
                    url: "student/" + keys,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var students = data.data;
                    callback(students);
                }, function errorCallback(err) {
                    callback(null, err);
                });
            };

            this.setStatus = function (key, status, callback) {
                $http({
                    method: 'PUT',
                    url: "student/" + key + "/status",
                    host: host,
                    data: {
                        "status": status
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function () {
                    callback();
                }, function errorCallback(err) {
                    callback(err);
                });
            };

            this.setR1Grade = function (key, grade, callback) {
                $http({
                    method: 'PUT',
                    url: "student/" + key + "/grade",
                    host: host,
                    data: {
                        "grade": grade
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function () {
                    callback();
                }, function errorCallback(err) {
                    callback(err);
                });
            };

            this.getComments = function(key,callback) {
                $http({
                    method: 'GET',
                    url: "comment/" + key,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var comments = data.data.comments;
                    for(var i = 0;i < comments.length;i++) {
                        comments[i].key = data.data.keys[i];
                    }
                    callback(comments);
                }, function errorCallback(err) {
                    callback(null, err);
                });
            };

            this.addComment = function(key ,text, callback){
                $http({
                    method: 'POST',
                    url: "comment/" + key,
                    host: host,
                    data: {
                        "message": text
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var comment = data.data.comment;
                    comment.key = data.data.key;
                    callback(comment);
                }, function errorCallback(err) {
                    callback(null, 'err');
                });
            };
            this.editComment = function(key ,text, callback) {
                $http({
                    method: 'PUT',
                    url: "comment/" + key,
                    host: host,
                    data: {
                        "message": text
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var comment = data.data.comment;
                    comment.key = data.data.key;
                    callback(comment);
                }, function errorCallback(err) {
                    callback(null,err);
                });
            };
            this.removeComment = function(commentKey, callback){
            $http({
                method: 'DELETE',
                url: "comment/" + commentKey,
                host: host,
                headers: {
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function () {
                callback();
            }, function errorCallback(err) {
                callback(err);
            });
        };
    }])
        .service('CompanyService', ['$http', function ($http){
            var self = this;
            self.getOrganizationData = function(key, callback){
                $http({
                    method: 'GET',
                    url: "organization/" + key,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var organization = data.data;
                    callback(organization);
                }, function errorCallback(err) {
                    callback(null, err);
                });
            };

            self.getOrganizationCurrentFromLogin = function(callback){
                $http({
                    method: 'GET',
                    url: "organization/current",
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var organization = data.data;
                    callback(organization);
                }, function errorCallback(err) {
                    callback(null, err);
                });
            };

            self.getOrganizationStudents = function (callback) {
                $http({
                    method: 'GET',
                    url: "organization/students",
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    var students = data.data.students;
                    for(var i = 0; i < students.length; i++) {
                        students[i].key = data.data.keys[i]
                    }
                    callback(students);
                }, function errorCallback(response) {
                    callback('', response);
                });
            };

            self.addStudentToWishList = function (studentKey, callback) {
                $http({
                    method: 'POST',
                    url: "organization/addStudent",
                    host: host,
                    data: {
                        "studentKey": studentKey
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            };

            self.removeStudentFromWishList = function (studentKey, callback) {
                $http({
                    method: 'POST',
                    url: "organization/removeStudent",
                    host: host,
                    data: {
                        "studentKey": studentKey
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            };

            self.switchStudentsInWishList = function (studentKey1, studentKey2, callback) {
                $http({
                    method: 'PUT',
                    url: "organization/switchStudents",
                    host: host,
                    data: {
                        "studentKey1": studentKey1,
                        "studentKey2": studentKey2
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            };
        }])
        .service('TechPointDashboardService',['$http', function ($http) {
            var self = this;
            self.queryUserSummaryData = function(callback){
                $http({
                    method: 'GET',
                    url: "student",
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer '+getToken('auth')
                    }
                }).then(function (data) {
                    callback(data.data.students, data.data.keys);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('','err');
                });
            };
    }])
        .service('ReviewerDashboardService',['$http', function ($http) {
            var self = this;

            self.queryReviewGroup = function (callback) {
                $http({
                    method: 'POST',
                    url: "reviewer/getReviewGroupForReviewer",
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data.data.students, data.data.users.students, data.data.studentGrades);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            };
        }])
        .service('ReviewerProfileService', ['$http', function ($http) {
            var self = this;
            self.getReviewerGradeForStudent = function (studentKey, callback) {
                $http({
                    method: 'GET',
                    url: "reviewer/getReviewerGradeForStudent/" + studentKey,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    self.reviewerGrade = data.data.grade;
                    callback(data.data.grade);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            };

            self.postReviewerGradeForStudent = function (studentKey, reviewerGrade) {
                $http({
                    method: 'POST',
                    url: "reviewer/postReviewerGradeForStudent",
                    data: {
                        "studentKey": studentKey,
                        "reviewerGrade": reviewerGrade
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                }, function errorCallback(response) {
                    console.log('error occurred: ', response);
                });
            };
        }])
        .service('AccountControlService',['$http', function ($http){
            var self = this;
            self.getOrganizations = function(callback) {
                var route = "organization";
                $http({
                    method: 'GET',
                    url: route,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    for (var i = 0; i < data.data.organizations.length; i++) {
                        data.data.organizations[i].key = data.data.keys[i]
                    }
                    callback(data.data.organizations);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err')
                });
            };
            self.getUsers = function(orgKey,callback) {
                var route = "user/org/"+orgKey;
                $http({
                    method: 'GET',
                    url: route,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    if(data.data.users) {
                        for (var i = 0; i < data.data.users.length; i++) {
                            data.data.users[i].key = data.data.keys[i];
                        }
                    }
                    callback(data.data.users);
                }, function errorCallback(err) {
                    callback(null,err)
                });
            };
            self.addUser = function(user, callback){
                var route = "user";
                $http({
                    method: 'POST',
                    url: route,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    },
                    data: {
                        name: user.name,
                        email: user.email,
                        //TODO encrypt into nonplain-text
                        password: user.password,
                        orgKey: user.organization
                    }
                }).then(function (data) {
                    //success
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err')
                });
            };
            self.addCompany = function(company, callback){
                var route = "organization";
                $http({
                    method: 'POST',
                    url: route,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    },
                    data: {
                        name: company
                    }
                }).then(function (data) {
                    //success
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occurred: ', response);
                    callback('', 'err')
                });
            };
            self.updateUser = function(user, callback){
                var route = "user/"+user.key;
                $http({
                    method: 'PUT',
                    url: route,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    },
                    data: {
                        name: user.name,
                        email: user.email,
                        password: user.password
                    }
                }).then(function (data) {
                    //success
                    callback(data);
                }, function errorCallback(err) {
                    callback(null, err)
                });
            };
            self.deleteUser = function (key, callback) {
                var route = "user/" + key;
                $http({
                    method: 'DELETE',
                    url: route,
                    host: host,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err')
                });
            };
    }])
        .service('AuthService',['$http', function ($http) {
            var self = this;
            self.login = function(email,password,callback) {
                $http({
                    method: 'POST',
                    url: "auth/login",
                    host: host,
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
                   sessionStorage.setItem("userKey", data.data.userKey);
                   callback(data.data['token'],data.data.organizationName);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', '', 'err');
                });
            };
            self.logout = function (callback) {
                logout();
                callback();
            };
        }])
        .service('TechPointReviewerControlService', ['$http', function ($http) {
            var self = this;
            self.reviewGroups = null;
            self.reviewGroupKeys = null;

            self.createReviewGroups = function (minStudents, minReviewers, callback) {
                $http({
                    method: 'POST',
                    url: "reviewer/create",
                    data: {
                        "minStudents": minStudents,
                        "minReviewers": minReviewers
                    },
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                });
            };

            self.queryReviewGroups = function (callback) {
                $http({
                    method: 'GET',
                    url: "reviewer/getReviewGroups",
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    self.reviewGroups = data.data.users;
                    self.reviewGroupKeys = data.data.keys;
                    callback(self.reviewGroups, self.reviewGroupKeys);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            };
        }]);