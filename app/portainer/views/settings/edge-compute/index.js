import angular from 'angular';
import controller from './settingsEdgeComputeController';

angular.module('opendocking.app').component('settingsEdgeComputeView', {
  templateUrl: './settingsEdgeCompute.html',
  controller,
});
