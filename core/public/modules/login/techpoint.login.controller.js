angular.module('Xtern')
    .controller('TechpointLogin',['$scope','$state','AuthService','TechPointDashboardService', function($scope, $state, AuthService) {
        var formConfig = function() {
            $('#techpointLogin').form({
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
                onSuccess: function(event, fields) {
                    //event.preventDefault();
                    authenticate(fields);
                },
                onFailure: function(formErrors, fields) {
                    return '';
                }
            });
        };

        formConfig();

        $scope.login = function() {
            $('#techpointLogin').form('validate form');
        };
        var authenticate = function(fields) {
            console.log(fields);
            var tempFields = {
                email: "xniccum@gmail.com",
                password: "admin1"
            };
            AuthService.login(tempFields.email, tempFields.password, function(token, err) {
                if (err) {
                    console.log('Login unsuccessful');
                    $('#techpointLogin .ui.error.message').html(
                        '<ui class="list"><li>Invalid Username or Password</li></ui>'
                    );
                } else {
                    AuthService.renderTokens(function(token, err) {
                        if (err) {
                            console.log('Render Token unsuccessful', err);
                            $('#techpointLogin .ui.error.message').html(
                                '<ui class="list"><li>A server error occured</li></ui>'
                            ).show();
                        } else {
                            console.log('Login Success');
                            $scope.isCompany = false;
                            $state.go('techpoint.dashboard');
                        }
                    });
                }
            });
        };

        $scope.$on('$viewContentLoaded', function(event, viewConfig) {
            formConfig();
        });
    }]);
