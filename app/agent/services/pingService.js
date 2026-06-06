import angular from 'angular';

angular.module('opendocking.agent').service('AgentPingService', AgentPingService);

function AgentPingService(AgentPing) {
  return { ping };

  function ping(endpointId) {
    return AgentPing.ping({ endpointId }).$promise;
  }
}
