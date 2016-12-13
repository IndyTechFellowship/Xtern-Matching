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
                    if(event)
                        event.preventDefault();
                    authenticate(fields);
                    return false;
                },
                onFailure: function(formErrors, fields) {    
                    return false;
                }
            });
        };

        formConfig();

        $scope.login = function() {
            $('#techpointLogin').form('validate form');
        };
        var authenticate = function(fields) {
            $('#techpointLogin .ui.button').addClass("disabled");
            AuthService.login(fields.email, fields.password, function(token, err) {
                if (err) {
                    $('#techpointLogin .ui.error.message').html(
                        '<ui class="list"><li>Invalid Username or Password</li></ui>'
                    ).show();
                    $('#techpointLogin .ui.button').removeClass("disabled");
                } else {
                    AuthService.renderTokens(function(token, err) {
                        if (err) {
                            console.log('Render Token unsuccessful', err);
                            $('#techpointLogin .ui.error.message').html(
                                '<ui class="list"><li>A server error occured</li></ui>'
                            ).show();
                            $('#techpointLogin .ui.button').removeClass("disabled");
                        } else {
                            $('#techpointLogin .ui.button').removeClass("disabled");
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
