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
            
        }

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
                    key: 212
                });
                $('.two.fields.group-role').hide();
            }
            else {
                 $('.two.fields.group-role').show();
                 $('#accountModalform').form('reset');                 
            }
        };

        $scope.launchAddEditUserModal = function (user) {
            console.log(user);
            $('#accountsModal').modal('show');
            resetUserForm(user);
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

        var submitUser = function (fields) {
            console.log('pased and submitting', fields);
            fields.name = fields.firstName + " " + fields.lastName;
            if (!fields.key) {
                AccountControlService.addUser(fields, function (data) {
                    $scope.selectedGroup.refresh();
                });
            } else {
                AccountControlService.updateUser(fields, function (data) {
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
            declarePageVars();
            formConfig();
            modalConfig();
            $scope.selectedGroup.refresh();
        }

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);