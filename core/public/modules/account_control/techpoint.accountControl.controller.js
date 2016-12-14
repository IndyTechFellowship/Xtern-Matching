angular.module('Xtern')
    .controller('TechPointAccountCtrl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;

        $scope.techPointUsers = [];
        $scope.instructorUsers = [];
        $scope.companyUsers = [];
        $scope.UserFormData = {};


        var declarePageVars = function () {
            $scope.selectedGroup = {
                active: 'TechPoint',
                activeCompany: $scope.companyList[0].key,
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
            $scope.companyList.forEach(function (company) {
               if(company.name == 'Techpoint') {
                   $scope.selectedGroup.activeCompany = company.key;
               }
            });

            $scope.tableHeaders = [
                { title: 'Name', sortPropertyName: 'name', displayPropertyName: 'name', asc: true },
                { title: 'Email', sortPropertyName: 'email', displayPropertyName: 'email', asc: true }
            ];

            //Set up CompanyAbbr
            $scope.companyListAbbr = $scope.companyList.filter(function(item) {
                return !(item.name == 'Techpoint' || item.name == 'Instructor');
            });
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
                $scope.UserFormData.organization = $scope.companyList;
                $scope.UserFormData.newUser = false;
            }
            else {
                $scope.UserFormData.key = null;
                $scope.UserFormData.firstName = '';
                $scope.UserFormData.lastName = '';
                $scope.UserFormData.email = '';
                $scope.UserFormData.password = '';
                $scope.UserFormData.organization = $scope.companyList[0].key;
                $scope.UserFormData.newUser = true;
            }
        };


        var refreshCompany = function (organizationKey) {
            AccountControlService.getUsers(organizationKey,function (users) {
                $scope.companyUsers.length = 0; //We want to keep array refrences but replace all of the elements
                console.log("here: ", users);
                users.forEach(function(user) {
                    $scope.companyUsers.push(user);
                });
            });
        };

        var swapActiveArray = function (group) {
            if (group == 'TechPoint') {
                $scope.companyList.forEach(function (company) {
                    if(company.name == 'Techpoint') {
                        refreshAccounts(company.key, $scope.techPointUsers);
                    }
                });
                //$scope.selectedGroup.selectedUsers = $scope.techPointUsers;
            } else if (group == 'Instructor') {
                $scope.companyList.forEach(function (company) {
                    if(company.name == 'Techpoint') {
                        refreshAccounts(company.key, $scope.techPointUsers);
                    }
                });
                //$scope.selectedGroup.selectedUsers = $scope.techPointUsers;
            }
            else if (group == 'Company') {
                refreshAccounts($scope.selectedGroup.activeCompany, $scope.companyUsers);
                $scope.selectedGroup.selectedUsers = $scope.companyUsers;
            } else {
                //ran out of cases
                $scope.selectedUsers.length = 0;
            }
        };

        var refreshAccounts = function (organizationKey,array) {
            AccountControlService.getUsers(organizationKey,function (users) {
                array.length = 0; //We want to keep array refrences but replace all of the elements
                users.forEach(function(user) {
                    array.push(user);
                });
                $scope.selectedGroup.selectedUsers = array;
            });
        };

        var submitUser = function () {
            console.log('passed and submitting');
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

        $scope.launchAddEditUserModal = function (user) {
            $('#accountsModal').modal('show');
            resetUserForm(user);
            $('#accountModalform .error.message').empty();
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

        $scope.deleteUser = function (user) {
            AccountControlService.deleteUser(user.key, function () {
                $scope.selectedGroup.refresh();
            });
        };

        $scope.rowClick = function(row) {
            console.log($scope.selectedGroup.selectedUsers);
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
            AccountControlService.getOrganizations(function (organizations) {
                $scope.companyList = organizations;
                declarePageVars();
                formConfig();
                modalConfig();
                $scope.selectedGroup.refresh();
            });
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);