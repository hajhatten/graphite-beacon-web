var graphiteBeaconWebControllers = angular.module('graphiteBeaconWebControllers', []);

graphiteBeaconWebControllers.controller('AlertListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/alerts').success(function(data) {
      $scope.alerts = data;
    });
  }]);

graphiteBeaconWebControllers.controller('AlertDetailCtrl', ['$scope', '$http', '$location', '$routeParams',
  function ($scope, $http, $location, $routeParams) {
    $http.get('api/alert/' + $routeParams.id).success(function(data) {
      $scope.alert = data;
    });
  }]);

graphiteBeaconWebControllers.controller('ConfigurationsCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/configurations').success(function(data) {
      $scope.configurations = data;
    });
  }]);
  
