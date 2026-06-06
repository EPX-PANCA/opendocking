import angular from 'angular';
import CreateRegistryController from './createRegistryController';

angular.module('opendocking.app').component('createRegistry', {
  templateUrl: './createRegistry.html',
  controller: CreateRegistryController,
  bindings: {
    $transition$: '<',
  },
});
