var graphiteBeaconWebroutes = angular.module('graphiteBeaconWebRoutes', ['ngRoute', 'graphiteBeaconWebAlertsController', 'graphiteBeaconWebConfigurationsController']);

graphiteBeaconWebroutes.config(['$routeProvider', function($routeProvider) {
    $routeProvider

      // alert routes
      .when('/alerts', {
        templateUrl: '/html/templates/alertsList.html',
        controller: 'AlertListCtrl'
      })
      .when('/alerts/deleted', {
        templateUrl: '/html/templates/alertsListDeleted.html',
        controller: 'AlertDeletedListCtrl'
      })
      .when('/alert/:id', {
        templateUrl: '/html/templates/alertDetail.html',
        controller: 'AlertDetailCtrl'
      })

      // configuration routes
      .when('/configurations', {
        templateUrl: '/html/templates/configurationsList.html',
        controller: 'ConfigurationListCtrl'
      })
      .when('/configurations/deleted', {
        templateUrl: '/html/templates/configurationsListDeleted.html',
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
