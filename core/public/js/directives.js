angular.module('Xtern')
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
    });