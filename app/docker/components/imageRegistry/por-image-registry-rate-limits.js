import angular from 'angular';

import controller from './por-image-registry-rate-limits.controller';

angular.module('opendocking.docker').component('porImageRegistryRateLimits', {
  bindings: {
    endpoint: '<',
    registry: '<',
    setValidity: '<',
    isAdmin: '<',
    isDockerHubRegistry: '<',
    isAuthenticated: '<',
    registryId: '<',
  },
  controller,
  transclude: {
    rateLimitExceeded: '?porImageRegistryRateLimitExceeded',
  },
  templateUrl: './por-image-registry-rate-limits.html',
});
