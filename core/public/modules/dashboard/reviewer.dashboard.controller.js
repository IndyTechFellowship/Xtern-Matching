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


            //Params from the parent
        var STARTCHARTSANDSTATS = $scope.startchartsandchats;
        var STARTFILTERS = $scope.startfilters;
        var TABLEHEADERS = $scope.tableheaders;
        var PATH = $scope.path;

        //vars
        $scope.summaryData = null;
        $scope.rawData = null;
        $scope.personsCount = 0;


        //DataLoad
        // $scope.$watch('data', function (newVal, oldval) {
        //     if (newVal)
        //         dataLoad(newVal);
        // }, true);

        var dataLoad = function (data, keys) {
            // $scope.summaryData = $.map(data, function (person) {
            //     return rowClass(person)
            // });
            // $scope.rawData = $.map(data, function (person) {
            //     return rowClass(person)
            // });
            let students = [];
            for(var i = 0; i < data.length; i++) {
                students[i] = rowClass(data[i],keys[i]);
            }
            $scope.summaryData = students;
            $scope.rawData = students;
            // loadFilters();
            // initCharts();
            $scope.personsCount = $scope.summaryData.length;

            $('.ui.dropdown').dropdown();//activates semantic drowpdowns
        };

                //Table Stuff
        $scope.tableHeaders = [];
        var setTableHeaders = function (jsonArray) {
            $scope.tableHeaders = jsonArray;
        };

        $scope.sort = function (header) {
            var prop = header.sortPropertyName;
            var asc = header.asc === 'ascending';
            header.asc = asc ? 'descending' : 'ascending';
            $scope.tableHeaders.forEach(function (header) {
                header.selected = false;
            });
            header.selected = true;
            var ascSort = function (a, b) {
                return a[prop] < b[prop] ? -1 : a[prop] > b[prop] ? 1 : 0;
            };
            var descSort = function (a, b) {
                return ascSort(b, a);
            };
            var sortFunc = asc ? ascSort : descSort;
            $scope.summaryData.sort(sortFunc);
        };

                //Table Click
        $scope.rowClick = function (key) {
            $state.go(PATH + '.profile', {key: key});
        };

        var run = function(data, keys){
            // generateChartAndStatus(STARTCHARTSANDSTATS);
            setTableHeaders(TABLEHEADERS);            
            // generateFilterObjects(STARTFILTERS);
            dataLoad(data, keys);

            //DOM
            $('.ui.accordion').accordion();
        };



        TechPointDashboardService.queryUserSummaryData(function (data, keys) {
            $scope.DATA = data;
            //$scope.KEYS = keys;
            run(data,keys);
        });
}]);