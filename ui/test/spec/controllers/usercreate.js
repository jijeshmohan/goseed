'use strict';

describe('Controller: UsercreateCtrl', function () {

  // load the controller's module
  beforeEach(module('uiApp'));

  var UsercreateCtrl,
  mockUser,
  httpBackend,
  location,
  scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope,$httpBackend,User,$location) {
    scope = $rootScope.$new();
    httpBackend=$httpBackend;
    mockUser=User;
    location=$location;
    UsercreateCtrl = $controller('UsercreateCtrl', {
      $scope: scope,
      'User': mockUser
    });
  }));

   afterEach(function() {
        httpBackend.verifyNoOutstandingExpectation();
        httpBackend.verifyNoOutstandingRequest();
    }); 

  it('should create new user', function () {
    httpBackend.expectPOST('/users', {name: 'new_user',email: 'new@email.com'}).respond({ id: 200, name: 'new_user',email: 'new@email.com' });
    scope.user=new mockUser();
    scope.user.name="new_user";
    scope.user.email="new@email.com";
    scope.submit();
    httpBackend.flush();
    expect(location.path()).toBe('/users');
    // expect(mockUser.save).toHaveBeenCalled();
  });
});
