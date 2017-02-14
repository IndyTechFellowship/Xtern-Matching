'use strict';
angular.module('Xtern')
    .controller('CompanyStudentProfileCtrl', function($scope, $location, ProfileService, CompanyService, $stateParams) {
        $scope.addStudent = function (key) {
            CompanyService.addStudentToWishList(key, function(data) {
                // $scope.recruitmentList.push($scope.studentData);
                console.log("Student added");
            });
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            $scope.studentKey = $stateParams.key;
            $('.ui.dropdown').dropdown();
            $('.ui.sticky').sticky({
                context: '#example1'
            });
        });
});