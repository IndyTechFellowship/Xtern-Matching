angular.module('Xtern')
    .controller('CompanyMain', ['$scope', '$rootScope', '$state', 'AuthService','CompanyService', 'ProfileService', function ($scope, $rootScope, $state, AuthService, CompanyService, ProfileService) {
        var self = this;
        $scope.loggedIn = !!getToken("role");
        // CompanyService.getCurrentCompany(function(company) {
        // $scope.companyData = getToken("organization"); //huh?
        // });

        isLoggedInCompany = function() {
            return getToken('auth') !== null;

        };
        // });

        $scope.loggedIn = isLoggedInCompany();
        $scope.isCompany = true;

        CompanyService.getCurrentCompany(function(company) {
            $scope.companyData = company;
        });




        $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams, options) {
                $scope.loggedIn = !!getToken("role");
                // $scope.companyData = getToken("organization");
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
        var self = this;
        $scope.recruitmentList = [];

        CompanyService.getCurrentCompany(function(company) {
            $scope.companyData = company;
            console.log("company data in recruiting controller:");
            console.log($scope.companyData);
            ProfileService.getStudentDataForIds($scope.companyData.studentIds, function(data) {
            	$scope.recruitmentList = data;
            });
        });

        // $scope.companyData = getToken("organization");
        //     ProfileService.getStudentDataForIds($scope.companyData.studentIds, function(data) {
        //         $scope.recruitmentList = data;
        //     });

        $scope.sortableOptions = {
            containment: '#table-container',
            containerPositioning: 'relative'
        };

        $scope.removeRecruit = function (_id) {
            console.log("remove recruit:");
            console.log(_id);
            CompanyService.removeStudentFromWishList(_id, function(data) {
                for (var i = $scope.recruitmentList.length - 1; i >= 0; i--) {
                    if ($scope.recruitmentList[i]._id == _id) {
                        $scope.recruitmentList.splice(i, 1);
                    }
                }
            });

        };

        $scope.viewRecruit = function (_id) {
            $state.go('company.profile', { _id: _id });
        };

        $scope.addStudent = function (_id) {
            console.log("add student:");
            console.log(_id);
        };

        $scope.dragControlListeners = {
            orderChanged: function(obj) {

                CompanyService.switchStudentsInWishList($scope.recruitmentList[obj.source.index]._id, $scope.recruitmentList[obj.dest.index]._id, function(data) {
                    console.log("order changed: ");
                    console.log(data);
                });
            }
        };

    }]);
