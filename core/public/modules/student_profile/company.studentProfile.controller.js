angular.module('Xtern')
    .controller('CompanyStudentProfileCtrl', function($scope, $location, ProfileService, CompanyService, $stateParams) {
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

        $scope.addStudent = function (key) {
            CompanyService.addStudentToWishList(key, function(data) {
                // $scope.recruitmentList.push($scope.studentData);
                console.log("Student added");
            });
        };

        $scope.$on('$viewContentLoaded', function (evt) {
            $('.ui.dropdown').dropdown();
            $('.ui.sticky').sticky({
                context: '#example1'
            });
        });
});