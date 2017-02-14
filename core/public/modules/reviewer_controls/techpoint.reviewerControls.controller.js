'use strict';
angular.module('Xtern').controller('TechPointReviewerCtrl', function ($scope, $rootScope, $state, TechPointReviewerControlService) {
    var PATH = 'techpoint';

    $scope.reviewGroups = null;
    $scope.reviewGroupKeys = null;
    $scope.selectedGroup = null;

    $scope.createReviewGroups = function () {
        TechPointReviewerControlService.createReviewGroups(20, 2, function (data) {
            TechPointReviewerControlService.queryReviewGroups(function (groups, keys) {
                $scope.reviewGroups = groups;
                $scope.reviewGroupKeys = keys;
            });
        });
    };


    $scope.selectReviewGroup = function (group) {
        $scope.selectedGroup = group;
        console.log("Group Selected", $scope.selectedGroup);
        console.log("Group Selected", $scope.reviewGroups[$scope.selectedGroup]);
    };

    $('.ui.dropdown').dropdown();//activates semantic drowpdowns

    TechPointReviewerControlService.queryReviewGroups(function (groups, keys) {
        $scope.reviewGroups = groups;
        $scope.reviewGroupKeys = keys;
        // console.log("$scope", $scope);
    });
});