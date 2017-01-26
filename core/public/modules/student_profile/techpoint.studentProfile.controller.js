angular.module('Xtern')
    .controller('TechPointStudentProfileCtrl', function($scope, $location, ProfileService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

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
    $scope.comments = [];
    $scope.commentMessage = "";

    $scope.addComment = function(studentKey) {
        if($scope.commentMessage === "") {
            //TODO do something
        } else {
            ProfileService.addComment(studentKey,$scope.commentMessage, function (comment, err) {
                if(!err) {
                    $scope.comments.push(comment);
                    $scope.commentMessage = "";
                } else {
                    //err
                }
            });
        }
    };

    $scope.removeComment = function(key) {
        ProfileService.removeComment(key, function (data, err) {
            if(err) {
                //error
            } else {
                //remove comment to the active scope
                $scope.comments = $scope.comments.filter(function (comment) {
                    return comment.key === key;
                });
            }
        });
    };

    $scope.selectStatus = function(option) {
        $scope.studentData.status = option;
    };

    $scope.selectR1Grade = function(option) {
        $scope.studentData.r1Grade = option;
    };

    $scope.$on('$viewContentLoaded', function (evt) {
        $('.ui.sticky').sticky({
            context: '#example1'
        });
        $('.ui.dropdown').dropdown();
        $scope.studentKey = $stateParams.key;
        console.log("state-Key: ",$stateParams.key);
        ProfileService.getComments($scope.studentKey,function (comments,err) {
            if(err) {
                console.log(err);
            } else {
                $scope.comments = comments;
            }
        });
    });

});