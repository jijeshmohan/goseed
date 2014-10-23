'use strict';

/**
 * @ngdoc function
 * @name uiApp.controller:UserCtrl
 * @description
 * # UserCtrl
 * Controller of the uiApp
 */
angular.module('uiApp')
.controller('UserCtrl', function ($scope,User,$http,$rootScope) {
  $scope.error= $rootScope.msg;
  $rootScope.msg=null;
  $scope.removeUser = function (id) {
    $scope.users.forEach(function(user, index) {
      if (id === user.id) {
        user.$delete({id: user.id},function(){
          $scope.users.splice(index, 1); 
        });
      }
    });
  };

  User.query(function(data) {
    $scope.users = data;
  },function(res){
    $scope.error="Unable to get users list from server";
  });
});
