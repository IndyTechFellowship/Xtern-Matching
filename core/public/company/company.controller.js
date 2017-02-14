'use strict';
angular.module('Xtern')
    .controller('CompanyMain', ['$scope', '$rootScope', '$state', 'AuthService', 'CompanyService', function ($scope, $rootScope, $state, AuthService, CompanyService) {
        $scope.loggedIn = !!getToken("organization");
        $scope.companyData = getToken("organization");

        $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams, options) {
                $scope.loggedIn = !!getToken("organization");
                $scope.companyData = getToken("organization");
                if (toState.name == "company.profile") {
                    $('#profile').show();
                }
                else {
                    $('#profile').hide();
                }
            });

        $scope.addStudentToCompany = function (studentID) {
            CompanyService.addStudentToWishList(studentID, function (data) {
                console.log(data);
            });
        };

        $scope.logout = function () {
            AuthService.logout(function (err) {
                if (err) {
                    console.log('Logout unsuccessful');
                } else {
                    $scope.loggedIn = false;
                    $state.go('company.login');
                }
            });
        };
    }]);