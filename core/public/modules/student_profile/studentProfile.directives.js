angular.module('Xtern')
    .directive('studentProfileComment', function() {
        return {
            restrict: 'E',
            templateUrl: 'public/modules/student_profile/comments/comment.html',
            controller: 'CommentCtrl'
        };
    })
    .directive('techLabels', function () {
        return {
            restrict: 'E',
            // scope: {
            //     labels: labels
            // },
            templateUrl: 'public/modules/student_profile/tech_labels/techLabels.html',
            controller: 'TechLabelsCtrl'
        };
    })
    .directive('studentDataPage', function() {
        return {
            restrict: 'E',
            templateUrl: 'public/modules/student_profile/partials/studentDataPage.html',
            controller: 'StudentProfileCtrl'
        };
    });