'use strict';

(function () {
    //Node Modules
    var angular = require('angular');
    require('angular-ui-router');
    require('angular-sanitize');
    require('angular-centered');
    require('chart.js');
    require('angular-chart.js');
    require('ng-sortable');
    require('pdfjs-dist/build/pdf.worker.js');
    require('pdfjs-dist/build/pdf.js');
    require('pdfjs-dist/build/pdf.combined.js');
    require('pdfjs-dist/build/pdf.worker.entry.js');


    var app = angular.module('Xtern', ["ui.router", "angular-centered", "chart.js", "as.sortable","ngSanitize","DataManager"]);//ngSanitize "ui.router", "angular-centered", "chart.js", "as.sortable","ngSanitize"

    require('./public/techpoint/techpoint.controller.js');
    require('./public/company/company.controller.js');
    require('./public/reviewer/reviewer.controller.js');
    require('./public/modules/');


    app.config(function ($stateProvider, $urlRouterProvider, $locationProvider) {
        // $locationProvider.html5Mode(true);
        $urlRouterProvider.when('/techpoint', '/techpoint/login');
        $urlRouterProvider.when('/techpoint/', '/techpoint/login');
        $urlRouterProvider.when('/company', '/company/login');
        $urlRouterProvider.when('/company/', '/company/login');
        $urlRouterProvider.when('/reviewer', '/reviewer/login');
        $urlRouterProvider.when('/reviewer/', '/reviewer/login');
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
            .state('techpoint.reviewerControls', {
                url: "/reviewerControls",
                templateUrl: "public/modules/reviewer_controls/partials/reviewerControls.html",
                controller: 'TechPointReviewerCtrl',
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
                templateUrl: "public/modules/recruting/partials/company.recruting.html",
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
            })
            // Reviewer
            .state('reviewer', {
                url: "/reviewer",
                abstract: true,
                templateUrl: "public/reviewer/partials/reviewer.html",
                controller: 'ReviewerMain'
            })
            .state('reviewer.login', {
                url: "/login",
                templateUrl: "public/modules/login/partials/reviewer.login.html",
                controller: 'ReviewerLogin',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedIn($q, 'ALREADY_AUTHENTICATED_REVIEWER');
                    }]
                }
            })
            .state('reviewer.profile', {
                url: "/profile/:key",
                templateUrl: "public/modules/student_profile/partials/reviewer.studentProfile.html",
                controller: 'ReviewerStudentProfileCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInReviewer($q);
                    }]
                }
            })
            .state('reviewer.dashboard', {
                url: "/dashboard",
                templateUrl: "public/modules/dashboard/partials/reviewer.dashboard.html",
                controller: 'ReviewerDashboardCtrl',
                resolve: {
                    security: ['$q', function ($q) {
                        return isLoggedInReviewer($q, 'ALREADY_AUTHENTICATED_REVIEWER');
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
                    case 'NOT_AUTHENTICATED_REVIEWER':
                        // go to the login page
                        $state.go('reviewer.login');
                        break;
                    case 'ALREADY_AUTHENTICATED_REVIEWER':
                        //go to the dash board
                        $state.go('reviewer.dashboard');
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
        });
    });
})();



