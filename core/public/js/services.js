(function (){
    let app = angular.module('DataManager',[]);
	let host = location.host;
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
                    let student = cleanStudents(data.data);
                    student.key = key;
                    callback(student);
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
                    let comments = data.data.comments;
                    for(let i = 0;i < comments.length;i++) {
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
                    let comment = data.data.comment;
                    comment.key = data.data.key;
                    callback(comment);
                }, function errorCallback(err) {
                    callback(null, 'err');
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
            let self = this;
            self.organization = null;
            self.organizationKey = null;

            self.getOrganizationData = function(key, callback){
                if(!self.organization || self.organizationKey !== key) {
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
                        //console.log(data.data);
                        self.organization = data.data;
                        self.organizationKey = key;
                        callback(self.organization);
                    }, function errorCallback(err) {
                        callback(null, err);
                    });
                } else {
                     callback(self.organization);
                }
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
                    //self.organization = data.data;
                    //self.organizationKey = key;
                    var students = data.data.students;
                    for(var i = 0; i < students.length; i++) {
                        students[i].key = data.data.keys[i]
                    }
                    //console.log(students);
                    callback(students);
                }, function errorCallback(response) {
                    //console.log('Company Services: error occured: ' + response);
                    callback('', 'err');
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
                    console.log('error occured: ', response );
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

            self.switchStudentsInWishList = function (studentKey, pos, callback) {
                $http({
                    method: 'PUT',
                    url: "organization/moveStudent",
                    host: host,
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
            let self = this;
            self.studentSummaryData = null;
            self.studentKeys = null;

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
                    self.studentSummaryData = data.data.students;
                    self.studentKeys = data.data.keys;
                    callback(self.studentSummaryData, self.studentKeys);
                }, function errorCallback(response) {
                    console.log('error occured: ', response);
                    console.log('Here: '+getToken('auth'));
                    callback('','err');
                });
            };
    }])
        .service('AccountControlService',['$http', function ($http){
            let self = this;
            self.userData = null;
            self.getOrganizations = function(callback) {
                var route = "organization";
                $http({
                    method: 'GET',
                    url: route,
                    host: host,
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
                console.log('Key: ', orgKey);
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
                console.log('Here: ', user);
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
                    // console.log('Here: ' + getToken('auth'));
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
                }, function errorCallback(response) {
                    console.log('error occured: ' ,  response);
                    callback('', 'err')
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
                    console.log('error occured: ' ,  response);
                    callback('', 'err')
                });
            };
    }])
        .service('AuthService',['$http', function ($http) {
            let self = this;
            self.userKey = null;

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
                   callback(data.data['token'],data.data.organizationName);
                }, function errorCallback(response) {
                    console.log('error occured: ' + response);
                    callback('','','err');
                });
            };

            self.logout = function (callback) {
                logout();
                callback();
            };

    }])
        .service('ResumeService',['$http', function ($http) {
            let self = this;
        
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
    }]);
})();
