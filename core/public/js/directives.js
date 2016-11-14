angular.module('Xtern')
    .directive('studentProfileComment', function() {
        return {
            restrict: 'E',
            templateUrl: 'public/shared/partials/components/comment.html',
            // template: '<p>test</p>'
            controller: 'CommentCtrl'
        };
    })
    .directive('techLabels', function () {
        return {
            restrict: 'E',
            // scope: {
            //     labels: labels
            // },
            templateUrl: 'public/shared/partials/components/techLabels.html',
            controller: 'TechLabelsCtrl'
        };
    })
    .directive('onLastRepeat', function (callback) {
        return function (scope, element, attrs) {
            if (scope.$last) setTimeout(function () {
                //scope.$emit('onRepeatLast', element, attrs);
                callback();
            }, 1);
        };
    })
    .directive('onLastRepeatDropDown', function () {
        return function (scope, element, attrs) {
            if (scope.$last) setTimeout(function () {
                $('.ui.dropdown').dropdown();
            }, 1);
        };
    })
    .directive('simpleFilter', function () {
        return {
            restrict: 'E',
            scope: {
                filterobject: '=',//two way binding
                change: '&'//function

            },
            templateUrl: 'public/shared/partials/components/simpleMSFilter.html'
        };
    })
    .directive('toggleFilter', function () {
        return {
            restrict: 'E',
            scope: {
                filterobject: '=',//two way binding
                change: '&'//function
            },
            templateUrl: 'public/shared/partials/components/toggleMSFilter.html'
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
            templateUrl: 'public/shared/partials/components/missionControlPage.html'
        };
    });