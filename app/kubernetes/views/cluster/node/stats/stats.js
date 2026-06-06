angular.module('opendocking.kubernetes').component('kubernetesNodeStatsView', {
  templateUrl: './stats.html',
  controller: 'KubernetesNodeStatsController',
  controllerAs: 'ctrl',
  bindings: {
    $transition$: '<',
  },
});
