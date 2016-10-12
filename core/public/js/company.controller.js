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
                    console.log('bad login')
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
        companyId = 5733953138851840;

        CompanyService.getCompanyDataForId(companyId, function(data)            
        // CompanyService.getCompanyDataForId($stateParams._id, function(data)            
        {
            $scope.companyData = data;
            // console.log("companyData: "+data);
            console.log("company data in recruiting controller:");
            console.log($scope.companyData);
            console.log($scope.companyData.studentIds);
            // $scope.recruitmentList = [];
            for (i=0;i<$scope.companyData.studentIds.length; i++) {
                ProfileService.getStudentDataForId($scope.companyData.studentIds[i], function(data) {
                    data.name = data.firstName+" "+data.lastName;
                    $scope.studentData = data;
                    // console.log($scope.studentData);
                    $scope.recruitmentList.push($scope.studentData);
                });
            }

        });

        // $scope.recruitmentList = [
        //     {
        //         _id: "57269aa3bf79bbf8cc55d9d",
        //         name: "Verna Gomez",
        //         gradYear: 2019,
        //         university: "Rose-Hulman Institute of Technology",
        //         summary: "Front End",
        //         notes: "Verna would be a great addition to Aarons Front end team. They use similar tools"
        //     },
        //     {
        //         _id: "573a010c27b02303a5819515",
        //         name: "Henderson Whitley",
        //         gradYear: 2018,
        //         university: "Indiana State University",
        //         summary: "Security",
        //         notes: "Phasellus ex nisl, pulvinar tempus dolor non, aliquam maximus ante.  Sed et nunc lectus. Phasellus eget lectus sit amet felis interdum tristique."
        //     },
        //     {
        //         _id: "573a010cdaf1dc6ea094593e",
        //         name: "Cross Berg",
        //         gradYear: 2018,
        //         university: "Indiana State University",
        //         summary: "Security",
        //         notes: "Fusce a est pulvinar, dictum tellus ut, cursus risus. Morbi bibendum elementum risus, in cursus dolor dictum luctus."
        //     },
        //     {
        //         _id: "573a010cc2cac4dfe9497bb2",
        //         name: "Loraine Pace",
        //         gradYear: 2017,
        //         university: "Rose-Hulman Institute of Technology",
        //         summary: "SE - Full Stack (Eric's Team)",
        //         notes: "Ut sollicitudin nunc ac mauris hendrerit consectetur. Pellentesque imperdiet ullamcorper augue et fermentum."
        //     },
        //     {
        //         _id: "573a010c7afd6700cb9c8598",
        //         name: "Maureen Mclean",
        //         gradYear: 2018,
        //         university: "Indiana State University",
        //         summary: "CPE - Hardware",
        //         notes: "Integer laoreet ornare interdum. Nunc dapibus elit et purus scelerisque rhoncus. Nullam sagittis nulla eget diam scelerisque euismod."
        //     },
        //     {
        //         _id: "573a010cfcbfb6015c7a6669",
        //         name: "Bell Simon",
        //         gradYear: 2017,
        //         university: "Rose-Hulman Institute of Technology",
        //         summary: "Backend API - Main Product",
        //         notes: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In ac suscipit velit. Fusce sollicitudin non massa ac blandit."
        //     }
        // ];

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

        // $scope.dragControlListeners.itemMoved = function(obj) {
        //     console.log("item moved: ");
        //     console.log(obj);
        // };

        $scope.dragControlListeners = {
            // accept: function (sourceItemHandleScope, destSortableScope) {return boolean}//override to determine drag is allowed or not. default is true.
            // itemMoved: function(obj) {
            //     console.log("item moved: ");
            //     console.log(obj);
            // },
            orderChanged: function(obj) {

                CompanyService.switchStudentsInWishList($scope.recruitmentList[obj.source.index]._id, $scope.recruitmentList[obj.dest.index]._id, function(data) {
                    console.log("order changed: ");
                    console.log(data);
                });
            }
            // containment: '#board'//optional param.
            // clone: true //optional param for clone feature.
            // allowDuplicates: false //optional param allows duplicates to be dropped.
        };

    }]);
