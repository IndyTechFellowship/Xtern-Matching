angular.module('Xtern')
    .controller('CompanyMain', ['$scope', '$rootScope', '$state', 'TechPointDashboardService', function($scope, $rootScope, $state, TechPointDashboardService){
        var self = this;
        $scope.loggedIn = isLoggedIn('company');
        $scope.companyName = getToken('companyName');

        $rootScope.$on('$stateChangeStart',
            function(event, toState, toParams, fromState, fromParams, options)
            {
                console.log('change');
                $scope.loggedIn = isLoggedIn('company');
                $scope.companyName = getToken('companyName');
                console.log($scope.loggedIn, $scope.companyName);
            });

        $scope.logout = function () {
            logoutStorage("company");
            $state.go('company.login');
            $scope.loggedIn = false;
        };
    }])
    .controller('CompanyDashboardCtrl', ['$scope', 'TechPointDashboardService', function($scope, TechPointDashboardService){
        //BEGIN CONFIG DATA
        $scope.STARTCHARTSANDSTATS = {
            University: {
                isChart:false,
                title: "Universities",
                icon:'university',
                dataLabel: 'university',
                nestedData: false
            },
            gender:{
                isChart:true,
                title:"Gender",
                labels:['Male','Female'],
                dataLabel:'gender',
                nestedData: false
            },
            Interests: {
                isChart:true,
                title: "Interests",
                dataLabel: 'interestedIn',
                labels: [],
                nestedData: true
            },
            Major: {
                isChart:true,
                title: "Major",
                labels: [],
                dataLabel: 'major',
                nestedData: false
            },
            technology:{
                isChart:true,
                title:"Technology",
                dataLabel:'knownTech',
                labels:[],
                nestedData: true
            }
        };
        $scope.STARTFILTERS = {
            Grade: {
                isToggle: false,
                label: "Grade",
                dataLabel: 'gradeLabel',
                simpleFilter:true,
                nestedHeaders:true,
            },
            GradYear: {
                isToggle: false,
                label: "Graduation Year",
                dataLabel: 'gradYear',
                simpleFilter:true,
                nestedHeaders:true,
            },
            University: {
                isToggle: false,
                label: "University",
                dataLabel: 'university',
                simpleFilter:true,
                nestedHeaders:true,
            },
            Technologies: {
                isToggle: true,
                label: "Technologies",
                dataLabel: 'knownTech',
                simpleFilter:false,
                nestedHeaders:false,
            },
            Interests: {
                isToggle: true,
                label: "Interests",
                dataLabel: 'interestedIn',
                simpleFilter:false,
                nestedHeaders:false,
            },
            Major: {
                isToggle: false,
                label: "Major",
                dataLabel: 'major',
                simpleFilter:true,
                nestedHeaders:true,
            },
            WorkStatus: {
                isToggle: false,
                label: "Work Status",
                dataLabel: 'workStatus',
                simpleFilter:true,
                nestedHeaders:true,
            },
            Name: {
                isToggle: false,
                label: "Name",
                dataLabel: 'name',
                nestedHeaders:true,
            }
        };
        $scope.TABLEHEADERS = [{title: 'Name', displayPropertyName: 'name', sortPropertyName: 'name', sort: 'ascending', selected: false},
            {title: 'Major', displayPropertyName: 'major', sortPropertyName: 'major', sort: 'ascending', selected: false},
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
        $scope.PATH ='company';

        TechPointDashboardService.queryUserSummaryData(function (data) {
            $scope.DATA = data;
        });
        //END CONFIG DATA
    }])
    .controller('CompanyLogin',['$scope','$state', function($scope, $state) {
        console.log('attached');
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
        $scope.login = function(){
            //if($('.ui.form').form('validate form')){
            //    $('.ui.form .message').show();
            //    return false;
            //}
            // do more stuff
            setToken("token_placeholder","company");
            setToken("ININ", "companyName");
            $state.go('company.dashboard');
        };
    }]);
