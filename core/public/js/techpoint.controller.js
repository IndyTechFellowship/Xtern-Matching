angular.module('Xtern')
    .controller('TechPointMain', ['$scope', '$rootScope', '$state', 'TechPointDashboardService', function($scope, $rootScope, $state, TechPointDashboardService){
        var self = this;
        $scope.loggedIn = isLoggedIn('techPoint');

        $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams, options) {
                $scope.loggedIn = isLoggedIn('techPoint');
                if (toState.name == "techpoint.profile") {
                    $('#profile').show();
                }
                else {
                    $('#profile').hide();
                }
            });

        $scope.logout = function () {
            logoutStorage("auth");
            $state.go('techpoint.login');
            $scope.loggedIn = false;
        };
    }])
    .controller('TechPointDashboardCtrl', ['$scope', 'TechPointDashboardService', function($scope, TechPointDashboardService){
        //BEGIN CONFIG DATA
        $scope.STARTCHARTSANDSTATS = {
            University: {
                isChart:false,
                title: "Universities",
                icon:'university',
                dataLabel: 'university',
                nestedData: false
            },
            status:{
                isChart:true,
                title:"Stage",
                labels: ["Stage 1 Approved", "Remaining", "Denied"],
                dataLabel:'status',
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
            Status: {
                isToggle: false,
                label: "Status",
                dataLabel: 'status',
                simpleFilter:true,
                nestedHeaders:true,
            },
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
        $scope.TABLEHEADERS = [
            {
                title: 'Name',
                displayPropertyName: 'namelink',
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
            },
            {
                title: 'Status',
                displayPropertyName: 'status',
                sortPropertyName: 'status',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Grade',
                displayPropertyName: 'gradeLabel',
                sortPropertyName: 'gradeValue',
                sort: 'descending',
                selected: false
            }];
        $scope.DATA = null;
        $scope.PATH ='techpoint';

        TechPointDashboardService.queryUserSummaryData(function (data) {
            // for(row in data){
            //     data[row].namelink = '<a href="/profile/' + data[row]._id + '">' + data[row].name + "</a>";
            // }
            $scope.DATA = data;
        });


        //END CONFIG DATA
    }])
    .controller('TechpointLogin',['$scope','$state','AuthService','TechPointDashboardService', function($scope, $state, AuthService, TechPointDashboardService) {
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
            AuthService.login('xniccum@gmail.com','admin1', function (token,err) {
                if (err) {
                    console.log('bad login')
                } else {
                    console.log('Login Success');
                    setToken(token,"auth");
                    console.log('Moving');
                    console.log(getToken("auth"));
                    $state.go('techpoint.dashboard');
                    console.log('Didn\'t move');
                }
            });
        };
    }]);
