(function () {

    var app = angular.module('Xtern', ["ui.router","angular-centered","chart.js","as.sortable","DataManager","ngSanitize"]);//ngSanitize

    app.config(function($stateProvider, $urlRouterProvider) {
        $urlRouterProvider.when('/techpoint', '/techpoint/login');
        $urlRouterProvider.when('/techpoint/', '/techpoint/login');
        $urlRouterProvider.when('/company', '/company/login');
        $urlRouterProvider.when('/company/', '/company/login');
        $urlRouterProvider.otherwise("/techpoint/login");
        $stateProvider
            //Techpoint
            .state('techpoint', {
                url:"/techpoint",
                abstract: true,
                templateUrl: "public/partials/techpoint/techpoint.html",
                controller: 'TechPointMain'
            })
            .state('techpoint.dashboard',{
                url:"/dashboard",
                templateUrl: "public/partials/techpoint/techpoint.missionControl.html",
                //resolve: { authenticate: authenticate }
                controller: 'TechPointDashboardCtrl',
                resolve: {
                    security: ['$q', function($q){
                      //  console.log($q, status);
                        if(!isLoggedIn('techPoint')){
                            var errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
                            return $q.reject(errorObject);
                        }
                    }]
                }
            })
            .state('techpoint.profile',{
                url:"/profile/:_id",
                templateUrl: "public/partials/studentProfile.html",
                //resolve: { authenticate: authenticate }
                controller: 'StudentProfileCtrl',
                resolve: {
                    security: ['$q', function($q){
                        //  console.log($q, status);
                        if(!isLoggedIn('techPoint')){
                            var errorObject = { code: 'NOT_AUTHENTICATED_TECHPOINT' };
                            return $q.reject(errorObject);
                        }
                    }]
                }
            })
            .state('techpoint.login', {
                url:"/login",
                templateUrl: "public/partials/techpoint/techpoint.login.html",
                controller: 'TechpointLogin',
                resolve: {
                    security: ['$q', function ($q) {
                        //  console.log($q, status);
                        if (isLoggedIn('techPoint')) {
                            console.log(isLoggedIn('techPoint'));
                       // if(false){
                            var errorObject = {code: 'ALREADY_AUTHENTICATED_TECHPOINT'};
                            return $q.reject(errorObject);
                        }
                    }]
                }
            })
            //Company
            .state('company', {
                url:"/company",
                abstract: true,
                templateUrl: "public/partials/company/company.html",
                controller: 'CompanyMain'
            })
            .state('company.dashboard',{
                url:"/dashboard",
                templateUrl: "public/partials/company/company.missionControl.html",
                //resolve: { authenticate: authenticate }
                controller: 'CompanyDashboardCtrl',
                resolve: {
                    security: ['$q', function($q){
                        //  console.log($q, status);
                        if(!isLoggedIn('company')){
                            var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
                            return $q.reject(errorObject);
                        }
                    }]
                }
            }).state('company.recruting',{
                url:"/recruting",
                templateUrl: "public/partials/company/company.recruting.html",
                //resolve: { authenticate: authenticate }
                controller: 'CompanyRecruiting',
                resolve: {
                    security: ['$q', function($q){
                        //  console.log($q, status);
                        if(!isLoggedIn('company')){
                            var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
                            return $q.reject(errorObject);
                        }
                    }]
                }
            })
            .state('company.profile',{
                url:"/profile/:_id",
                templateUrl: "public/partials/studentProfile.html",
                //resolve: { authenticate: authenticate }
                controller: 'StudentProfileCtrl',
                resolve: {
                    security: ['$q', function($q){
                        //  console.log($q, status);
                        if(!isLoggedIn('company')){
                            var errorObject = { code: 'NOT_AUTHENTICATED_COMPANY' };
                            return $q.reject(errorObject);
                        }
                    }]
                }
            })
            .state('company.login', {
                url:"/login",
                templateUrl: "public/partials/company/company.login.html",
                controller: 'CompanyLogin',
                resolve: {
                    security: ['$q', function ($q) {
                        //  console.log($q, status);
                        if (isLoggedIn('company')) {
                            console.log(isLoggedIn('company'));
                            // if(false){
                            var errorObject = {code: 'ALREADY_AUTHENTICATED_COMPANY'};
                            return $q.reject(errorObject);
                        }
                    }]
                }
            })
			.state('student-profile', {
				url: "/student-profile",
				templateUrl: "public/partials/studentProfile.html"
			});
	});
    app.run(function($state, $rootScope){
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

var removeDataColors = function(data){
    data.knownTech = [];
    for(var i in data.languages){
        data.knownTech.push(data.languages[i].name);
    }
    //data.knownTech.sort();
};

// There should be a better way to do this, but I am blanking now -- maybe filter
// Corrects data formatting
var rowClass = function(data){
    data.name = data.firstName + " " + data.lastName;
    //data.gradeLabel = data.r1Grade.text;
    //data.gradeValue = data.r1Grade.value;
    data.namelink = '<a ui-sref="profile/' + data._id + '">' + data.name + "</a>";
    removeDataColors(data);
    
    //console.log(data);
    return data;
};