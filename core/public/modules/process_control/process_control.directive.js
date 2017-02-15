'use strict';
angular.module('Xtern')
    .directive('decisionBoard', function () {
        return {
            restrict: 'E',
            scope: {
                title: '<',//one way binding
                slider: '=',//two way binding,
                charts: '=',
                histogram: '=',
                displayList:'=',
                length:'='
                //change: '&'//function

            },
            templateUrl: 'public/modules/process_control/partials/decisionBoard.html'
        };
    });