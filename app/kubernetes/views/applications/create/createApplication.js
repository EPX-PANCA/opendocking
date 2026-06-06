angular.module('opendocking.kubernetes').component('kubernetesCreateApplicationView', {
  templateUrl: './createApplication.html',
  controller: 'KubernetesCreateApplicationController',
  controllerAs: 'ctrl',
  bindings: {
    endpoint: '<',
  },
});
