angular.module('Xtern')
    // .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', 'rzModule' , function ($scope, $rootScope, $state, AccountControlService, rzModule) {
    .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;
        $scope.showDecisionboard = true;
        $scope.showInstructorStats = false;
        $scope.companyList = [];

        $scope.activeStep = 'p1';


        $scope.phase1 = {
            fullList: [],
            list: [],
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
                },
                workStatus: {
                    labels: ['2016'],
                    data: [1],
                    name: 'Work Status'
                },
                ethnicity: {
                    labels: ['2016'],
                    data: [1],
                    name: 'Ethnicity'
                }
            },
            histogram: {
                labels: ['5', '6'],
                data: [[1, 2]],
                name: 'Histogram of Scores'
            }
        }

        $scope.stepClick = function (dest) {
            if (dest === 'p1') {
                $scope.showDecisionboard = true;
                $scope.showInstructorStats = false;
            } else if (dest === 'p2I') {
                $scope.showDecisionboard = false;
                $scope.showInstructorStats = true;
            }
        };

        var phase1FilterLoad = function () {
            $scope.phase1.list = $scope.phase1.fullList.filter(function (val) {
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
            var chartData3 = renderChartData($scope.phase1.list, 'workStatus');
            $scope.phase1.charts.workStatus.data = chartData3.values;
            $scope.phase1.charts.workStatus.labels = chartData3.keys;
            var chartData4 = renderChartData($scope.phase1.list, 'ethnicity');
            $scope.phase1.charts.ethnicity.data = chartData4.values;
            $scope.phase1.charts.ethnicity.labels = chartData4.keys;

        };

        var phase1HistLoad = function (metadata) {
            var histData = renderHistogramData(metadata, 'score');
            $scope.phase1.histogram.data = [histData.values];
            $scope.phase1.histogram.labels = histData.keys;
            console.log($scope.phase1.histogram);
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

        //Instructor Stats
        $scope.phase2 = {

        };

        $scope.phase2Instrutor = {
            studentHist: {
                labels: ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
                data: [[1, 6, 11, 12, 18, 14, 12, 12, 10, 1]],
                name: "Distribution of Student Scores"
            },
            reviewerHist: {
                labels: ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
                data: [[2, 6, 11, 15, 18, 16, 12, 10, 8, 1]],
                name: "Distribution of Reviewer Scores (AVG)"
            },
            progressPie: {
                labels: ['Completed', 'In Progress', 'Remaining'],
                data: [50, 23, 30],
                name: ""
            }
        };


        // Chart Util Funcitons
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


        var setup = function () {
            AccountControlService.getOrganizations(function (organizations) {
                $scope.companyList = organizations;
            });

            $scope.phase1.fullList = DECISION_BOARD_LIST;
            phase1HistLoad(DECISION_BOARD_LIST);
            phase1FilterLoad();

            $('.ui.sticky').sticky({context: '#processBoard', pushing: true});
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);