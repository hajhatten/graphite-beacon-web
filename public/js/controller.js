var graphiteBeaconWebControllers = angular.module('graphiteBeaconWebControllers', []);

graphiteBeaconWebControllers.controller('AlertListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/alerts').success(function(data) {
      $scope.alerts = data;
    });
    $scope.go = function (url) {
        $location.path(url);
    }
    $scope.delete = function (url) {
        $http.delete(url).success(function(data) {
          console.log(data);
          $http.get('api/alerts').success(function(data) {
            $scope.configurations = data;
          });
        });
    }
  }]);

graphiteBeaconWebControllers.controller('AlertDetailCtrl', ['$scope', '$http', '$location', '$routeParams',
  function ($scope, $http, $location, $routeParams) {
    $http.get('api/alert/' + $routeParams.id).success(function(data) {
      $scope.alert = data;
    });
    $scope.go = function (url) {
        $location.path(url);
    }
  }]);

graphiteBeaconWebControllers.controller('ConfigurationListCtrl', ['$scope', '$http', '$location',
  function ($scope, $http, $location) {
    $http.get('api/configurations').success(function(data) {
      $scope.configurations = data;
    });
    $scope.go = function (url) {
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
  
graphiteBeaconWebControllers.controller('ConfigurationDetailCtrl', ['$scope', '$http', '$location', '$routeParams',
  function ($scope, $http, $location, $routeParams) {
    $http.get('api/configuration/' + $routeParams.id).success(function(data) {
      $scope.conf = data;
    });    
  }]);