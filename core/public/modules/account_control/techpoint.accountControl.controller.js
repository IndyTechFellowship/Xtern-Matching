angular.module('Xtern')
    .controller('TechPointAccountCtrl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;

        $scope.techPointUsers = [];
        $scope.instructorUsers = [];
        $scope.companyUsers = [];
        $scope.UserFormData = {};

        //TODO: REPLACE WITH BACKEND CALL
        $scope.companyList = COMPANY_GLOBAL_LIST;

        var declarePageVars = function () {
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
                selectedUsers: $scope.techPointUsers,
                refresh: function () {
                    console.log('2');
                    console.log('refreshed');
                    swapActiveArray($scope.selectedGroup.active);
                }
            };

            $scope.tableHeaders = [
                { title: 'Name', sortPropertyName: 'name', displayPropertyName: 'name', asc: true },
                { title: 'Email', sortPropertyName: 'email', displayPropertyName: 'email', asc: true }
            ];

            //Set up CompanyAbbr
            $scope.companyListAbbr = $scope.companyList.filter(function(item){
                return !(item == 'TechPoint' || item == 'Instructor');
            });
        };

        $scope.sort = function (header, event) {
            var prop = header.sortPropertyName;
            var asc = header.asc;
            header.asc = !header.asc;
            var ascSort = function (a, b) { return a[prop] < b[prop] ? -1 : a[prop] > b[prop] ? 1 : 0; };
            var descSort = function (a, b) { return ascSort(b, a); };
            var sortFunc = asc ? ascSort : descSort;
            $scope.techPointUsers.sort(sortFunc);
        };

        var resetUserForm = function (user) {
            //$('#accountModalform').form('clear')
            if (user) {
                var nameArr = user.name.split(' ', 2);
                console.log(user);
                $scope.UserFormData.key = user.key;
                $scope.UserFormData.firstName = nameArr[0];
                $scope.UserFormData.lastName = nameArr[1];
                $scope.UserFormData.email = user.email;
                $scope.UserFormData.password = user.password;
                //$scope.UserFormData.organization = 'ININ';
                $scope.UserFormData.newUser = false;
            }
            else {
                $scope.UserFormData.key = null;
                $scope.UserFormData.firstName = '';
                $scope.UserFormData.lastName = '';
                $scope.UserFormData.email = '';
                $scope.UserFormData.password = '';
                //$scope.UserFormData.organization = 'ININ';
                $scope.UserFormData.newUser = true;
            }
        };

        $scope.launchAddEditUserModal = function (user) {
            console.log(user);
            $('#accountsModal').modal('show');
            resetUserForm(user);
            $('#accountModalform .error.message').empty();
        };


        var refreshCompany = function (company) {
            AccountControlService.getUsers(function (users,keys) {
                $scope.companyUsers.length = 0; //We want to keep array refrences but replace all of the elements
                for(var i = 0; i < keys.length ; i++) {
                    users[i] = keys[i];
                    $scope.companyUsers.push(users[i]);
                }
            });
        };

        var refreshAccounts = function (array) {
            AccountControlService.getUsers(function (users,keys) {
                array.length = 0; //We want to keep array refrences but replace all of the elements
                for(var i = 0; i < keys.length ; i++) {
                    users[i] = keys[i];
                    array.push(users[i]);
                }
            });
        };

        var swapActiveArray = function (group) {
            if (group == 'TechPoint') {
                console.log('3.a');
                $scope.selectedGroup.selectedUsers = $scope.techPointUsers;
                refreshAccounts($scope.techPointUsers);
            } else if (group == 'Instructor') {
                console.log('3.b');
                $scope.selectedGroup.selectedUsers = $scope.techPointUsers;
                refreshAccounts($scope.techPointUsers);
            }
            else if (group == 'Company') {
                console.log('3.c');
                $scope.selectedGroup.selectedUsers = $scope.companyUsers;
                refreshAccounts($scope.companyUsers);
            } else {
                //ran out of cases
                $scope.selectedUsers.length = 0;
            }
        };

        var submitUser = function () {
            console.log('pased and submitting', $scope.UserFormData);
            $scope.UserFormData.name = $scope.UserFormData.firstName + " " + $scope.UserFormData.lastName;
            if ($scope.UserFormData.newUser) {
                AccountControlService.addUser($scope.UserFormData, function () {
                    $scope.selectedGroup.refresh();
                });
            } else {
                AccountControlService.updateUser($scope.UserFormData, function () {
                    $scope.selectedGroup.refresh();
                });
            }
            $('#accountsModal').modal('hide');
        };

        $scope.deleteUser = function (user) {
            AccountControlService.deleteUser(user.key, function () {
                $scope.selectedGroup.refresh();
            });
        };

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
                    onSuccess: function (event, fields) {
                        submitUser();
                    },
                    onFailure: function (formErrors, fields) {
                        return;
                    },
                    keyboardShortcuts: false
                });
        };

        var modalConfig = function () {
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
        };
        var setup = function () {
            declarePageVars();
            formConfig();
            modalConfig();
            $scope.selectedGroup.refresh();
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);