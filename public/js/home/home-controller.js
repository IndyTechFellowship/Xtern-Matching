'use strict';
var app = angular.module('Xtern-Matching');
app.controller('HomeController', [function () {
  var self = this;
  self.sampleVariable = 'Loading...';
  $http.get('/api/v1/jsonInfo').then(function (result) {
    self.sampleVariable = result.data.stam;
  });
}]);
