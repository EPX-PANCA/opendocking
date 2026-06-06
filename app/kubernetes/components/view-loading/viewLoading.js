angular.module('opendocking.kubernetes').component('kubernetesViewLoading', {
  templateUrl: './viewLoading.html',
  bindings: {
    viewReady: '<',
  },
});
