angular.module('Xtern')
    .controller("UploadCtrl", function($scope,ResumeService){
		this.uploadResume = function(id) {
			ResumeService.uploadResume(id);
		}
	});

