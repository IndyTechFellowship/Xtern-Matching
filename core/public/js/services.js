/**
 * Created by Steven on 3/28/2016.
 */
 (function (){
    var app = angular.module('DataManager',[]);
	//var host = "http://xtern-matching.appspot.com/student/"
	//var host = "http://xtern-matching-143216.appspot.com/student/" //DEV Server
	var host = "http://localhost:8080/student/"
    app.service('ProfileService', ['$http', function ($http){
        var self = this;
        self.profile = null;

        self.getStudentDataForId = function(id, callback){
            console.log(id);
            if(!self.profile || self.profile._id !== id) {
                $http.get(host+id).then(function(data) {
                    console.log('get student data:' + data.data.length);
                    self.profile = $.grep(data.data, function(e){return e._id == id;})[0];
                    callback(self.profile);
              });
            } else {
                callback(self.profile);
            }
        };

    }]).service('TechPointDashboardService',['$http', function ($http){
        var self = this;
        self.userSummaryData = null;

        self.queryUserSummaryData = function(callback){
            if(!self.userSummaryData){
                $http.get(host).then(function (data) {
                    self.userSummaryData = data.data;
                    callback(self.userSummaryData);
                });
            }
            else{
                callback(self.userSummaryData);
            }
        };
    }]);
})();
