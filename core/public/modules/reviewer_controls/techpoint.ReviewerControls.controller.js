angular.module('Xtern').controller('TechPointReviewerCtrl', function ($scope, $rootScope, $state, TechPointReviewerControlService) {
    var self = this;
    PATH ='techpoint';
    $scope.createReviewGroups = function() {
        TechPointReviewerControlService.createReviewGroups(20, 2);
    };


});