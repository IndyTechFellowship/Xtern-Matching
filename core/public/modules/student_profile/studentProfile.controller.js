angular.module('Xtern')
    .controller('StudentProfileCtrl', function($scope, $location, ProfileService, CompanyService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

    $scope.comment = {};
    $scope.isCompany = getToken('isCompany');

    // TODO: turn this into point grade
    $scope.statusOptions = [
        'Stage 1 Approved',
        'Stage 2 Approved',
        'Stage 3 Approved',
        'Undecided',
        'Rejected (Stage 1)',
        'Rejected (Stage 2)',
        'Rejected (Stage 3)'
    ];

    $scope.r1GradeOptions = [1,2,3,4,5,6,7,8,9,10];

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

})