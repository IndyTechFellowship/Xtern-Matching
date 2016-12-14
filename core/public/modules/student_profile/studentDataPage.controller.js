angular.module('Xtern')
    .controller('StudentDataPageCtrl', function($scope, $location, ProfileService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

    $scope.studentData = null;

    ProfileService.getStudentData($stateParams.key, function(data) {
        $scope.studentData = data;
        $scope.studentData.key = $stateParams.key;
        //PDFObject.embed(data.resume, "#example1");
        //https://gist.github.com/fcingolani/3300351
        PDFJS.disableWorker = true;
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

    $('.ui.sticky').sticky({
        context: '#example1'
    });

    $(function () {
        $('.ui.dropdown').dropdown();
    });
});