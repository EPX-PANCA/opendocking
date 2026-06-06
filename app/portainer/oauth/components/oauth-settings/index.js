import angular from 'angular';
import controller from './oauth-settings.controller';

angular.module('opendocking.oauth').component('oauthSettings', {
  templateUrl: './oauth-settings.html',
  bindings: {
    settings: '=',
    teams: '<',
    onSaveSettings: '<',
    saveButtonState: '<',
  },
  controller,
});
