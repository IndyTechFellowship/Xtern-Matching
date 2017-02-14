'use strict';
angular.module('Xtern')
    .directive('simpleFilter', function () {
        return {
            restrict: 'E',
            scope: {
                filterobject: '=',//two way binding
                change: '&'//function

            },
            templateUrl: 'public/modules/dashboard/partials/simpleMSFilter.html'
        };
    })
    .directive('toggleFilter', function () {
        return {
            restrict: 'E',
            scope: {
                filterobject: '=',//two way binding
                change: '&'//function
            },
            templateUrl: 'public/modules/dashboard/partials/toggleMSFilter.html'
        };
    })
    .directive('missionControl', function () {// until a better name arrives
        return {
            restrict: 'E',
            scope: {
                startchartsandchats: '=',//two way binding
                startfilters: '=',//two way binding
                tableheaders: '=',//two way binding
                data: '=', //two way binding
                path: '<' // one way binding
            },
            templateUrl: 'public/modules/dashboard/partials/missionControlPage.html'
        };
    });