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
    $scope.user=new User();
    $scope.submit = function () {
      User.save($scope.user,function(u){
        $rootScope.msg = 'User '+u.name+' created!';
        $location.path('/users');
      },function(res){
        $scope.alert={type:'alert-danger',msg:"Unable to save contact"};
      });
    };
  });
