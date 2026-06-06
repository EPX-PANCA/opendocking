angular.module('opendocking.kubernetes').component('kubernetesApplicationLogsView', {
  templateUrl: './logs.html',
  controller: 'KubernetesApplicationLogsController',
  controllerAs: 'ctrl',
  bindings: {
    $transition$: '<',
  },
});
