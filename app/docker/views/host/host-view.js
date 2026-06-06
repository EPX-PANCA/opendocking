angular.module('opendocking.docker').component('hostView', {
  templateUrl: './host-view.html',
  controller: 'HostViewController',
  bindings: {
    endpoint: '<',
  },
});
