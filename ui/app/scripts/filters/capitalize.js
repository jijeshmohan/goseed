'use strict';

/**
 * @ngdoc filter
 * @name uiApp.filter:capitalize
 * @function
 * @description
 * # capitalize
 * Filter in the uiApp.
 */
angular.module('uiApp')
  .filter('capitalize', function () {
    return function (input,all) {
      return (!!input) ? input.replace(/([^\W_]+[^\s-]*) */g, function(txt){return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();}) : '';
    };
  });
