var graphiteBeaconWeb = angular.module('graphiteBeaconWeb', ['ngRoute', 'graphiteBeaconWebControllers']);

graphiteBeaconWeb.config(['$routeProvider', function($routeProvider) {
    $routeProvider
    
      // alert routes
      .when('/alerts', {
        templateUrl: '/html/templates/alertlist.html',
        controller: 'AlertListCtrl'
      })
      .when('/alert/:id', {
        templateUrl: '/html/templates/alertdetail.html',
        controller: 'AlertDetailCtrl'
      })
      
      // configuration routes
      .when('/configurations', {
        templateUrl: '/html/templates/configurations.html',
        controller: 'ConfigurationsCtrl'
      })
      
      // catch all route
      .otherwise({
        redirectTo: '/alerts'
      });
      
  }]);
