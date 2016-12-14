angular.module('Xtern')
    .controller('CompanyStudentProfileCtrl', function($scope, $location, ProfileService, CompanyService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

    $scope.comment = {};

    $('.ui.sticky').sticky({
        context: '#example1'
    });

    $(function () {
        $('.ui.dropdown').dropdown();
    });

    // $scope.selectStatus = function(option) {
    //     $scope.studentData.status = option;
    // };

    // $scope.selectR1Grade = function(option) {
    //     $scope.studentData.r1Grade = option;
    // };

    $scope.addComment = function(){
        // TODO: fix/update this for new data format
        var author_name = "controller test author";
        var group_name = "controller test group";
        var text = "controller test text bla bla bla. bla bla bla.";

        ProfileService.addCommentToStudent($scope.studentData._id, author_name, group_name, text, function (data) {
            // console.log(data);
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
            // console.log(data);git
        });

        // TODO: fix/update this for new data format
        for(var i = $scope.studentData.comments.length - 1; i >= 0; i--){
            if($scope.studentData.comments[i].text == text){
                $scope.studentData.comments.splice(i,1);
            }
        }
    };

    $scope.addStudent = function (_id) {
        console.log("add student:");
        console.log(_id);
        CompanyService.addStudentToWishList(_id, function(data) {
            // $scope.recruitmentList.push($scope.studentData);
            console.log("Student added");
        });

    };
});