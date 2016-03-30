var graphiteBeaconWebConfigurationsController = angular.module('graphiteBeaconWebConfigurationsController', []);

graphiteBeaconWebConfigurationsController.controller('ConfigurationListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/configurations').success(function(data) {
      $scope.configurations = data;
    });
    $scope.show = function (url) {
        $location.path(url);
    }
    $scope.delete = function (url) {
        $http.delete(url).success(function(data) {
          console.log(data);
          $http.get('api/configurations').success(function(data) {
            $scope.configurations = data;
          });
        });
    }
  }]);
  
graphiteBeaconWebConfigurationsController.controller('ConfigurationDeletedListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/configurations/deleted').success(function(data) {
      $scope.configurations = data;
    });
    $scope.show = function (url) {
        $location.path(url);
    }
    $scope.undelete = function (url) {
        $http.put(url).success(function(data) {
          $http.get('api/configurations/deleted').success(function(data) {
            $scope.configurations = data;
          });
        });
    }
  }]);
  
graphiteBeaconWebConfigurationsController.controller('ConfigurationDetailCtrl', ['$scope', '$http', '$location', '$routeParams',
  function ($scope, $http, $location, $routeParams) {
    $http.get('api/configuration/' + $routeParams.id).success(function(data) {
      $scope.conf = data;
    });
    $scope.goBack = function() {
      window.history.back();
    };
  }]);