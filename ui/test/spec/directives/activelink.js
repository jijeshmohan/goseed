'use strict';

describe('Directive: activeLink', function () {

  // load the directive's module
  beforeEach(module('uiApp'));

  var element,
    scope;

  beforeEach(inject(function ($rootScope) {
    scope = $rootScope.$new();
  }));

  it('should make element active using route', inject(function ($compile) {
    element = angular.element('<div active-link><ul><li data-match-route="/users"></li><li data-match-route="/"><li></ul></div>');
    element = $compile(element)(scope);
    expect(element.find('ul').length).toEqual(1);
  }));
});
