angular.module('opendocking.kubernetes').component('kubernetesApplicationStatsView', {
  templateUrl: './stats.html',
  controller: 'KubernetesApplicationStatsController',
  controllerAs: 'ctrl',
  bindings: {
    $transition$: '<',
  },
});
