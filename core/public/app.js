(function () {
    var app = angular.module('Xtern', ["ui.router", "angular-centered", "chart.js", "as.sortable", "DataManager", "ngSanitize", 'rzModule']);

    app.config(function ($stateProvider, $urlRouterProvider, $locationProvider) {
        // $locationProvider.html5Mode(true);
        $urlRouterProvider.when('/techpoint', '/techpoint/login');
        $urlRouterProvider.when('/techpoint/', '/techpoint/login');
        $urlRouterProvider.when('/company', '/company/login');
        $urlRouterProvider.when('/company/', '/company/login');
        $urlRouterProvider.otherwise("/techpoint/login");
        $stateProvider
            //Techpoint
            .state('techpoint', {
                url: "/techpoint",
                abstract: true,
                templateUrl: "public/techpoint/partials/techpoint.html",
                controller: 'TechPointMain'
            })
            .state('techpoint.dashboard', {
                url: "/dashboard",
                templateUrl: "public/modules/dashboard/partials/techpoint.missionControl.html",
                controller: 'TechPointDashboardCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInTechPoint($q);
                    }]
                }
            })
            .state('techpoint.accounts', {
                url: "/accounts",
                templateUrl: "public/modules/account_control/partials/accounts.html",
                controller: 'TechPointAccountCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInTechPoint($q);
                    }]
                }
            })
            .state('techpoint.process', {
                url: "/process_control",
                templateUrl: "public/modules/process_control/partials/process_control.html",
                controller: 'TechPointProcessControl',
                // abstract: true,
                // redirectTo:'techpoint.process.phase1',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInTechPoint($q);
                    }]
                }
            })           
            .state('techpoint.profile', {
                url: "/profile/:key",
                templateUrl: "public/modules/student_profile/partials/techpoint.studentProfile.html",
                controller: 'TechPointStudentProfileCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInTechPoint($q);
                    }]
                }
            })
            .state('techpoint.login', {
                url: "/login",
                templateUrl: "public/modules/login/partials/techpoint.login.html",
                controller: 'TechpointLogin',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedIn($q, 'ALREADY_AUTHENTICATED_TECHPOINT');
                    }]
                }
            })
            //Company
            .state('company', {
                url: "/company",
                abstract: true,
                templateUrl: "public/company/partials/company.html",
                controller: 'CompanyMain'
            })
            .state('company.dashboard', {
                url: "/dashboard",
                templateUrl: "public/modules/dashboard/partials/company.missionControl.html",
                //resolve: { authenticate: authenticate }
                controller: 'CompanyDashboardCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInCompany($q);
                    }]
                }
            })
            .state('company.recruting', {
                url: "/recruting",
                templateUrl: "public/company/partials/company.recruting.html",
                //resolve: { authenticate: authenticate }
                controller: 'CompanyRecruiting',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInCompany($q);
                    }]
                }
            })
            .state('company.profile', {
                url: "/profile/:key",
                templateUrl: "public/modules/student_profile/partials/company.studentProfile.html",
                //resolve: { authenticate: authenticate }
                controller: 'CompanyStudentProfileCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInCompany($q);
                    }]
                }
            })
            .state('company.login', {
                url: "/login",
                templateUrl: "public/modules/login/partials/company.login.html",
                controller: 'CompanyLogin',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedIn($q, 'ALREADY_AUTHENTICATED_COMPANY');
                    }]
                }
            });
    });
    app.run(function ($state, $rootScope) {
        $rootScope.$on('$stateChangeError', function (evt, toState, toParams, fromState, fromParams, error) {
            if (angular.isObject(error) && angular.isString(error.code)) {
                switch (error.code) {
                    case 'NOT_AUTHENTICATED_TECHPOINT':
                        // go to the login page
                        $state.go('techpoint.login');
                        break;
                    case 'ALREADY_AUTHENTICATED_TECHPOINT':
                        //go to the dash board
                        $state.go('techpoint.dashboard');
                        break;
                    case 'NOT_AUTHENTICATED_COMPANY':
                        // go to the login page
                        $state.go('company.login');
                        break;
                    case 'ALREADY_AUTHENTICATED_COMPANY':
                        //go to the dash board
                        $state.go('company.dashboard');
                        break;
                    case 'NOT_AUTHENTICATED_INSTRUCTOR':
                        // go to the login page
                        //$state.go('company.login');
                        break;
                    case 'ALREADY_AUTHENTICATED_INSTRUCTOR':
                        //go to the dash board
                        //$state.go('company.dashboard');
                        break;
                    default:
                        // set the error object on the error state and go there
                        $state.get('error').error = error;
                        $state.go('error');
                }
            }
            else {
                // unexpected error
                $state.go('techpoint.login');
            }
        })
    });
})();

//---------------------Classes and Function - to be moved later --------------//

var removeDataColors = function (data) {
    data.knownTech = [];
    for (var i in data.languages) {
        data.knownTech.push(data.languages[i].name);
    }
    //data.knownTech.sort();
};

// There should be a better way to do this, but I am blanking now -- maybe filter
// Corrects data formatting
var rowClass = function (data, key) {
    data.name = data.firstName + " " + data.lastName;
    data.namelink = '<a ui-sref="profile/' + key + '">' + data.name + "</a>";
    data.gradeLabel = data.grade;
    data.key = key;
    removeDataColors(data);

    //console.log(data);
    return data;
};

var removedDuplicates = function (arr) {
    return arr.filter(function (elem, index, self) {
        return index == self.indexOf(elem);
    });
};

var cleanStudents = function (student) {
    //student.interestedIn = removedDuplicates(student.interestedIn);
    //student.languages = removedDuplicates(student.interestedIn);
    return student;
};

