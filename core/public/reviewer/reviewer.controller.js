angular.module('Xtern')
    .controller('ReviewerMain', ['$scope', '$rootScope', '$state', 'AuthService','CompanyService', 'ProfileService', function ($scope, $rootScope, $state, AuthService, CompanyService, ProfileService) {
        var self = this;
        $scope.loggedIn = !!getToken("organization");

        $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams, options) {
                $scope.loggedIn = !!getToken("organization");
            });

        $scope.logout = function () {
            AuthService.logout(function (err) {
                if (err) {
                    console.log('Logout unsuccessful');
                } else {
                    $scope.loggedIn = false;
                    $state.go('reviewer.login');
                }
            });
        };
    }]);
