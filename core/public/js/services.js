(function (){
    var app = angular.module('DataManager',[]);
	var host = "http://localhost:8080/";
    app.service('ProfileService', ['$http', function ($http){
        var self = this;
        self.profile = null;
        self.studentKey = null;
        self.comments = null;
        self.commentKeys = null;

        self.getStudentData = function(key, callback){
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
                    // console.log('SUCCESS: get student data', data.data);
                    self.profile = cleanStudents(data.data);
                    self.studentKey = key;
                    callback(self.profile);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('', 'err');
                });
            } else {
                 callback(self.profile);
            }
        };

        self.getCommentData = function(callback) {
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
                console.log('error occured: ', response);
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
                console.log('error occured: ', response);
                callback('', 'err');
            });
        };

    }])
        .service('CompanyService', ['$http', function ($http){
        var self = this;
        self.organization = null;
        self.organizationKey = null;

        self.getOrganizationData = function(key, callback){
            if(!self.organization || self.organizationKey !== key) {
                $http({
                    method: 'GET',
                    url: host + "organization/" + key,
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
                    console.log('Company Services: error occured: ',  response);
                    callback('', 'err');
                });
            } else {
                 callback(self.organization);
            }
        };

        self.getOrganizationStudents = function (callback) {
            $http({
                method: 'GET',
                url: host + "organization/students",
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer ' + getToken('auth')
                }
            }).then(function (data) {
                //self.organization = data.data;
                //self.organizationKey = key;
                var students = data.data.students;
                for(var i = 0; i < students.length; i++) {
                    students[i].key = data.data.keys[i]
                }
                console.log(students);
                callback(students);
            }, function errorCallback(response) {
                console.log('Company Services: error occured: ' + response);
                callback('', 'err');
            });
        };

        self.addStudentToWishList = function (studentKey, callback) {
            $http({
                method: 'POST',
                url: host + "organization/addStudent",
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
                console.log('error occured: ', response );
                callback('', 'err');
            });
        };

        self.removeStudentFromWishList = function (studentKey, callback) {
            $http({
                method: 'POST',
                url: host + "organization/removeStudent",
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

        self.switchStudentsInWishList = function (studentKey, pos, callback) {
            $http({
                method: 'PUT',
                url: host + "organization/moveStudent",
                data: {
                    "studentKey": studentKey,
                    "position": pos
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
                console.log('error occured: ', response);
                callback('','err');
            });
        };
    }])
    .service('ReviewerDashboardService',['$http', function ($http) {
        var self = this;
        self.studentSummaryData = null;
        self.studentKeys = null;

        self.queryReviewGroup = function(callback){
            $http({
                method: 'POST',
                url: host + "reviewer/getReviewGroupForReviewer",
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function (data) {
                self.studentSummaryData = data.data.students;
                self.studentKeys = data.data.users.students;
                self.studentGrades = data.data.studentGrades;
                callback(self.studentSummaryData, self.studentKeys, self.studentGrades);
            }, function errorCallback(response) {
                console.log('error occured: ', response);
                callback('','err');
            });
        };
    }])
    .service('ReviewerProfileService',['$http', function ($http) {
        var self = this; 
        self.getReviewerGradeForStudent = function(studentKey, callback){
            $http({
                method: 'GET',
                url: host + "reviewer/getReviewerGradeForStudent/" + studentKey,
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function (data) {
                self.reviewerGrade = data.data.grade;
                callback(data.data.grade);
            }, function errorCallback(response) {
                console.log('error occured: ', response);
                callback('','err');
            });
        };

        self.postReviewerGradeForStudent = function(studentKey, reviewerGrade){
            $http({
                method: 'POST',
                url: host + "reviewer/postReviewerGradeForStudent",
                data: {
                    "studentKey": studentKey,
                    "reviewerGrade": reviewerGrade
                },
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function (data) {
            }, function errorCallback(response) {
                console.log('error occured: ', response);
            });
        };
    }])
        .service('AccountControlService',['$http', function ($http){
            var self = this;
            self.userData = null;
            self.getOrganizations = function(callback) {
                var route = "organization";
                $http({
                    method: 'GET',
                    url: host + route,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer '+ getToken('auth')
                    }
                }).then(function (data) {
                    for(var i = 0; i < data.data.organizations.length; i++) {
                        data.data.organizations[i].key = data.data.keys[i]
                    }
                    callback(data.data.organizations);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('','err')
                });
            };
            self.getUsers = function(orgKey,callback) {
                // console.log('Key: ', orgKey);
                var route = "user/org/"+orgKey;
                $http({
                    method: 'GET',
                    url: host + route,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    if(data.data && data.data.users && data.data.keys){
                        for(var i = 0; i < data.data.users.length; i++) {
                            data.data.users[i].key = data.data.keys[i];
                        }
                        callback(data.data.users);
                    }else{
                        callback([]);
                    }
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    callback('','err')
                });
            };
            self.addUser = function(user, callback){
                // console.log('Here: ', user);
                var route = "user";
                $http({
                    method: 'POST',
                    url: host + route,
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
                    // console.log('Here: ' + getToken('auth'));
                    callback('', 'err')
                });
            };
            self.updateUser = function(user, callback){
                var route = "user/"+user.key;
                $http({
                    method: 'PUT',
                    url: host + route,
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
                }, function errorCallback(response) {
                    console.log('error occured: ' ,  response);
                    callback('', 'err')
                });
            };
            self.deleteUser = function (key, callback) {
                var route = "user/" + key;
                $http({
                    method: 'DELETE',
                    url: host + route,
                    headers: {
                        'Content-Type': "application/json",
                        'Accept': "application/json",
                        'Authorization': 'bearer ' + getToken('auth')
                    }
                }).then(function (data) {
                    callback(data);
                }, function errorCallback(response) {
                    console.log('error occured: ' ,  response);
                    callback('', 'err')
                });
            };
    }])
        .service('AuthService',['$http', function ($http) {
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
                    console.log('error occured: ', response);
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
                // $http({
                //     method: 'POST',
                //     url: host + "auth/logout",
                //     data: {},
                //     headers: {
                //         'Content-Type': "application/json",
                //         'Accept': "application/json",
                //         'Authorization': 'bearer '+getToken('auth')
                //     }
                // }).then(function () {
                //     logout();
                //     callback();
                // }, function errorCallback(response) {
                //     // console.log('error occured: '+response);
                //     callback('err')
                // });
                logout();
                callback();
            };

    }])
        .service('ResumeService',['$http', function ($http) {
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
                console.log('error occured: ', response);
                console.log('Here: '+getToken('auth'));
            });
        };
    }])
    .service('TechPointReviewerControlService',['$http', function ($http) {
        var self = this;
        self.reviewGroups = null;
        self.reviewGroupKeys = null;

        self.createReviewGroups = function(minStudents, minReviewers, callback){
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
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function (data) {
                callback(data);
            }, function errorCallback(response) {
                console.log('error occured: ', response);
            });
        };

        self.queryReviewGroups = function(callback){
            $http({
                method: 'GET',
                url: "reviewer/getReviewGroups",
                headers: {
                    'Content-Type': "application/json",
                    'Accept': "application/json",
                    'Authorization': 'bearer '+getToken('auth')
                }
            }).then(function (data) {
                self.reviewGroups = data.data.users;
                self.reviewGroupKeys = data.data.keys;
                callback(self.reviewGroups, self.reviewGroupKeys);
            }, function errorCallback(response) {
                console.log('error occured: ', response);
                callback('','err');
            });
        };
    }]);
})();
