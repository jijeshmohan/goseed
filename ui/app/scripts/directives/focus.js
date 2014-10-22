'use strict';

/**
 * @ngdoc directive
 * @name uiApp.directive:focus
 * @description
 * # focus
 */
angular.module('uiApp')
  .directive('focus', function ($timeout) {
      return {
        scope : {
          trigger : '@focus'
        },
        link : function(scope, element) {
          scope.$watch('trigger', function(value) {
            if (value === "true") {
              $timeout(function() {
                element[0].focus();
              });
            }
          });
        }
      };
  });
