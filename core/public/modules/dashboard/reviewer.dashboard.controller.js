angular.module('Xtern').controller('ReviewerDashboardCtrl', function($scope, $state, ReviewerDashboardService, ReviewerProfileService){
    PATH ='reviewer';

    $scope.isCompany = false;
    $scope.summaryData = null;
    $scope.rawData = null;
    $scope.personsCount = 0;

    $scope.filters ={
        graded: false
    };

    $scope.change = function(){        
        $scope.summaryData = $scope.rawData.filter(function(row){
            if($scope.filters.graded && row.currentReviewerGrade){
                return false;
            }
            return true;
        });
        $scope.personsCount = $scope.summaryData.length;
    };

    $scope.tableHeaders = [
        {title: 'Name', displayProperty:'name', sortPropertyName: 'name', asc: true},
        {title: 'Your Grade', displayProperty:'currentReviewerGrade', sortPropertyName: 'sortReviewerGrade', asc: true}
    ];

    $scope.rowClick = function (key) {
        $state.go(PATH + '.profile', {key: key});
    };

    ReviewerDashboardService.queryReviewGroup(function (data, keys, grades) {
        var students = [];
        for(var i = 0; i < data.length; i++) {
            students[i] = rowClass(data[i],keys[i]);
            students[i].name = students[i].firstName + " " + students[i].lastName;
            if(grades[i] !== null && grades[i] > 0) {
                students[i].currentReviewerGrade = grades[i];
                students[i].sortReviewerGrade = grades[i];
            }else if(grades[i]){
                students[i].sortReviewerGrade = -1;
            }
        }
        $scope.summaryData = students;
        $scope.rawData = students;
        $scope.personsCount = $scope.summaryData.length;
        $('.ui.dropdown').dropdown();//activates semantic drowpdowns
        $('.ui.accordion').accordion();
    });

    //Standard sort function
    $scope.sort = function (header, event) {
        var prop = header.sortPropertyName;
        var asc = header.asc;
        header.asc = !header.asc;
        var ascSort = function (a, b) {
            return a[prop] < b[prop] ? -1 : a[prop] > b[prop] ? 1 : 0;
        };
        var descSort = function (a, b) {
            return ascSort(b, a);
        };
        var sortFunc = asc ? ascSort : descSort;
        $scope.summaryData.sort(sortFunc);
    };

});