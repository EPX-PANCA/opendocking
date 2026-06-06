angular.module('opendocking.kubernetes').component('kubernetesSummaryView', {
  templateUrl: './summary.html',
  controller: 'KubernetesSummaryController',
  controllerAs: '$ctrl',
  bindings: {
    formValues: '<',
    oldFormValues: '<',
  },
});
