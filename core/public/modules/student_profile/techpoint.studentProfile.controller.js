'use strict';
angular.module('Xtern')
    .controller('TechPointStudentProfileCtrl', function($scope, $location, ProfileService, $stateParams) {
        $('.ui.dropdown').dropdown();
        $scope.statusOptions = [
            'Stage 1 Approved',
            'Stage 2 Approved',
            'Stage 3 Approved',
            'Undecided',
            'Rejected (Stage 1)',
            'Rejected (Stage 2)',
            'Rejected (Stage 3)'
        ];
        $scope.r1GradeOptions = [];

        $scope.selectStatus = function(option) {
            ProfileService.setStatus($scope.studentKey,option, function (err) {
                if(err) {}
            });
        };

        $scope.selectR1Grade = function(option) {
            ProfileService.setR1Grade($scope.studentKey,option, function (err) {
                if(err) {}
            });
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            $scope.studentKey = $stateParams.key
            $('.ui.sticky').sticky({
                context: '#example1'
            });
            $('.ui.dropdown').dropdown();
            for(let i=0; i < 10; i += 0.5) {
                $scope.r1GradeOptions.push(i);
            }
        });
});
