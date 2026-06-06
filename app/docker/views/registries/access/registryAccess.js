angular.module('opendocking.docker').component('dockerRegistryAccessView', {
  templateUrl: './registryAccess.html',
  controller: 'DockerRegistryAccessController',
  bindings: {
    endpoint: '<',
  },
});
