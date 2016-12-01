angular.module('Xtern')
    .controller('StudentProfileCtrl', function($scope, $location, ProfileService, CompanyService, $stateParams) {
    $('.ui.dropdown').dropdown();//activites semantic dropdowns

    $scope.comment = {};
    $scope.studentData = null;

    ProfileService.getStudentDataForId($stateParams._id, function(data)
    {
        $scope.studentData = data;
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

    // TODO: turn this into point grade
    $scope.statusOptions = [
        'Stage 1 Approved',
        'Stage 2 Approved',
        'Stage 3 Approved',
        'Undecided',
        'Rejected (Stage 1)',
        'Rejected (Stage 2)',
        'Rejected (Stage 3)'
    ];

    $scope.r1GradeOptions = [
        {
            "text":"A",
            "value":4
        },{
            "text":"B+",
            "value":3.5
        },{
            "text":"B",
            "value":3
        },{
            "text":"B-",
            "value":2.8
        },{
            "text":"C",
            "value":2
        },{
            "text":"D",
            "value":1
        }];

    $('.ui.sticky').sticky({
        context: '#example1'
    });

    $(function () {
        $('.ui.dropdown').dropdown();
    });

    $scope.selectStatus = function(option) {
        $scope.studentData.status = option;
    };

    $scope.selectR1Grade = function(option) {
        $scope.studentData.r1Grade = option;
    };

    $scope.addComment = function(){
        // TODO: fix/update this for new data format
        var author_name = "controller test author";
        var group_name = "controller test group";
        var text = "controller test text bla bla bla. bla bla bla.";

        ProfileService.addCommentToStudent($scope.studentData._id, author_name, group_name, text, function (data) {
            // console.log(data);
        });

        $scope.comment.author = 'test user'; //temporary
        $scope.comment.group = 'test users'; //temporary
        var newComment = angular.copy($scope.comment);
        $scope.studentData.comments.push(newComment);
    };

    $scope.removeComment = function(commentToRemove) {
        var author_name = "controller test author";
        var group_name = "controller test group";
        var text = "controller test text bla bla bla. bla bla bla.";

        ProfileService.removeCommentFromStudent($scope.studentData._id, author_name, group_name, text, function (data) {
            // console.log(data);git
        });

        // TODO: fix/update this for new data format
        for(var i = $scope.studentData.comments.length - 1; i >= 0; i--){
            if($scope.studentData.comments[i].text == text){
                $scope.studentData.comments.splice(i,1);
            }
        }
    };

    $scope.addStudent = function (_id) {
        console.log("add student:");
        console.log(_id);
        CompanyService.addStudentToWishList(_id, function(data) {
            // $scope.recruitmentList.push($scope.studentData);
            console.log("Student added");
        });

    };
})