angular.module('Xtern')
    // .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', 'rzModule' , function ($scope, $rootScope, $state, AccountControlService, rzModule) {
    .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService','DecisionBoardService', 'TechPointReviewerControlService', 'CompanyService',function ($scope, $rootScope, $state, AccountControlService,DecisionBoardService,TechPointReviewerControlService, CompanyService) {
        var self = this;
        $scope.current = "p1";
        $scope.company = {
            active: "",
            list:[],
            studentList:[],
            changeCompany:function(){
                $scope.company.studentList=[];
                //console.log($scope.company.active)
                CompanyService.getOrganizationStudentsWithKey($scope.company.active, function(arr){
                    $scope.company.studentList = arr;
                });
            }
        };

        $scope.viewRecruit = function (key) {
            $state.go('techpoint.profile', { key: key });
        };

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
                labels: ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
                data: [[2, 6, 11, 15, 18, 16, 12, 10, 8, 1]],
                name: 'Distribution of Student Scores'
            }
        };

        $scope.createReviewGroups = function() {
            TechPointReviewerControlService.createReviewGroups(20, 3, function(data) {});
        };

        $scope.stepClick = function (dest) {
            $scope.current=dest;
        };

        var phase1FilterLoad = function () {
            $scope.phase1.list = $scope.phase1.fullList.filter(function (val) {
                return val["grade"] >= $scope.phase1.slider.value;
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

        var phase2FilterLoad = function () {
            $scope.phase2.list = $scope.phase2.fullList.filter(function (val) {
                return val["grade"] >= $scope.phase2.slider.value;
            });
            // console.log(list);
            // $scope.phase1.list = $scope.phase1.fullList;
            $scope.phase2.displayList = splitInToTwo($scope.phase2.list);
            var chartData = renderChartData($scope.phase2.list, 'gender');
            $scope.phase2.charts.gender.data = chartData.values;
            $scope.phase2.charts.gender.labels = chartData.keys;
            var chartData2 = renderChartData($scope.phase2.list, 'gradYear');
            $scope.phase2.charts.class.data = chartData2.values;
            $scope.phase2.charts.class.labels = chartData2.keys;
            var chartData3 = renderChartData($scope.phase2.list, 'workStatus');
            $scope.phase2.charts.workStatus.data = chartData3.values;
            $scope.phase2.charts.workStatus.labels = chartData3.keys;
            var chartData4 = renderChartData($scope.phase2.list, 'ethnicity');
            $scope.phase2.charts.ethnicity.data = chartData4.values;
            $scope.phase2.charts.ethnicity.labels = chartData4.keys;
        };

        var phase1HistLoad = function (metadata) {
            var histData = renderHistogramData(metadata, 'grade');
            $scope.phase1.histogram.data = [histData.values];
            $scope.phase1.histogram.labels = histData.keys;
        };

        var phase2HistLoad = function (metadata) {
            var studentHistData = renderHistogramData(metadata, 'grade');
            $scope.phase2.student_histogram.data = [studentHistData.values];
            $scope.phase2.student_histogram.labels = studentHistData.keys;
            var reviewerHist = renderReviewerHistData($scope.phase2.instructors);
            $scope.phase2.reviewer_histogram.data = [reviewerHist.values];
            $scope.phase2.reviewer_histogram.labels = reviewerHist.keys;
        };

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
            fullList: [],
            list: [],
            displayList: [],
            instructors: {},
            remaining: 0,
            slider: {
                value: 7.5,
                options: {
                    id: 'first',
                    floor: 0,
                    ceil: 10,
                    step: 0.5,
                    precision: 1,
                    onChange: function (sliderId, modelValue, highValue, pointerType) {
                        phase2FilterLoad()
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
            student_histogram: {
                labels: ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
                data: [[1, 6, 11, 12, 18, 14, 12, 12, 10, 1]],
                name: 'Distribution of Student Scores (AVG)'
            },
            reviewer_histogram: {
                labels: ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
                data: [[1, 6, 11, 12, 18, 14, 12, 12, 10, 1]],
                name: 'Distribution of Reviewer Scores (AVG)'
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
            }
            var output = { keys: [], values: [] };
            for (key in _map) {
                output.keys.push(key);
                output.values.push(_map[key]);
            }
            return output;
        };

        var renderReviewerHistData = function (data) {
            var _map = [];
            //create map
            for (var i in data) {
                var score = 0;
                for(var j in data[i]){
                    score+= data[i][j];
                }
                var itemVal = Math.ceil(score/data[i].length);
                if (_map[itemVal]) {
                    _map[itemVal]++;
                } else {
                    _map[itemVal] = 1;
                }
            }
            var output = { keys: [], values: [] };
            for (key in _map) {
                output.keys.push(key);
                output.values.push(_map[key]);
            }
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
            }
            var output = { keys: [], values: [] };
            for (key in _map) {
                output.keys.push(key);
                output.values.push(_map[key]);
            }
            return output;
        };

        $scope.refreshPhaseOne = function(hardReload){
            DecisionBoardService.getPhaseOne(function(list){
                $scope.phase1.fullList = list;
                phase1HistLoad(list);
                phase1FilterLoad();
            },hardReload);
        };

        var phase2datascrub = function(list){
            list.forEach(function(student){
                var tempGrade = 0;
                var reviewCount = 0;
                for(var i in student.reviewerGrades){
                    var review = student.reviewerGrades[i];
                    tempGrade += parseInt(review.grade);
                    reviewCount ++;
                    if($scope.phase2.instructors[review.reviewer]){
                        $scope.phase2.instructors[review.reviewer].push(review.grade);
                    }
                    else{
                        $scope.phase2.instructors[review.reviewer] = [review.grade];
                    }
                }
                student.grade = tempGrade/reviewCount;
            });
            $scope.phase2.fullList = list;
        };

        $scope.refreshPhaseTwo = function(hardReload){
            DecisionBoardService.getPhaseTwo(function(list){
                if(list){
                    phase2datascrub(list);
                    phase2HistLoad($scope.phase2.fullList);
                    phase2FilterLoad();
                }
            },hardReload);

            DecisionBoardService.getReviewerOverview(function(data){
               //{noReviews: 170, someReviews: 30, allReviews: 0}allReviews: 0noReviews: 170someReviews: 30__proto__: Object
                $scope.phase2.reviewer_histogram.data = [data.allReviews, data.someReviews, data.noReviews];
                $scope.phase2.remaining = data.someReviews + data.noReviews;

            });
        };

        var setup = function () {
            AccountControlService.getOrganizations(function (organizations) {
                $scope.company.list = organizations.filter(function(org){
                    return org.name != 'Reviewers' && org.name !== 'Techpoint';
                });
            });
            $scope.refreshPhaseOne(false);
            $scope.refreshPhaseTwo(false);
            $('.ui.sticky').sticky({context: '#processBoard', pushing: true});
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);