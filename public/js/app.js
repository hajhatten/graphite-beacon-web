var graphiteBeaconWeb = angular.module('graphiteBeaconWeb', ['ngRoute', 'graphiteBeaconWebControllers']);

graphiteBeaconWeb.config(['$routeProvider', function($routeProvider) {
    $routeProvider
    
      // alert routes
      .when('/alerts', {
        templateUrl: '/html/templates/alertList.html',
        controller: 'AlertListCtrl'
      })
      .when('/alert/:id', {
        templateUrl: '/html/templates/alertDetail.html',
        controller: 'AlertDetailCtrl'
      })
      
      // configuration routes
      .when('/configurations', {
        templateUrl: '/html/templates/configurationList.html',
        controller: 'ConfigurationListCtrl'
      })
      .when('/configuration/:id', {
        templateUrl: '/html/templates/configurationDetail.html',
        controller: 'ConfigurationDetailCtrl'
      })
      
      // catch all route
      .otherwise({
        redirectTo: '/alerts'
      });
      
  }]);