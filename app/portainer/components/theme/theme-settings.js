import angular from 'angular';
import controller from './theme-settings.controller';

angular.module('opendocking.app').component('themeSettings', {
  templateUrl: './theme-settings.html',
  controller,
});
