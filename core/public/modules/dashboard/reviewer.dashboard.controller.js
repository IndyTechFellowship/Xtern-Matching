angular.module('Xtern')
    .controller('ReviewerDashboardCtrl', ['$scope', 'TechPointDashboardService', function($scope, TechPointDashboardService){
    //BEGIN CONFIG DATA
    $scope.STARTCHARTSANDSTATS = {
        University: {
            isChart:false,
            title: "Universities",
            icon:'university',
            dataLabel: 'university',
            nestedData: false
        },
    };
    $scope.STARTFILTERS = {
        Grade: {
            isToggle: false,
            label: "Grade",
            dataLabel: 'gradeLabel',
            simpleFilter:true,
            nestedHeaders:true,
        }
    };
    $scope.TABLEHEADERS = [
        {
            title: 'Name',
            displayPropertyName: 'name',
            sortPropertyName: 'name',
            sort: 'ascending',
            selected: false
        },
        {
            title: 'Grade',
            displayPropertyName: 'gradeLabel',
            sortPropertyName: 'gradeValue',
            sort: 'descending',
            selected: false
        }];
    $scope.DATA = null;
    $scope.PATH ='reviewer';
    $scope.isCompany = false;
}]);