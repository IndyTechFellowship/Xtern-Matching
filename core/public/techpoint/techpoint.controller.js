'use strict';
angular.module('Xtern')
    .controller('TechPointMain', ['$scope', '$rootScope', '$state', 'AuthService', function($scope, $rootScope, $state, AuthService){
        $scope.loggedIn = !!getToken("organization");
        //$scope.loggedIn = isLoggedInTechPoint();
        //$scope.isCompany = false;

       $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams, options) {
                $scope.loggedIn = !!getToken("organization");
                if (toState.name == "techpoint.profile") {
                    $('#profile').show();
                }
                else {
                    $('#profile').hide();
                }
            });

        $scope.logout = function () {
            AuthService.logout(function (err) {
                if(err) {
                    console.log('Logout unsuccessful');
                } else {
                    $scope.loggedIn = false;
                    $state.go('techpoint.login');
                }
            });
        };
    }]);


