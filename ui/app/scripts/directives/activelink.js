'use strict';

/**
 * @ngdoc directive
 * @name uiApp.directive:activeLink
 * @description
 * # activeLink
 */
angular.module('uiApp')
.directive('activeLink', function ($location) {
  return {
    restrict: 'A',
  link: function postLink(scope, element, attrs, controller) {
    scope.$watch(function() {
      return $location.path();
    }, function(newValue, oldValue) {

      $('li[data-match-route]', element).each(function(k, li) {
        var $li = angular.element(li),
      // data('match-rout') does not work with dynamic attributes
      pattern = $li.attr('data-match-route'),
      regexp = new RegExp('^' + pattern + '$', ['i']);
      if(regexp.test(newValue)) {
        $li.addClass('active');
      } else {
        $li.removeClass('active');
      }

      });
    });
  }
  }
  });
