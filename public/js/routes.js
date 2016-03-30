var graphiteBeaconWebroutes = angular.module('graphiteBeaconWebRoutes', ['ngRoute', 'graphiteBeaconWebAlertsController', 'graphiteBeaconWebConfigurationsController']);

graphiteBeaconWebroutes.config(['$routeProvider', function($routeProvider) {
    $routeProvider
    
      // alert routes
      .when('/alerts', {
        templateUrl: '/html/templates/alertList.html',
        controller: 'AlertListCtrl'
      })
      .when('/alerts/deleted', {
        templateUrl: '/html/templates/alertListDeleted.html',
        controller: 'AlertDeletedListCtrl'
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
      .when('/configurations/deleted', {
        templateUrl: '/html/templates/configurationListDeleted.html',
        controller: 'ConfigurationDeletedListCtrl'
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