'use strict';
angular.module('Xtern')
    .controller('CompanyRecruiting', ['$scope', '$state', 'CompanyService', function ($scope, $state, CompanyService) {
        $scope.recruitmentList = [];

        $scope.companyData = getToken("organization");
        CompanyService.getOrganizationStudents(function (data) {
            $scope.recruitmentList = data;
        });

        $scope.sortableOptions = {
            containment: '#table-container',
            containerPositioning: 'relative'
        };

        $scope.removeRecruit = function (key) {
            console.log("remove recruit: " + key);
            CompanyService.removeStudentFromWishList(key, function (data) {
                for (var i = $scope.recruitmentList.length - 1; i >= 0; i--) {
                    if ($scope.recruitmentList[i].key == key) {
                        $scope.recruitmentList.splice(i, 1);
                    }
                }
            });

        };

        $scope.viewRecruit = function (key) {
            $state.go('company.profile', {key: key});
        };

        $scope.addStudent = function (key) {
            console.log("add student:");
            console.log(key);
        };

        $scope.dragControlListeners = {
            orderChanged: function (obj) {
                console.log(obj.source.index + ' ' + obj.dest.index);
                CompanyService.switchStudentsInWishList($scope.recruitmentList[obj.source.index].key, obj.dest.index, function (data) {
                    console.log("order changed: " + data);
                });
            }
        };
    }]);
