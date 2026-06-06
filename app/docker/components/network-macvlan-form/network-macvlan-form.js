angular.module('opendocking.docker').component('networkMacvlanForm', {
  templateUrl: './networkMacvlanForm.html',
  controller: 'NetworkMacvlanFormController',
  bindings: {
    data: '=',
    applicationState: '<',
  },
});
