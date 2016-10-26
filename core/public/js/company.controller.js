angular.module('Xtern')
    .controller('CompanyMain', ['$scope', '$rootScope', '$state', 'TechPointDashboardService', 'CompanyService', function ($scope, $rootScope, $location, $state, CompanyService, TechPointDashboardService, $stateParams) {
        var self = this;
        $scope.loggedIn = isLoggedIn('company');
        $scope.companyName = getToken('companyName');
        $scope.isCompany = true;

        CompanyService.getCompanyDataForId(5733953138851840, function(data)            
        // CompanyService.getCompanyDataForId($stateParams._id, function(data)            
        {
            $scope.companyData = data;
            console.log("companyData: "+data);
        });

        $rootScope.$on('$stateChangeStart', function (event, toState, toParams, fromState, fromParams, options) {
           $scope.loggedIn = isLoggedIn('company');
            $scope.companyName = getToken('companyName');
           if(toState.name == "company.profile"){
                $('#profile').show();
            }
            else{
                $('#profile').hide();
            }
        });

        $scope.addStudentToCompany = function (studentID) {
            CompanyService.addStudentToWishList(studentID, function (data) {
                console.log(data);
            });
        };

        $scope.logout = function () {
            localStorage.removeItem("auth");
            localStorage.removeItem("role");
            logoutStorage("auth");
            $state.go('company.login');
            $scope.loggedIn = false;
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
    .controller('CompanyLogin', ['$scope', '$state','AuthService', function ($scope, $state,AuthService) {
        $scope.company = "ININ";
        //$('.ui.form')
        //    .form({
        //        fields: {
        //            email: {
        //                identifier: 'email',
        //                rules: [
        //                    {
        //                        type: 'empty',
        //                        prompt: 'Please enter your e-mail'
        //                    },
        //                    {
        //                        type: 'email',
        //                        prompt: 'Please enter a valid e-mail'
        //                    }
        //                ]
        //            },
        //            password: {
        //                identifier: 'password',
        //                rules: [
        //                    {
        //                        type: 'empty',
        //                        prompt: 'Please enter your password'
        //                    },
        //                    {
        //                        type: 'length[6]',
        //                        prompt: 'Your password must be at least 6 characters'
        //                    }
        //                ]
        //            }
        //        }
        //    });
        $scope.login = function () {
            //if($('.ui.form').form('validate form')){
            //    $('.ui.form .message').show();
            //    return false;
            //}
            // do more stuff
            //setToken("token_placeholder", "company");
            AuthService.login('xniccum@gmail.com','admin1', function (token,err) {
                if (err) {
                    console.log('bad login');
                } else {
                    localStorage.setItem("auth", token);
                    $scope.isCompany = true;
                    $state.go('company.dashboard');
                }
            });
        };
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
