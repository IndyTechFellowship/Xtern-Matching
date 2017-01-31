angular.module('Xtern').controller('ReviewerDashboardCtrl', function($scope, $state, ReviewerDashboardService, ReviewerProfileService){
    PATH ='reviewer';

    $scope.isCompany = false;
    $scope.summaryData = null;
    $scope.rawData = null;
    $scope.personsCount = 0;

    $scope.rowClick = function (key) {
        $state.go(PATH + '.profile', {key: key});
    };

    ReviewerDashboardService.queryReviewGroup(function (data, keys, grades) {
        var students = [];
        for(var i = 0; i < data.length; i++) {
            students[i] = rowClass(data[i],keys[i]);
            if(grades[i] !== null && grades[i] > 0) {
                students[i].currentReviewerGrade = grades[i];
            }
        }
        $scope.summaryData = students;
        $scope.rawData = students;
        $scope.personsCount = $scope.summaryData.length;
        $('.ui.dropdown').dropdown();//activates semantic drowpdowns
        $('.ui.accordion').accordion();
    });
});