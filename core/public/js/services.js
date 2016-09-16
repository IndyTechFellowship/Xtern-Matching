/**
 * Created by Steven on 3/28/2016.
 */
 (function (){
    var app = angular.module('DataManager',[]);

    app.service('ProfileService', ['$http', function ($http){
        var self = this;
        self.profile = null;

        // self.query = function(callback){
        //     if(!self.profile){
        //         $http({
        //             method: 'GET',
        //             url: 'data_mocks/user_profile_1.json'
        //         }).then(function(data){
        //             console.log(data);
        //             self.profile = data.data;
        //             callback(self.profile);
        //             console.log('get student data:' + data.data.length);
        //             var test = $.grep(data.data, function(e){return e.id == '57269aa3bf79bbf8cc55d9dc';});
        //             console.log(test);
        //             // id = '57269aa3bf79bbf8cc55d9dc';

        //           //   for (var i = 0; i < data.data.length; i++) {
        //           //       if (data.data[i].id === id) {
        //           //           console.log(data.data[i]);
        //           //     } else {
        //           //       console.log('nope');
        //           //     }
        //           // }
        //       });
        //     } else {
        //         callback(self.profile);
        //     }
        // };

        self.getStudentDataForId = function(id, callback){
            console.log(id);
            if(!self.profile || self.profile._id !== id){
                $http({
                    method: 'GET',
                    // url: 'public/data_mocks/user_profile_1.json'
                    url: 'public/data_mocks/StudentData5_1.json'
                }).then(function(data){
                    console.log('get student data:' + data.data.length);
                    // self.profile = $.grep(data.data, function(e){return e.id == '57269aa3bf79bbf8cc55d9dc';});
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
                $http({
                    method: 'GET',
                    url:'public/data_mocks/StudentData5_1.json'
                }).then(function (data) {
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
