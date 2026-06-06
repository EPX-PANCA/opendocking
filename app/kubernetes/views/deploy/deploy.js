angular.module('opendocking.kubernetes').component('kubernetesDeployView', {
  templateUrl: './deploy.html',
  controller: 'KubernetesDeployController',
  controllerAs: 'ctrl',
  bindings: {
    endpoint: '<',
  },
});
