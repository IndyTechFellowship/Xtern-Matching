angular.module('Xtern')
    .controller('TechLabelsCtrl', function($scope) {
        $scope.colorForLanguage = function(category) {
            if(category === "Full-Stack") {
                return 'red';
            } else if (category === "Front-End") {
                return 'green';
            } else if (category === "Mobile") {
                return 'blue';
            } else if (category === "General") {
                return 'yellow';
            } else if (category === "Database") {
                return 'purple';
            } else {
                return 'black';
            }
        };
    });