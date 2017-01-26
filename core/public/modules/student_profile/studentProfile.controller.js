angular.module('Xtern')
    .controller('StudentProfileCtrl', function($scope, $location, ProfileService) {
    $('.ui.dropdown').dropdown();

    $scope.isCompany = getToken('isCompany');
    $scope.$on('$viewContentLoaded', function (evt) {
        $('.ui.dropdown').dropdown();
    });

});