angular.module('Xtern')
    .controller('TechPointMain', ['$scope', '$rootScope', '$state', function ($scope, $rootScope, $state) {
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
    .controller('TechPointAccountCtrl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;

        $scope.techPointUsers = [];
        $scope.instructorUsers = [];
        $scope.companyUsers = [];
        $scope.UserFormData = {};

        $scope.companyList = ["ININ", "Salesforce","Instructor","Company"];

        $scope.selectedGroup = {
            active: 'TechPoint',
            activeCompany: 'ININ',
            changeGroup: function (group) {
                $scope.selectedGroup.active = group;
                swapActiveArray(group);
            },
            changeCompany: function () {
                refreshCompany($scope.selectedGroup.activeCompany);
            },
            selectedUsers: $scope.techPointUsers
        };

        $scope.tableHeaders = [
            { title: 'Name', sortPropertyName: 'name', displayPropertyName: 'name', asc: true },
            { title: 'Email', sortPropertyName: 'email', displayPropertyName: 'email', asc: true }
        ];

        $scope.sort = function (header, event) {
            var prop = header.sortPropertyName;
            var asc = header.asc;
            header.asc = !header.asc;
            var ascSort = function (a, b) { return a[prop] < b[prop] ? -1 : a[prop] > b[prop] ? 1 : 0; };
            var descSort = function (a, b) { return ascSort(b, a); };
            var sortFunc = asc ? ascSort : descSort;
            $scope.techPointUsers.sort(sortFunc);
        };

        $scope.launchAddEditUserModal = function (user) {
           
            $('#accountsModal').modal('show');
             resetUserForm(user);
           // $('#accountModalform').form('reset');
            $('#accountModalform .error.message').empty();
        }

        var resetUserForm = function(user){
            //$('#accountModalform').form('clear')
            if (user) {
                var nameArr = user.name.split(' ', 2);
                console.log(user);
                $scope.UserFormData.firstName = nameArr[0];
                $scope.UserFormData.lastName = nameArr[1];
                $scope.UserFormData.email = user.email;
                $scope.UserFormData.password = user.password;
                $scope.UserFormData.role = user.role;
                $scope.UserFormData.organization = user.organization;
                $scope.UserFormData.newUser = false;
            }
            else{
                $scope.UserFormData.firstName ='';
                $scope.UserFormData.lastName = '';
                $scope.UserFormData.email = '';
                $scope.UserFormData.password = '';
                $scope.UserFormData.role = 'TechPoint';
                $scope.UserFormData.organization = 'ININ';
                $scope.UserFormData.newUser = true;
            }
        };

        var refreshCompany = function (company) {
            AccountControlService.getUsers('Company', company, function (data) {
                $scope.companyUsers.length = 0; //We want to keep array refrences but replace all of the elements 
                data.forEach(function (user) {
                    $scope.companyUsers.push(user);
                });
            });
        };

        var refreshAccounts = function (group, company, array) {
            AccountControlService.getUsers(group, company, function (data) {
                array.length = 0; //We want to keep array refrences but replace all of the elements 
                data.forEach(function (user) {
                    array.push(user);
                });
            });
        }

        var swapActiveArray = function (group) {
            if (group == 'TechPoint') {
                $scope.selectedGroup.selectedUsers = $scope.techPointUsers;
                refreshAccounts(group, group, $scope.techPointUsers);
            } else if (group == 'Instructor') {
                $scope.selectedGroup.selectedUsers = $scope.techPointUsers;
                refreshAccounts(group, group, $scope.techPointUsers);
            }
            else if (group == 'Company') {
                $scope.selectedGroup.selectedUsers = $scope.companyUsers;
                refreshAccounts(group, $scope.selectedGroup.activeCompany, $scope.companyUsers);
            } else {
                //ran out of cases
                $scope.selectedUsers.length = 0;
            }
        };

        var submitUser = function(){
             console.log('pased and submitting', $scope.UserFormData);
             if($scope.UserFormData.newUser){
                 AccountControlService.addUser($scope.UserFormData, function(){});
             }else{
                AccountControlService.updateUser($scope.UserFormData, function(){});
             }
            $('#accountsModal').modal('hide');
        }

        var formConfig = function () {
            $('#accountModalform')
                .form({
                    fields: {
                        fname: {
                            identifier: 'first-name',
                            rules: [
                                {
                                    type: 'empty',
                                    prompt: 'Please enter a first name'
                                }
                            ]
                        },
                        lname: {
                            identifier: 'last-name',
                            rules: [
                                {
                                    type: 'empty',
                                    prompt: 'Please enter a last name'
                                }
                            ]
                        },
                        email: {
                            identifier: 'email',
                            rules: [
                                {
                                    type: 'email',
                                    prompt: 'Please enter a valid e-mail'
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
                        },
                        group: {
                            identifier: 'group',
                            rules: [
                                {
                                    type: 'empty',
                                    prompt: 'Please select a group'
                                }
                            ]
                        },
                    },
                    onSuccess:function(event, fields){
                        submitUser();
                    },
                    onFailure: function(formErrors, fields){
                        return;
                    },
                    keyboardShortcuts: false
                });
        }

        var modalConfig = function(){
            $('#accountsModal').modal({
                closable: false,
                onDeny: function () {
                    return true;
                },
                onApprove: function () {
                   $('#accountModalform').form('validate form');
                   return $('#accountModalform').form('is valid');
                }
            });
        }
        var setup = function () {
            refreshAccounts('TechPoint', 'TechPoint', $scope.techPointUsers);
            refreshAccounts('Instructor', 'Instructor', $scope.techPointUsers);
            swapActiveArray($scope.selectedGroup.active);
            formConfig();
            modalConfig();
        }
        setup();

    }])
    .controller('TechPointDashboardCtrl', ['$scope', 'TechPointDashboardService', function ($scope, TechPointDashboardService) {
        //BEGIN CONFIG DATA
        $scope.STARTCHARTSANDSTATS = {
            University: {
                isChart: false,
                title: "Universities",
                icon: 'university',
                dataLabel: 'university',
                nestedData: false
            },
            status: {
                isChart: true,
                title: "Stage",
                labels: ["Stage 1 Approved", "Remaining", "Denied"],
                dataLabel: 'status',
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
            Status: {
                isToggle: false,
                label: "Status",
                dataLabel: 'status',
                simpleFilter: true,
                nestedHeaders: true,
            },
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
        $scope.PATH = 'techpoint';

        TechPointDashboardService.queryUserSummaryData(function (data) {
            // for(row in data){
            //     data[row].namelink = '<a href="/profile/' + data[row]._id + '">' + data[row].name + "</a>";
            // }
            $scope.DATA = data;
        });


        //END CONFIG DATA
    }])
    .controller('TechpointLogin', ['$scope', '$state', 'AuthService', 'TechPointDashboardService', function ($scope, $state, AuthService, TechPointDashboardService) {
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
            AuthService.login('xniccum@gmail.com', 'admin1', function (token, err) {
                if (err) {
                    console.log('bad login')
                } else {
                    console.log('Login Success');
                    setToken(token, "auth");
                    console.log('Moving');
                    console.log(getToken("auth"));
                    $state.go('techpoint.dashboard');
                    console.log('Didn\'t move');
                }
            });
        };
    }]);
