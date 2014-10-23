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
      User.save(user,function(u){
        $rootScope.msg = 'User '+u.name+' created!';
        $location.path('/users');
      },function(res){
        $scope.alert={type:'alert-danger',msg:"Unable to save contact"};
      });
    };
  });
