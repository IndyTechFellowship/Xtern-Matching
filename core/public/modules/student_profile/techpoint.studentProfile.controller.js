angular.module('Xtern')
    .controller('TechPointStudentProfileCtrl', function($scope, $location, ProfileService, $stateParams) {
        $('.ui.dropdown').dropdown();

        $scope.statusOptions = [
            'Stage 1 Approved',
            'Stage 2 Approved',
            'Stage 3 Approved',
            'Undecided',
            'Rejected (Stage 1)',
            'Rejected (Stage 2)',
            'Rejected (Stage 3)'
        ];
        $scope.r1GradeOptions = [];
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

                    }
                });
            }
        };

        $scope.removeComment = function(key) {
            ProfileService.removeComment(key, function (err) {
                if(err) {
                    //error
                } else {
                    //remove comment to the active scope
                    $scope.comments = $scope.comments.filter(function (comment) {
                        return comment.key !== key;
                    });
                }
            });
        };

        $scope.selectStatus = function(option) {
            ProfileService.setStatus($scope.studentKey,option, function (err) {
                if(err) {}
            });
        };

        $scope.selectR1Grade = function(option) {
            ProfileService.setR1Grade($scope.studentKey,option, function (err) {
                if(err) {}
            });
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            $scope.studentKey = $stateParams.key;
            ProfileService.getComments($scope.studentKey, function (comments,err) {
                if(err) {
                    //error
                } else {
                    $scope.comments = comments;
                }
            });
            $('.ui.sticky').sticky({
                context: '#example1'
            });
            $('.ui.dropdown').dropdown();
            for(let i=0; i < 10; i += 0.5) {
                $scope.r1GradeOptions.push(i);
            }
        });
});