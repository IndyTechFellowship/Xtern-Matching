// Declare app level module which depends on filters, and services
'use strict';
angular.module('Xtern-Matching', ['ngResource', 'ngRoute', 'ui.bootstrap'])
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/home/home.html',
        controller: 'HomeController',
        controllerAs: 'home'
      })
      .otherwise({redirectTo: '/'});
  }]);
