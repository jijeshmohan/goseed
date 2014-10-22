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
  $scope.addUser = function () {
    var user = new User($scope.user);
    user.$save(function(u, putResponseHeaders) {
      $scope.users.push(u);
      $scope.user.Name = '';
      $scope.user.Email = '';	
    });
  };

  $scope.removeUser = function (id) {
    $scope.users.forEach(function(user, index) {
      if (id === user.Id) {
        user.$delete({id: user.Id},function(){
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
