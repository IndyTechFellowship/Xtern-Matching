'use strict';
angular.module('Xtern')
.controller('CompanyRecruiting', ['$scope', '$rootScope', '$state', 'CompanyService', function ($rootScope, $scope, $state, CompanyService) {
    $scope.recruitmentList = [];

    CompanyService.getOrganizationStudents(function(data) {
        console.log("ORG Students: ", data);
        $scope.recruitmentList = data;
    });

    $scope.sortableOptions = {
        containment: '#table-container',
        containerPositioning: 'relative'
    };

    $scope.removeRecruit = function (key) {
        console.log("remove recruit: "+key);
        CompanyService.removeStudentFromWishList(key, function(data) {
            for (var i = $scope.recruitmentList.length - 1; i >= 0; i--) {
                if ($scope.recruitmentList[i].key == key) {
                    $scope.recruitmentList.splice(i, 1);
                }
            }
        });

    };

    $scope.viewRecruit = function (key) {
        $state.go('company.profile', { key: key });
    };

    $scope.addStudent = function (key) {
        console.log("add student:");
        console.log(key);
    };

    $scope.dragControlListeners = {
        orderChanged: function(obj) {
            console.log("new index", obj.dest.index);
            console.log(obj.source.index+' '+obj.dest.index);
            CompanyService.switchStudentsInWishList($scope.recruitmentList[obj.source.index].key, $scope.recruitmentList[obj.dest.index].key, function(data) {
                console.log("order changed: ", data);
            });
        }
    };
}]);