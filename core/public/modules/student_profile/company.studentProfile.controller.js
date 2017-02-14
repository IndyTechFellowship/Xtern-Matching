'use strict';
angular.module('Xtern')
.controller('CompanyStudentProfileCtrl', function($scope, $rootScope, $location, ProfileService, CompanyService, $stateParams) {
    // var toastr = require('toastr/build/toastr.min.js');
    $scope.comment = {};
    $scope.isStudentApplicant = false;
    $scope.companyData = {};
    $scope.companyData.students = new Array();
    $scope.studentKey = $stateParams.key;

    CompanyService.getOrganizationCurrentFromLogin(function(company) {
        $scope.companyData = company;
        // console.log("company data in recruiting controller:", $scope);
        ProfileService.getStudentDataForIds(company.students, function(data) {
        });
        $scope.isStudentApplicant = isStudentApplicant($scope.studentKey);
    });

    var isStudentApplicant = function(studentId) {
        if($scope.companyData.students == {} || $scope.companyData.students === null) {
            CompanyService.getOrganizationCurrentFromLogin(function(company) {
                $scope.companyData = company;
                $scope.companyData.students = company.students;
                // console.log("company data in recruiting controller:", $scope);
                ProfileService.getStudentDataForIds(company.students, function(data) {
                });
                $scope.isStudentApplicant = isStudentApplicant($scope.studentKey);
            });
        } else {
            // console.log("STUDENTS", $scope.companyData.students);
            // console.log("STUDENT_ID", studentId);
            // console.log("INDEX", $scope.companyData.students.indexOf(studentId +""));
            return ($scope.companyData.students.indexOf(studentId+"") != -1);
        }
    };

    $rootScope.$on('$stateChangeStart', function (event, toState, toParams, fromState, fromParams, options) {
        // console.log("STATE_PARAM_ID", $stateParams);
        $scope.isStudentApplicant = isStudentApplicant($scope.studentKey);
    });

    $scope.addStudent = function () {
        CompanyService.addStudentToWishList($stateParams.key, function(data) {
                // $scope.recruitmentList.push($scope.studentData);
                $scope.isStudentApplicant = true;
                // toastr.success('Added Applicant', 'Student added to your Recruitment List');
            });
    };

    $scope.removeStudent = function () {
        CompanyService.removeStudentFromWishList($stateParams.key, function(data) {
            $scope.isStudentApplicant=false;
            // toastr.error('Removed Applicant', 'Student removed to your Recruitment List');
        });
    };

    $scope.$on('$viewContentLoaded', function (evt) {
        $scope.studentKey = $stateParams.key;
        $('.ui.dropdown').dropdown();
        $('.ui.sticky').sticky({
            context: '#example1'
        });
        $scope.isStudentApplicant = isStudentApplicant($scope.studentKey);
    });
});