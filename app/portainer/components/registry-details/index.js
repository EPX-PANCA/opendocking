import angular from 'angular';

export const registryDetails = {
  templateUrl: './registry-details.html',
  bindings: {
    registry: '<',
  },
};

angular.module('opendocking.app').component('registryDetails', registryDetails);
