'use strict';

describe('Controller: UserCtrl', function () {

  // load the controller's module
  beforeEach(module('uiApp'));

  var UserCtrl,
  scope,
  mockUser;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope,$q) {
    scope = $rootScope.$new();
    mockUser = {
      query: function() {
        var queryDeferred = $q.defer();
        return {$promise: queryDeferred.promise};
      }
    }
    spyOn(mockUser, 'query').andCallThrough();
    UserCtrl = $controller('UserCtrl', {
      $scope: scope,
      'User': mockUser
    });
  }));

  it('should have user query', function () {
    expect(mockUser.query).toHaveBeenCalled();
  });
});
