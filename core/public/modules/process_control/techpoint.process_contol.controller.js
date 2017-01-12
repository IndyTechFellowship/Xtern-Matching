angular.module('Xtern')
    // .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', 'rzModule' , function ($scope, $rootScope, $state, AccountControlService, rzModule) {
    .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;
        $scope.phase1 = {
            fullList: [],
            list:[],
            displayList: [],
            slider: {
                value: 7.5,
                options: {
                    id: 'first',
                    floor: 0,
                    ceil: 10,
                    step: 0.5,
                    precision: 1,
                    onChange: function (sliderId, modelValue, highValue, pointerType) {
                        phase1FilterLoad()
                        console.log(sliderId, modelValue, highValue, pointerType);
                    }
                }
            },
            charts: {
                gender: {
                    labels: ['Male'],
                    data: [70],
                    name: 'Gender'
                },
                class: {
                    labels: ['2016'],
                    data: [1],
                    name: 'Class Year'
                }
            },
            histogram: {
                labels: ['5','6'],
                data: [[1,2]],
                name: 'Histogram of Scores'
            }
        }

        $scope.companyList = [];


        var phase1FilterLoad = function () {
            $scope.phase1.list = $scope.phase1.fullList.filter(function(val){
                return val["score"] >= $scope.phase1.slider.value;
            });
            // console.log(list);
            // $scope.phase1.list = $scope.phase1.fullList;
            $scope.phase1.displayList = splitInToTwo($scope.phase1.list);
            var chartData = renderChartData($scope.phase1.list, 'gender');
            $scope.phase1.charts.gender.data = chartData.values;
            $scope.phase1.charts.gender.labels = chartData.keys;
            var chartData2 = renderChartData($scope.phase1.list, 'gradYear');
            $scope.phase1.charts.class.data = chartData2.values;
            $scope.phase1.charts.class.labels = chartData2.keys;

        };

        var phase1HistLoad = function(metadata){
            var histData = renderHistogramData(metadata, 'score');
            $scope.phase1.histogram.data = [histData.values];
            $scope.phase1.histogram.labels = histData.keys;
        }

        var splitInToTwo = function (inList) {
            var tempArr = [];
            for (i = 0; i < inList.length; i += 2) {
                tempArr.push({
                    right: inList[i],
                    left: i + 1 < inList.length ? inList[i + 1] : '-'
                });
            }
            return tempArr;
        };

        var renderHistogramData = function (data, field) {
            var _map = [];
            //create map
            for (var i in data) {
                var itemVal = Math.ceil(data[i][field]);
                if (_map[itemVal]) {
                    _map[itemVal]++;
                } else {
                    _map[itemVal] = 1;
                }
            };
            var output = { keys: [], values: [] };
            for (key in _map) {
                output.keys.push(key);
                output.values.push(_map[key]);
            };
            return output;
        };



        var renderChartData = function (data, field) {
            var _map = [];
            //create map
            for (var i in data) {
                var itemVal = data[i][field];
                if (_map[itemVal]) {
                    _map[itemVal]++;
                } else {
                    _map[itemVal] = 1;
                }
            };
            var output = { keys: [], values: [] };
            for (key in _map) {
                output.keys.push(key);
                output.values.push(_map[key]);
            };
            return output;
        };

        var generateHeaders = function (field, data, array) {
            array.length = 0;
            for (var rowIndex in data) {
                if (array.indexOf(data[rowIndex][field]) === -1) {
                    array.push(data[rowIndex][field]);
                }
            }
            array.sort();
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


        var setup = function () {
            AccountControlService.getOrganizations(function (organizations) {
                $scope.companyList = organizations;
            });
                      
            $scope.phase1.fullList = DECISION_BOARD_LIST;
            phase1HistLoad(DECISION_BOARD_LIST);
            phase1FilterLoad();

        };

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);