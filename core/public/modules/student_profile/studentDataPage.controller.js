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
            let viewport = page.getViewport(1.1);
            let canvas = document.createElement('canvas');
            let ctx = canvas.getContext('2d');
            let renderContext = {
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
            for(let num = 1; num <= pdfDoc.numPages; num++)
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