angular.module('Xtern')
    .controller('TechPointAccountCtrl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
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
                    // refreshCompany($scope.selectedGroup.activeCompany);
                    swapActiveArray('Company');
                },
                selectedUsers: $scope.techPointUsers,
                refresh: function () {
                    swapActiveArray($scope.selectedGroup.active);
                }
            };

            $scope.tableHeaders = [
                { title: 'Name', sortPropertyName: 'name', displayPropertyName: 'name', asc: true },
                { title: 'Email', sortPropertyName: 'email', displayPropertyName: 'email', asc: true }
            ];

            //Set up CompanyAbbr
            $scope.companyListAbbr = $scope.companyList.filter(function (item) {
                return !(item.name == 'Techpoint' || item.name == 'Instructor' || item.name == 'Instructors');
            });
        };

        var resetUserForm = function (user) {
            if (user) {
                var nameArr = user.name.split(' ', 2);
                $scope.UserFormData.newUser = false;
                $('#accountModalform').form('set values', {
                    firstName: nameArr[0],
                    lastName: nameArr[1],
                    email: user.email,
                    password: user.password,
                    role: user.organization,
                    organization: user.organization,
                    key: user.key
                });
                $('.two.fields.group-role').hide();
            }
            else {
                $('.two.fields.group-role').show();
                $('#accountModalform').form('reset');
                $scope.hideCompanyDropdown();
                setSelectOptions();
            }
        };

        $scope.launchAddEditUserModal = function (user) {
            $('#accountsModal').modal('show');
            resetUserForm(user);
            $('#accountModalform .error.message').empty();
        };

        var refreshAccounts = function (group, company, array) {
            AccountControlService.getUsers(group, company, function (data) {
                array.length = 0; //We want to keep array refrences but replace all of the elements 
                data.forEach(function (user) {
                    array.push(user);
                });
            });
        }
        var refreshAccounts = function (organizationKey, array) {
            AccountControlService.getUsers(organizationKey, function (users) {
                array.length = 0; //We want to keep array refrences but replace all of the elements
                users.forEach(function (user) {
                    array.push(user);
                });
                $scope.selectedGroup.selectedUsers = array;
            });
        };

        var swapActiveArray = function (group) {
            if (group == 'TechPoint') {
                $scope.selectedGroup.selectedUsers = $scope.techPointUsers;
                $scope.companyList.forEach(function (company) {
                    if (company.name == 'Techpoint') {
                        refreshAccounts(company.key, $scope.techPointUsers);
                    }
                });
            } else if (group == 'Instructor') {
                $scope.selectedGroup.selectedUsers = $scope.instructorUsers;
                $scope.companyList.forEach(function (company) {
                    if (company.name == 'Instructor' || company.name== 'Instructors') {
                        refreshAccounts(company.key, $scope.instructorUsers);
                    }
                });
            }
            else if (group == 'Company') {
                $scope.selectedGroup.selectedUsers = $scope.companyUsers;
                if($scope.selectedGroup.activeCompany){
                    refreshAccounts($scope.selectedGroup.activeCompany, $scope.companyUsers);
                }
            } else {
                //ran out of cases
                $scope.selectedUsers.length = 0;
            }
        };

        var submitUser = function (fields) {
            //                 console.log('passed and submitting');
            //     $scope.UserFormData.name = $scope.UserFormData.firstName + " " + $scope.UserFormData.lastName;
            //     if ($scope.UserFormData.newUser) {
            //         AccountControlService.addUser($scope.UserFormData, function () {
            fields.name = fields.firstName + " " + fields.lastName;
            if (!fields.key) {
                if (fields.role != 'Company') {
                    fields.organization = fields.role;
                };
                AccountControlService.addUser(fields, function (data) {

                    $scope.selectedGroup.refresh();
                });
            } else {
                AccountControlService.updateUser(fields, function (data) {
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

        var accountControlFormConfig = function () {
            $('#accountModalform')
                .form({
                    fields: {
                        fname: {
                            identifier: 'firstName',
                            rules: [
                                {
                                    type: 'empty',
                                    prompt: 'Please enter a first name'
                                }
                            ]
                        },
                        lname: {
                            identifier: 'lastName',
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
                            identifier: 'role',
                            rules: [
                                {
                                    type: 'empty',
                                    prompt: 'Please select a group'
                                }
                            ]
                        },
                        organization: {
                            identifier: 'organization',
                            rules: [
                                {
                                    type: 'empty',
                                    prompt: 'Please select a group'
                                }
                            ]
                        },
                    },
                    onSuccess: function (event, fields) {
                        submitUser(fields);
                    },
                    onFailure: function (formErrors, fields) {
                        return;
                    },
                    keyboardShortcuts: false
                });
        };
        var accountControlModalConfig = function () {

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
                accountControlFormConfig();
                accountControlModalConfig();
                $scope.selectedGroup.refresh();
            });
        };

        $scope.hideCompanyDropdown = function () {
            $('#companyDropdown').hide();
        };

        $scope.showCompanyDropdown = function () {
            $('#companyDropdown').show();
        };

        var setSelectOptions = function () {
            $('.role.dropdown').dropdown({
                onChange: function (text, value) {
                    if (value == 'Company') {
                        $('#companyDropdown').show();
                    } else {
                        $('#companyDropdown').hide();
                    }
                }
            });
        };
        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);