'use strict';

/**
 * @ngdoc function
 * @name uiApp.controller:UsercreateCtrl
 * @description
 * # UsercreateCtrl
 * Controller of the uiApp
 */
angular.module('uiApp')
  .controller('UsercreateCtrl', function ($scope,User, $location, $rootScope) {
    $scope.submit = function () {
      var user = new User($scope.user);
      user.$save(function(u, putResponseHeaders) {
        $rootScope.msg = 'User '+u.Name+' created!';
        $location.path('/users');
      },function(res){
        $scope.alert={type:'alert-danger',msg:"Unable to save contact"};
      });
    };
  });
