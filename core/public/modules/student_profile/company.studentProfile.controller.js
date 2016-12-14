angular.module('Xtern')
    .controller('CompanyStudentProfileCtrl', function($scope, $rootScope, $location, ProfileService, CompanyService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

    $scope.comment = {};
    $scope.isStudentApplicant = false;

    var isStudentApplicant = function(studentId) {
        return ($scope.companyData.studentIds.indexOf(parseInt(studentId)) != -1);
    };

    CompanyService.getCompanyDataForId(getToken("organization"), function(company) {
        $scope.companyData = company;
        $scope.isStudentApplicant = isStudentApplicant($stateParams._id);
    });

    $rootScope.$on('$stateChangeStart', function (event, toState, toParams, fromState, fromParams, options) {
        $scope.isStudentApplicant = isStudentApplicant($stateParams._id);
    });

    $('.ui.sticky').sticky({
        context: '#example1'
    });

    $(function () {
        $('.ui.dropdown').dropdown();
    });

    $scope.addComment = function(){
        // TODO: fix/update this for new data format
        var author_name = "controller test author";
        var group_name = "controller test group";
        var text = "controller test text bla bla bla. bla bla bla.";

        ProfileService.addCommentToStudent($scope.studentData._id, author_name, group_name, text, function (data) {
        });

        $scope.comment.author = 'test user'; //temporary
        $scope.comment.group = 'test users'; //temporary
        var newComment = angular.copy($scope.comment);
        $scope.studentData.comments.push(newComment);
    };

    $scope.removeComment = function(commentToRemove) {
        var author_name = "controller test author";
        var group_name = "controller test group";
        var text = "controller test text bla bla bla. bla bla bla.";

        ProfileService.removeCommentFromStudent($scope.studentData._id, author_name, group_name, text, function (data) {
        });

        // TODO: fix/update this for new data format
        for(var i = $scope.studentData.comments.length - 1; i >= 0; i--){
            if($scope.studentData.comments[i].text == text){
                $scope.studentData.comments.splice(i,1);
            }
        }
    };

    $scope.addStudent = function (_id) {
        $scope.isStudentApplicant=true;
        CompanyService.addStudentToWishList(_id, function(data) {
            toastr.success('Added Applicant', 'Student added to your Recruitment List');
        });
    };

    $scope.removeStudent = function (_id) {
        $scope.isStudentApplicant=false;
        CompanyService.removeStudentFromWishList(_id, function(data) {
            toastr.error('Removed Applicant', 'Student removed to your Recruitment List');
        });
    };
});