angular.module('Xtern')
    .controller('ReviewerStudentProfileCtrl', function($scope, $location, ProfileService, ReviewerProfileService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

    $scope.comment = {};
    $scope.reviewerGrade = 0;
    $scope.reviewerGradeOptions = [1,2,3,4,5,6,7,8,9,10];

    $scope.selectReviewerGrade = function(option) {
        $scope.reviewerGrade = option;

        ReviewerProfileService.postReviewerGradeForStudent($stateParams.key, option);
    };

    // $scope.addComment = function(){
    //     // TODO: fix/update this for new data format
    //     var text = "controller test text bla bla bla. bla bla bla.";

    //     ProfileService.addCommentToStudent(text, function (comment) {
    //         $scope.comment.authorName = comment.author; //temporary
    //         var newComment = angular.copy($scope.comment);
    //         $scope.studentData.comments.push(newComment);
    //     });
    // };

    // $scope.removeComment = function(commentToRemove) {
    //     var author_name = "controller test author";
    //     var group_name = "controller test group";
    //     var text = "controller test text bla bla bla. bla bla bla.";

    //     ProfileService.removeCommentFromStudent($scope.studentData._id, author_name, group_name, text, function (data) {
    //         // console.log(data);git
    //     });

    //     // TODO: fix/update this for new data format
    //     for(var i = $scope.studentData.comments.length - 1; i >= 0; i--){
    //         if($scope.studentData.comments[i].text == text){
    //             $scope.studentData.comments.splice(i,1);
    //         }
    //     }
    // };

    var ReviewerStudentProfileCtrlSetup = function(){
        $('.ui.sticky').sticky({
            context: '#example1'
        });
        
        $('.ui.dropdown').dropdown();        
    };


    $scope.$on('$viewContentLoaded', function (evt) {
        console.log("SCOPE");
        console.log($scope);
        ReviewerStudentProfileCtrlSetup();
        ReviewerProfileService.getReviewerGradeForStudent($stateParams.key, function (reviewerGrade) {
            console.log($stateParams.key);
            console.log("REVIEWER GRADE", reviewerGrade);
            $scope.reviewerGrade = reviewerGrade;
        });
        // console.log("reviewergrades from json", $stateParams.key);
    });
});