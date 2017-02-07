angular.module('Xtern')
    .controller('CommentCtrl', function($scope) {
        console.log($scope);
        if($scope.comment) {
            $('.dimmable.card').dimmer({
                on: 'hover'
            });
        }
    });