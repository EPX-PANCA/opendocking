import angular from 'angular';
import controller from './registryController';

angular.module('opendocking.app').component('editRegistry', {
  templateUrl: './registry.html',
  controller,
  bindings: {
    $transition$: '<',
  },
});
