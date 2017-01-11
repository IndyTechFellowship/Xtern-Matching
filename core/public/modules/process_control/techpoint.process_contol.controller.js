angular.module('Xtern')
    // .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', 'rzModule' , function ($scope, $rootScope, $state, AccountControlService, rzModule) {
        .controller('TechPointProcessControl', ['$scope', '$rootScope', '$state', 'AccountControlService', function ($scope, $rootScope, $state, AccountControlService) {
        var self = this;
        $scope.phase1 = {
            list: [],
            displayList:[],
            slider:{
                value: 7.5,
                options: {
                    id:'first',
                    floor:0,
                    ceil:10,
                    step: 0.5,
                    precision: 1,
                    onChange: function(sliderId, modelValue, highValue, pointerType){
                        console.log(sliderId, modelValue, highValue, pointerType);
                    }
                }
            }, charts:{
                gender:{
                    labels:['Male', 'Female'],
                    data: [70,60],
                    name:'Gender'
                },
                class:{
                    labels:['2016', '2015', '2014'],
                    data: [70,60,110],
                    name:'Class Year'
                }
            }
        }

        $scope.companyList = [];


        var phase1Load = function(){
            $scope.phase1.list = DECISION_BOARD_LIST;
            $scope.phase1.displayList = splitInToTwo($scope.phase1.list);
        };

        var splitInToTwo = function(inList){
            var tempArr =[];
            for(i = 0; i<inList.length; i+=2){
                tempArr.push({
                    right: inList[i],
                    left: i + 1 < inList.length? inList[i+1] : '-'
                });
            }
            return tempArr;
        };

        var setup = function () {
            AccountControlService.getOrganizations(function (organizations) {
                $scope.companyList = organizations;                               
            });
            phase1Load();
            
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            setup();
        });
    }]);