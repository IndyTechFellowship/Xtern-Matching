angular.module('Xtern')
    .controller('CompanyMain', ['$scope', '$rootScope', '$state', 'AuthService', function ($scope, $rootScope, $state, AuthService) {
        var self = this;
        $scope.loggedIn = isLoggedInCompany();
        $scope.companyName = getToken('companyName');
        $scope.isCompany = true;

        $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams, options) {
                $scope.loggedIn = isLoggedInCompany();
                $scope.companyName = getToken('companyName');
                if (toState.name == "company.profile") {
                    $('#profile').show();
                }
                else {
                    $('#profile').hide();
                }
            });

        CompanyService.getCompanyDataForId(5733953138851840, function(data)
            // CompanyService.getCompanyDataForId($stateParams._id, function(data)
        {
            $scope.companyData = data;
            console.log("companyData: "+data);
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
                    localStorage.removeItem("auth");
                    $state.go('company.login');
                }
            });
        };
    }])
    .controller('CompanyDashboardCtrl', ['$scope', 'TechPointDashboardService', function ($scope, TechPointDashboardService) {
        //BEGIN CONFIG DATA
        $scope.STARTCHARTSANDSTATS = {
            University: {
                isChart: false,
                title: "Universities",
                icon: 'university',
                dataLabel: 'university',
                nestedData: false
            },
            gender: {
                isChart: true,
                title: "Gender",
                labels: ['Male', 'Female'],
                dataLabel: 'gender',
                nestedData: false
            },
            Interests: {
                isChart: true,
                title: "Interests",
                dataLabel: 'interestedIn',
                labels: [],
                nestedData: true
            },
            Major: {
                isChart: true,
                title: "Major",
                labels: [],
                dataLabel: 'major',
                nestedData: false
            },
            technology: {
                isChart: true,
                title: "Technology",
                dataLabel: 'knownTech',
                labels: [],
                nestedData: true
            }
        };
        $scope.STARTFILTERS = {
            Grade: {
                isToggle: false,
                label: "Grade",
                dataLabel: 'gradeLabel',
                simpleFilter: true,
                nestedHeaders: true,
            },
            GradYear: {
                isToggle: false,
                label: "Graduation Year",
                dataLabel: 'gradYear',
                simpleFilter: true,
                nestedHeaders: true,
            },
            University: {
                isToggle: false,
                label: "University",
                dataLabel: 'university',
                simpleFilter: true,
                nestedHeaders: true,
            },
            Technologies: {
                isToggle: true,
                label: "Technologies",
                dataLabel: 'knownTech',
                simpleFilter: false,
                nestedHeaders: false,
            },
            Interests: {
                isToggle: true,
                label: "Interests",
                dataLabel: 'interestedIn',
                simpleFilter: false,
                nestedHeaders: false,
            },
            Major: {
                isToggle: false,
                label: "Major",
                dataLabel: 'major',
                simpleFilter: true,
                nestedHeaders: true,
            },
            WorkStatus: {
                isToggle: false,
                label: "Work Status",
                dataLabel: 'workStatus',
                simpleFilter: true,
                nestedHeaders: true,
            },
            Name: {
                isToggle: false,
                label: "Name",
                dataLabel: 'name',
                nestedHeaders: true,
            }
        };
        $scope.TABLEHEADERS = [
            {
                title: 'Name',
                displayPropertyName: 'name',
                sortPropertyName: 'name',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Major',
                displayPropertyName: 'major',
                sortPropertyName: 'major',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'School',
                displayPropertyName: 'university',
                sortPropertyName: 'university',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Graduation Year',
                displayPropertyName: 'gradYear',
                sortPropertyName: 'gradYear',
                sort: 'ascending',
                selected: false
            }];
        $scope.DATA = null;
        $scope.PATH = 'company';

        TechPointDashboardService.queryUserSummaryData(function (data) {
            $scope.DATA = data;
        });
        //END CONFIG DATA
    }])
    .controller('CompanyLogin', ['$scope', '$state', 'AuthService', function ($scope, $state, AuthService) {

        var formConfig = function () {
            $('#companyLogin').form({
                fields: {
                    email: {
                        identifier: 'email',
                        rules: [
                            {
                                type: 'empty',
                                prompt: 'Please enter your e-mail'
                            },
                            {
                                type: 'email',
                                prompt: 'Please enter a valid e-mail'
                            }
                        ]
                    },
                    dropdown: {
                        identifier: 'company-dropdown',
                        rules: [
                            {
                                type: 'empty',
                                prompt: 'Please select a dropdown value'
                            }
                        ]
                    },
                    password: {
                        identifier: 'password',
                        rules: [
                            {
                                type: 'empty',
                                prompt: 'Please enter a password'
                            },
                            {
                                type: 'minLength[6]',
                                prompt: 'Your password must be at least {ruleValue} characters'
                            }
                        ]
                    }
                },
                onSuccess: function (event, fields) {
                    authenticate(fields);
                },
                onFailure: function (formErrors, fields) {
                    return '';
                }
            });
        };
        formConfig();

        $scope.login = function () {
            //formConfig();
            $('#companyLogin').form('validate form');
        };
        var authenticate = function (fields) {
            console.log(fields);
            var tempFields = {
                email: "xniccum@gmail.com",
                password: "admin1"
            }
            AuthService.login(tempFields.email, tempFields.password, function (token, err) {
                if (err) {
                    console.log('Login unsuccessful');
                    $('#companyLogin .ui.error.message').html(
                        '<ui class="list"><li>Invalid Username or Passord</li></ui>'
                    );
                } else {
                    setToken(token, "auth");
                    setToken("ININ", "company");
                    AuthService.renderTokens(function (token, err) {
                        if (err) {
                            console.log('Render Token unsuccessful', err);
                            $('#companyLogin .ui.error.message').html(
                                '<ui class="list"><li>A server error occured</li></ui>'
                            ).show();
                        } else {
                            $scope.isCompany = true;
                            $state.go('company.dashboard');
                        }
                    });
                }
            });
        };

        $scope.$on('$viewContentLoaded', function (event, viewConfig) {
            $('select.dropdown').dropdown();
            formConfig();
        });
    }])
    .controller('CompanyRecruiting', ['$scope', '$state', 'ProfileService', 'CompanyService', function ($scope, $state, ProfileService, CompanyService) {
        var self = this;
        $scope.recruitmentList = [];
        // console.log(getToken('auth'));

        CompanyService.getCurrentCompany(function(company) {
            $scope.companyData = company;
            console.log("company data in recruiting controller:");
            console.log($scope.companyData);

            ProfileService.getStudentDataForIds($scope.companyData.studentIds, function(data) {
                $scope.recruitmentList = data;
            });
        });

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
