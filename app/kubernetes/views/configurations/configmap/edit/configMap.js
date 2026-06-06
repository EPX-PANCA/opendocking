angular.module('opendocking.kubernetes').component('kubernetesConfigMapView', {
  templateUrl: './configMap.html',
  controller: 'KubernetesConfigMapController',
  controllerAs: 'ctrl',
  bindings: {
    $transition$: '<',
  },
});
