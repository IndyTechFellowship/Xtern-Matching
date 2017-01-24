angular.module('Xtern')
    .controller('ReviewerDashboardCtrl', ['$scope', '$state', 'ReviewerDashboardService', function($scope, $state, ReviewerDashboardService){

    PATH ='reviewer';

    $scope.isCompany = false;
    $scope.summaryData = null;
    $scope.rawData = null;
    $scope.personsCount = 0;

    //Table Click
    $scope.rowClick = function (key) {
        $state.go(PATH + '.profile', {key: key});
    };

    ReviewerDashboardService.queryUserSummaryData(function (data, keys) {
        let students = [];
        for(var i = 0; i < data.length; i++) {
            students[i] = rowClass(data[i],keys[i]);
        }
        $scope.summaryData = students;
        $scope.rawData = students;
        $scope.personsCount = $scope.summaryData.length;
        $('.ui.dropdown').dropdown();//activates semantic drowpdowns
        $('.ui.accordion').accordion();
    });
}]);