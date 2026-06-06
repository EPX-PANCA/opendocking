angular.module('opendocking.docker').component('nodeDetailsView', {
  templateUrl: './node-details-view.html',
  controller: 'NodeDetailsViewController',
  bindings: {
    endpoint: '<',
  },
});
