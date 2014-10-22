
angular.module('uiApp')
	.factory("User", function($resource) {
  return $resource("/users/:id",{},{
  	 'save':   {method:'POST'}
  });
});
