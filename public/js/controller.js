angular.module('Xtern')
    .controller('StudentProfileCtrl', function($scope, $location, ProfileService, $stateParams) {
        $('.ui.dropdown').dropdown();//activites semantic dropdowns

        $scope.comment = {};
        $scope.studentData = null;        

        ProfileService.getStudentDataForId($stateParams._id, function(data)            
        {
            $scope.studentData = data;
        });

        $scope.statusOptions = [
            'Stage 1 Approved',
            'Stage 2 Approved',
            'Stage 3 Approved',
            'Undecided',
            'Rejected (Stage 1)',
            'Rejected (Stage 2)',
            'Rejected (Stage 3)'
        ];

        $scope.r1GradeOptions = [
            {
                "text":"A",
                "value":4
            },{
                "text":"B+",
                "value":3.5
            },{
                "text":"B",
                "value":3
            },{
                "text":"B-",
                "value":2.8
            },{
                "text":"C",
                "value":2
            },{
                "text":"D",
                "value":1
            }];

        $('.ui.sticky').sticky({
            context: '#example1'
        });

        $(function () {
            $('.ui.dropdown').dropdown();
        });

        $scope.selectStatus = function(option) {
            $scope.studentData.status = option;
        };

        $scope.selectR1Grade = function(option) {
            $scope.studentData.r1Grade = option;
        };

        $scope.addComment = function(){
            $scope.comment.author = 'test user'; //temporary
            $scope.comment.group = 'test users'; //temporary
            var newComment = angular.copy($scope.comment);
            $scope.studentData.comments.push(newComment);
        };

        $scope.removeComment = function(commentToRemove) {
            for(var i = $scope.studentData.comments.length - 1; i >= 0; i--){
                if($scope.studentData.comments[i].text == commentToRemove.text){
                    $scope.studentData.comments.splice(i,1);
                }
            }
        };
    })
    .controller('CommentCtrl', function() {
        $('.dimmable.card').dimmer({
            on: 'hover'
        });
    })
    .controller('TechLabelsCtrl', function($scope) {
        $scope.colorForLanguage = function(category) {
            if(category === "Full-Stack") {
                return 'red';
            } else if (category === "Front-End") {
                return 'green';
            } else if (category === "Mobile") {
                return 'blue';
            } else if (category === "General") {
                return 'yellow';
            } else if (category === "Database") {
                return 'purple';
            } else {
                return 'black';
            }
        };
    })
    .controller('MissionControl', ['$scope', '$state', function ($scope,$state) {
        //Params from the parent
        var STARTCHARTSANDSTATS = $scope.startchartsandchats;
        var STARTFILTERS = $scope.startfilters;
        var TABLEHEADERS = $scope.tableheaders;
        var PATH = $scope.path;
        
        //DataLoad
        $scope.$watch('data', function (newVal, oldval) {
            if (newVal)
                dataLoad(newVal);
        }, true);

        var dataLoad = function (data) {
            $scope.summaryData = $.map(data, function (person) {
                return rowClass(person)
            });
            $scope.rawData = $.map(data, function (person) {
                return rowClass(person)
            });
            loadFilters();
            initCharts();
            $scope.personsCount = $scope.summaryData.length;

            $('.ui.dropdown').dropdown();//activates semantic drowpdowns
        };

        //vars
        $scope.summaryData = null;
        $scope.rawData = null;
        $scope.personsCount = 0;


        //Graph Stuff
        var generateChartAndStatus = function (jsonArray) {
            $scope.chartsAndStats = [];
            for (var key in jsonArray) {
                var chartStat = jsonArray[key];
                if (chartStat.isChart) {
                    $scope.chartsAndStats.push({
                        chart: chartStat.isChart,
                        title: chartStat.title,
                        labels: chartStat.labels,
                        dataLabel: chartStat.dataLabel,
                        data: [],
                        init: chartStat.labels.length > 0 ?
                            function (data) {
                            } :
                            chartStat.nestedData ? function (data) {
                                generateHeadersNestedArray(this.dataLabel, data, this.labels);
                            } :
                                function (data) {
                                    generateHeaders(this.dataLabel, data, this.labels);
                                },
                        refresh: chartStat.nestedData ? function (data) {
                            this.data = generateNestedChartData(this.labels, data, this.dataLabel)
                        } :
                            function (data) {
                                this.data = generateChartData(this.labels, data, this.dataLabel)
                            }
                    });
                }
                else {
                    $scope.chartsAndStats.push({
                        chart: chartStat.isChart,
                        title: chartStat.title,
                        icon: chartStat.icon,
                        data: 0,
                        dataLabel: chartStat.dataLabel,
                        nestedData: chartStat.nestedData,
                        uniqueObjects: [],
                        init: function (data) {
                        },
                        refresh: chartStat.nestedData ? function (data) {
                            generateHeadersNestedArray(this.dataLabel, data, this.uniqueObjects);
                            this.data = this.uniqueObjects.length;
                        } :
                            function (data) {
                                generateHeaders(this.dataLabel, data, this.uniqueObjects);
                                this.data = this.uniqueObjects.length;
                            }
                    });
                }
            }
        };
        generateChartAndStatus(STARTCHARTSANDSTATS);

        var initCharts = function () {
            for (var stat in $scope.chartsAndStats) {
                $scope.chartsAndStats[stat].init($scope.summaryData);
            }
            refreshCharts();
        };

        var generateChartData = function (headers, data, dataLabel) {
            //initialize array
            var returnData = Array.apply(null, Array(headers.length)).map(function () {
                return 0
            });
            for (var rowIndex in data) {
                var value = data[rowIndex][dataLabel];
                for (var i = 0; i < headers.length; i++) {
                    if (headers[i] === value) {
                        returnData[i]++;
                        // break;
                    }
                }
            }
            return returnData;
        };

        var generateNestedChartData = function (headers, data, dataLabel) {
            //initialize array
            var returnData = Array.apply(null, Array(headers.length)).map(function () {
                return 0
            });
            for (var rowIndex in data) {
                for (var arrayIndex in data[rowIndex][dataLabel]) {
                    var value = data[rowIndex][dataLabel][arrayIndex];
                    for (var i = 0; i < headers.length; i++) {
                        if (headers[i] === value) {
                            returnData[i]++;
                            // break;
                        }
                    }
                }
            }
            return returnData;
        };

        var refreshCharts = function () {
            for (var stat in $scope.chartsAndStats) {
                $scope.chartsAndStats[stat].refresh($scope.summaryData);
            }
        };

        //Table Stuff
        $scope.tableHeaders = [];
        var setTableHeaders = function (jsonArray) {
            $scope.tableHeaders = jsonArray;
        };
        setTableHeaders(TABLEHEADERS);
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

        // FILTERS
        var generateFilterObjects = function (jsonArray) {
            $scope.filterObjects = [];
            for (var key in jsonArray) {
                var filter = jsonArray[key];
                $scope.filterObjects.push({
                    isToggle: filter.isToggle,
                    all: false,
                    label: filter.label,
                    dataLabel: filter.dataLabel,
                    filterList: [],
                    optionsList: [],
                    filterFunc: filter.simpleFilter ?
                        function (row) {
                            return isContained(row, this.filterList, this.dataLabel);
                        }
                        : function (row) {
                        return complexFilter(row, this.filterList, this.all, this.dataLabel);
                    },
                    generate: filter.simpleFilter ?
                        function (data) {
                            generateHeaders(this.dataLabel, data, this.optionsList);
                        }
                        : function (data) {
                        return generateHeadersNestedArray(this.dataLabel, data, this.optionsList);
                    }
                });
            }
        };
        generateFilterObjects(STARTFILTERS);

        //Filter Helper Functions
        var generateHeaders = function (field, data, array) {
            array.length = 0;
            for (var rowIndex in data) {
                if (array.indexOf(data[rowIndex][field]) === -1) {
                    array.push(data[rowIndex][field]);
                }
            }
            array.sort();
        };
        var generateHeadersNestedArray = function (field, data, array) {
            array.length = 0;
            for (var rowIndex in data) {
                for (var arrayIndex in data[rowIndex][field]) {
                    if (array.indexOf(data[rowIndex][field][arrayIndex]) === -1) {
                        array.push(data[rowIndex][field][arrayIndex]);
                    }
                }
            }
            array.sort();
        };
        var isContained = function (row, selectedOptions, prop) {
            var instance = "" + row[prop];
            var propArray = selectedOptions;
            return !propArray || propArray.length == 0 || $.inArray(instance, propArray) > -1;//contained
        };
        var complexFilter = function (row, array, all, field) {
            if (all)
                return containsAllSelected(row, array, field);
            else
                return containsAtLeastOne(row, array, field);
        };
        var containsAtLeastOne = function (row, selectedOptions, prop) {
            if (!selectedOptions || selectedOptions.length == 0) //no filters exist
                return true;
            for (var index in row[prop]) {
                if ($.inArray(row[prop][index], selectedOptions) > -1) {
                    return true;
                }
            }
            return false;
        };
        var containsAllSelected = function (row, selectedOptions, prop) {
            if (!selectedOptions || selectedOptions.length == 0) //no filters exist
                return true;
            for (var index in selectedOptions) {
                if ($.inArray(selectedOptions[index], row[prop]) < 0) {
                    return false;
                }
            }
            return true;
        };
        //Execute Filter Functions
        var loadFilters = function () {
            for (var filter in $scope.filterObjects) {
                $scope.filterObjects[filter].generate($scope.summaryData);
            }

        };
        var isInFilters = function (row) {
            var vaild = true;
            for (var filter in $scope.filterObjects) {
                vaild = vaild && $scope.filterObjects[filter].filterFunc(row);
            }
            return vaild;
        };

        $scope.filterTable = function () {
            $scope.summaryData = [];
            for (var index in $scope.rawData) {
                if (isInFilters($scope.rawData[index]))
                    $scope.summaryData.push($scope.rawData[index]);
            }
            refreshCharts();
            $scope.personsCount = $scope.summaryData.length;
        };
        $scope.clearAllFilters = function () {
            $scope.filterData.Selected.Universities.length = 0;
            $scope.filterData.Selected.Technologies.length = 0;
            $scope.filterData.Selected.Name.length = 0;
            $scope.filterData.Selected.Major.length = 0;
            $scope.filterData.Selected.WorkStatus.length = 0;
            $scope.filterData.Selected.Interests.length = 0;
            $scope.filterData.Selected.Status.length = 0;
            $scope.filterTable();
        };

        //DataLoad
        //Table Click
        $scope.rowClick = function(id){
            $state.go(PATH +'.profile', {_id:id});            
        }
        //DOM
        $('.ui.accordion')
            .accordion();
    }]);
