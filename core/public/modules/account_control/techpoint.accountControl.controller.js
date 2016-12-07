angular.module('Xtern')
    .controller('TechPointAccountCtrl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;

        $scope.techPointUsers = [];
        $scope.instructorUsers = [];
        $scope.companyUsers = [];
        $scope.UserFormData = {};

        $scope.companyList = ["ININ", "Salesforce", "Instructor", "TechPoint"];

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
                console.log('refreshed');
                swapActiveArray($scope.selectedGroup.active);
            }
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


        var resetUserForm = function (user) {
            //$('#accountModalform').form('clear')
            if (user) {
                var nameArr = user.name.split(' ', 2);
                console.log(user);
                $scope.UserFormData._id = user._id;
                $scope.UserFormData.firstName = nameArr[0];
                $scope.UserFormData.lastName = nameArr[1];
                $scope.UserFormData.email = user.email;
                $scope.UserFormData.password = user.password;
                $scope.UserFormData.role = user.role;
                $scope.UserFormData.organization = user.organization;
                $scope.UserFormData.newUser = false;
            }
            else {
                $scope.UserFormData._id = null;
                $scope.UserFormData.firstName = '';
                $scope.UserFormData.lastName = '';
                $scope.UserFormData.email = '';
                $scope.UserFormData.password = '';
                $scope.UserFormData.role = 'TechPoint';
                $scope.UserFormData.organization = 'ININ';
                $scope.UserFormData.newUser = true;
            }
        };

        $scope.launchAddEditUserModal = function (user) {
            console.log(user);
            $('#accountsModal').modal('show');
            resetUserForm(user);
            // $('#accountModalform').form('reset');
            $('#accountModalform .error.message').empty();
        }


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
        }

        $scope.deleteUser = function (user) {
            AccountControlService.deleteUser(user._id, function () {
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
        }

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
        }
        var setup = function () {
            // refreshAccounts('TechPoint', 'TechPoint', $scope.techPointUsers);
            // refreshAccounts('Instructor', 'Instructor', $scope.techPointUsers);
            formConfig();
            modalConfig();
            $scope.selectedGroup.refresh();
        }
        setup();

    }])