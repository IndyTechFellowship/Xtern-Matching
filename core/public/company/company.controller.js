angular.module('Xtern')
.controller('CompanyMain', ['$scope', '$rootScope', '$state', 'AuthService','CompanyService', 'ProfileService', function ($scope, $rootScope, $state, AuthService, CompanyService, ProfileService) {
    var self = this;
    $scope.loggedIn = !!getToken("organization");

    $rootScope.$on('$stateChangeStart',
        function (event, toState, toParams, fromState, fromParams, options) {
            $scope.loggedIn = !!getToken("organization");
            if (toState.name == "company.profile") {
                $('#profile').show();
            }
            else {
                $('#profile').hide();
            }
            CompanyService.getOrganizationCurrentFromLogin(function(company) {
                console.log("company data in recruiting controller:", company);
                // ProfileService.getStudentDataForIds(company.studentIds, function(data) {
                //     $scope.recruitmentList = data;
                // });
            });

        });

    CompanyService.getOrganizationCurrentFromLogin(function(company) {
        console.log("company data in recruiting controller:", company);
        // ProfileService.getStudentDataForIds(company.studentIds, function(data) {
        //     $scope.recruitmentList = data;
        // });
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
}])
.controller('CompanyRecruiting', ['$scope', '$rootScope', '$state', 'ProfileService', 'CompanyService', function ($rootScope, $scope, $state, ProfileService, CompanyService) {
    $scope.recruitmentList = [];

    CompanyService.getOrganizationStudents(function(data) {
        console.log("ORG Students: ", data);
        $scope.recruitmentList = data;
    });

    $scope.sortableOptions = {
        containment: '#table-container',
        containerPositioning: 'relative'
    };

    $scope.removeRecruit = function (key) {
        console.log("remove recruit: "+key);
        CompanyService.removeStudentFromWishList(key, function(data) {
            for (var i = $scope.recruitmentList.length - 1; i >= 0; i--) {
                if ($scope.recruitmentList[i].key == key) {
                    $scope.recruitmentList.splice(i, 1);
                }
            }
        });

    };

    $scope.viewRecruit = function (key) {
        $state.go('company.profile', { key: key });
    };

    $scope.addStudent = function (key) {
        console.log("add student:");
        console.log(key);
    };

    $scope.dragControlListeners = {
        orderChanged: function(obj) {
            console.log("new index", obj.dest.index);
            console.log(obj.source.index+' '+obj.dest.index);
            CompanyService.switchStudentsInWishList($scope.recruitmentList[obj.source.index].key, $scope.recruitmentList[obj.dest.index].key, function(data) {
                console.log("order changed: ", data);
            });
        }
    };
}]);
