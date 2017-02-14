'use strict';
angular.module('Xtern')
    .controller('StudentDataPageCtrl', function($scope, $location, ProfileService, $stateParams) {
    $scope.studentData = null;

    ProfileService.getStudent($stateParams.key, function(data) {
        $scope.studentData = data;
        $scope.studentData.key = $stateParams.key;
        PDFJS.disableWorker = true;
        //TODO make this loaded
        PDFJS.workerSrc = "node_modules/pdfjs-dist/build/pdf.worker.js";

        function renderPage(page) {
            var viewport = page.getViewport(1.1);
            var canvas = document.createElement('canvas');
            var ctx = canvas.getContext('2d');
            var renderContext = {
                canvasContext: ctx,
                viewport: viewport
            };

            canvas.height = viewport.height;
            canvas.width = viewport.width;
            canvas.style.overflow = "hidden";
            document.getElementById("example1").appendChild(canvas);

            page.render(renderContext);
        }
        function renderPages(pdfDoc){
            for(var num = 1; num <= pdfDoc.numPages; num++)
                pdfDoc.getPage(num).then(renderPage);
        }
        PDFJS.getDocument(data.resume).then(renderPages);
    });

    $scope.$on('$viewContentLoaded', function (evt) {
         $('.ui.dropdown').dropdown();
         $('.ui.sticky').sticky({
             context: '#example1'
         });
    });
});