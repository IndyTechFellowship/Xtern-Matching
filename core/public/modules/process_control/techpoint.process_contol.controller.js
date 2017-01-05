angular.module('Xtern')
    .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;
        $scope.companyList = [];

        var setup = function () {
            AccountControlService.getOrganizations(function (organizations) {
                $scope.companyList = organizations;                               
            });
        };


        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);