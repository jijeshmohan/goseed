
angular.module('uiApp')
	.factory("User", function($resource) {
  return $resource("/users/:id",{id: '@Id'},{
  	 'save':   {method:'POST'}
  });
});
