'use strict';
angular.module('Xtern')
    .controller('TechPointAccountCtrl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {

        $scope.keys = {
            techpoint: '',
            reviewer: '',
            companies: [],
            companyKeys: {}
        };
        $scope.techPointUsers = [];
        $scope.reviewerUsers = [];
        $scope.companyUsers = [];
        $scope.UserFormData = {};


        var declarePageVars = function () {
            $scope.selectedGroup = {
                active: 'TechPoint',
                activeCompany: $scope.keys.companies[0].key,
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
                {title: 'Name', sortPropertyName: 'name', displayPropertyName: 'name', asc: true},
                {title: 'Email', sortPropertyName: 'email', displayPropertyName: 'email', asc: true}
            ];
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

        $scope.launchAddCompanyModal = function (user) {
            $('#addCompanyForm').form('reset');
            $('#addCompanyForm .error.message').empty();
            $('#addCompanyModal').modal('show');
        };

        var refreshAccounts = function (organizationKey, array) {
            AccountControlService.getUsers(organizationKey, function (users) {
                array.length = 0; //We want to keep array refrences but replace all of the elements
                if(users) {
                    users.forEach(function (user) {
                        array.push(user);
                    });
                    $scope.selectedGroup.selectedUsers = array;
                }
            });
        };

        var swapActiveArray = function (group) {
            $scope.selectedGroup.selectedUsers = [];
            //$scope.selectedUsers.length = 0;
            if (group == 'TechPoint') {
                $scope.selectedGroup.selectedUsers = $scope.techPointUsers;
                refreshAccounts($scope.keys.techpoint, $scope.techPointUsers);

            } else if (group == 'Reviewer') {
                $scope.selectedGroup.selectedUsers = $scope.reviewerUsers;
                refreshAccounts($scope.keys.reviewer, $scope.reviewerUsers);

            }
            else if (group == 'Company') {
                $scope.selectedGroup.selectedUsers = $scope.companyUsers;
                if ($scope.selectedGroup.activeCompany) {
                    refreshAccounts($scope.selectedGroup.activeCompany, $scope.companyUsers);
                }
            } else {
                //ran out of cases
                $scope.selectedUsers.length = 0;
            }
        };

        var submitUser = function (fields) {
            fields.name = fields.firstName + " " + fields.lastName;
            if (!fields.key) {
                if (fields.role == "Techpoint" || fields.role == "TechPoint") {
                    fields.organization = $scope.keys.techpoint;
                } else if (fields.role == "Reviewer" || fields.role == "Reviewer") {
                    fields.organization = $scope.keys.reviewer;
                }

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

        var addCompany = function (name) {
            AccountControlService.addCompany(name, function (data) {
                AccountControlService.getOrganizations(function (organizations) {
                    $scope.companyList = organizations;
                    $scope.keys.companies.length =0;
                  //  $scope.keys.companyKeys = [];
                    organizations.forEach(function (org) {
                        if(!(org.name == "Reviewers" || org.name == "Reviewer") && !(org.name == "TechPoint" || org.name == "Techpoint")) {
                            $scope.keys.companies.push(org);
                            $scope.keys.companyKeys[org.name] = org.key;
                        }
                    });
                    $scope.keys.companies.sort();
                });
            });
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
            var ascSort = function (a, b) {
                return a[prop] < b[prop] ? -1 : a[prop] > b[prop] ? 1 : 0;
            };
            var descSort = function (a, b) {
                return ascSort(b, a);
            };
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
                        }
                    },
                    onSuccess: function (event, fields) {
                        submitUser(fields);
                    },
                    onFailure: function (formErrors, fields) {
                        return false;
                    },
                    keyboardShortcuts: false
                });

            $.fn.form.settings.rules.companyDoesNotExist = function (value, arg1) {
                for (var org in $scope.keys.companies) {
                    if ($scope.keys.companies[org].name == value) {
                        return false;
                    }
                }
                return true;
            };

            $('#addCompanyForm').form({
                fields: {
                    name: {
                        identifier: 'name',
                        rules: [
                            {type: 'empty', prompt: 'Company Name cannot be blank'},
                            {type: 'companyDoesNotExist[0]', prompt: '{value} is already a company'}
                        ]
                    },
                    confirm: {
                        identifier: 'confirm',
                        rules: [
                            {type: 'match[name]', prompt: 'Company Names don\'t match'}
                        ]
                    }
                },
                onSuccess: function (event, fields) {
                    addCompany(fields.name);
                    return true;
                },
                onFailure: function (formErrors, fields) {
                    return false;
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

            $('#addCompanyModal').modal({
                closable: false,
                onDeny: function () {
                    return true;
                },
                onApprove: function () {
                    $('#addCompanyForm').form('validate form');
                    return $('#addCompanyForm').form('is valid');
                }
            });
        };

        var setup = function () {
            AccountControlService.getOrganizations(function (organizations) {
                $scope.companyList = organizations;
                organizations.forEach(function (org) {
                    if (org.name == "Reviewers" || org.name == "Reviewer") {
                        $scope.keys.reviewer = org.key;
                    } else if (org.name == "TechPoint" || org.name == "Techpoint") {
                        $scope.keys.techpoint = org.key;
                    }
                    else {
                        $scope.keys.companies.push(org);
                        $scope.keys.companyKeys[org.name] = org.key;
                    }
                });
                $scope.keys.companies.sort();
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