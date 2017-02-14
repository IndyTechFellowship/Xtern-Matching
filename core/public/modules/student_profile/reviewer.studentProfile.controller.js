'use strict';
angular.module('Xtern').controller('ReviewerStudentProfileCtrl', function($scope, ProfileService, ReviewerProfileService, $stateParams) {
    $('.ui.dropdown').dropdown(); //activate Semantic UI dropdowns

    $scope.comment = {};
    $scope.reviewerGrade = 0;
    $scope.reviewerGradeOptions = [1,2,3,4,5,6,7,8,9,10];

    $scope.selectReviewerGrade = function(option) {
        $scope.reviewerGrade = option;
        ReviewerProfileService.postReviewerGradeForStudent($stateParams.key, option);
    };

    var ReviewerStudentProfileCtrlSetup = function(){
        $('.ui.sticky').sticky({
            context: '#example1'
        });

        $('.ui.dropdown').dropdown();        
    };

    $scope.$on('$viewContentLoaded', function (evt) {
        ReviewerStudentProfileCtrlSetup();
        ReviewerProfileService.getReviewerGradeForStudent($stateParams.key, function (reviewerGrade) {
            $scope.reviewerGrade = reviewerGrade;
        });
    });
});