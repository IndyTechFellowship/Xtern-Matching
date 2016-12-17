angular.module('Xtern')
    .controller('CompanyMain', ['$scope', '$rootScope', '$state', 'AuthService','CompanyService', 'ProfileService', function ($scope, $rootScope, $state, AuthService, CompanyService, ProfileService) {
        var self = this;
        $scope.loggedIn = !!getToken("organization");
        // CompanyService.getCurrentCompany(function(company) {
        $scope.companyData = getToken("organization");
        // });

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
    }])
    .controller('CompanyRecruiting', ['$scope', '$state', 'ProfileService', 'CompanyService', function ($scope, $state, ProfileService, CompanyService) {
        $scope.recruitmentList = [];

        $scope.companyData = getToken("organization");
        CompanyService.getOrganizationStudents(function(data) {
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
                console.log(obj.source.index+' '+obj.dest.index);
                CompanyService.switchStudentsInWishList($scope.recruitmentList[obj.source.index].key, obj.dest.index, function(data) {
                    console.log("order changed: "+data);
                });
            }
        };
    }]);
