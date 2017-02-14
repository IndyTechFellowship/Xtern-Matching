angular.module('Xtern')
    .controller('CommentCtrl', function($scope, ProfileService) {
        $scope.comments = [];
        $scope.newCommentMessage = "";

        $scope.addComment = function(studentKey) {
            if($scope.newCommentMessage !== "") {
                ProfileService.addComment(studentKey,$scope.newCommentMessage, function (comment, err) {
                    if(err) {
                        //error
                    } else {
                        $scope.comments.push(comment);
                        $scope.newCommentMessage = "";
                    }
                });
            }
        };
        $scope.removeComment = function(commentKey) {
            ProfileService.removeComment(commentKey, function (err) {
                if(err) {
                    //error
                } else {
                    //remove comment from the active scope
                    $scope.comments = $scope.comments.filter(function (comment) {
                        return comment.key !== commentKey;
                    });
                    $("#"+commentKey).remove();
                }
            });
        };
        $scope.editComment = function (comment) {
            let editModal = $('#editModal');
            $("#editComment").text(comment.message);
            $('#editForm').form({
                fields: {
                    editCommentMessage: {
                        identifier: 'editCommentMessage',
                        rules: [
                            {
                                type: 'empty',
                                prompt: 'Please set message'
                            }
                        ]
                    },
                },
                onSuccess: function (event, fields) {
                    ProfileService.editComment(comment.key,fields.editCommentMessage, function (editedComment, err) {
                        if(err) {
                            //error
                        } else {
                            $("#"+editedComment.key+" .description").text(editedComment.message);
                        }
                    });
                },
                onFailure: function (formErrors, fields) {},
                keyboardShortcuts: false
            });
            editModal.modal({
                onShow: function () {

                },
                onDeny: function () {
                    return true;
                },
                onApprove: function () {
                    let editForm = $("#editForm");
                    editForm.form('validate form');
                    return editForm.form('is valid');
                }
            });
            editModal.modal('show');
        };

        ProfileService.getComments($scope.studentKey, function (comments,err) {
            if(err) {}
            else {
                $scope.comments = comments;
                $scope.comments.forEach(function (comment) {
                    if(sessionStorage.getItem("userKey") === comment.author) {
                        $('.'+comment.author).dimmer({
                            on: 'hover'
                        });
                    }
                });
            }
        });
    });