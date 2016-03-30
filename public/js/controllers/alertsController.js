var graphiteBeaconWebAlertsController = angular.module('graphiteBeaconWebAlertsController', []);

graphiteBeaconWebAlertsController.controller('AlertListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/alerts').success(function(data) {
      $scope.alerts = data;
    });
    $scope.details = function (url) {
        $location.path(url);
    }
    $scope.delete = function (url) {
        $http.delete(url).success(function(data) {
          console.log(data);
          $http.get('api/alerts').success(function(data) {
            $scope.alerts = data;
          });
        });
    }
  }]);

graphiteBeaconWebAlertsController.controller('AlertDeletedListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/alerts/deleted').success(function(data) {
      $scope.alerts = data;
    });
    $scope.details = function (url) {
        $location.path(url);
    }
    $scope.undelete = function (url) {
        $http.put(url).success(function(data) {
          $http.get('api/alerts/deleted').success(function(data) {
            $scope.alerts = data;
          });
        });
    }
  }]);

graphiteBeaconWebAlertsController.controller('AlertDetailCtrl', ['$scope', '$http', '$location', '$routeParams',
  function ($scope, $http, $location, $routeParams) {
    $http.get('api/alert/' + $routeParams.id).success(function(data) {
      $scope.alert = data;
    });
    $scope.goBack = function() {
      window.history.back();
    };
  }]);