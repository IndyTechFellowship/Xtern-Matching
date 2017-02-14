'use strict';
angular.module('Xtern')
    .controller('ReviewerLogin', ['$scope', '$state', 'AuthService', function ($scope, $state, AuthService) {

        var formConfig = function () {
            $('#reviewerLogin').form({
                fields: {
                    email: {
                        identifier: 'email',
                        rules: [
                            {
                                type: 'empty',
                                prompt: 'Please enter your e-mail'
                            },
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
                    }
                },
                onSuccess: function (event, fields) {
                    if (event)
                        event.preventDefault();
                    authenticate(fields);
                    return false;
                },
                onFailure: function (formErrors, fields) {
                    return false;
                }
            });
        };
        formConfig();

        $scope.login = function () {
            //formConfig();
            $('#reviewerLogin').form('validate form');
        };
        var authenticate = function (fields) {
            $('#reviewerLogin .ui.button').addClass("disabled");
            AuthService.login(fields.email, fields.password, function (token, org, err) {
                if (err) {
                    $('#reviewerLogin .ui.error.message').html(
                        '<ui class="list"><li>Invalid Username or Password</li></ui>'
                    ).show();
                    $('#reviewerLogin .ui.button').removeClass("disabled");
                } else {
                    console.log('Login Success ' + org);
                    setToken(true, 'isReviewer');
                    $scope.isReviewer = true;
                    $('#reviewerLogin .ui.button').removeClass("disabled");
                    $state.go('reviewer.dashboard');
                }
            });
        };

        $scope.$on('$viewContentLoaded', function (event, viewConfig) {
            $('select.dropdown').dropdown();
            formConfig();
        });
    }]);