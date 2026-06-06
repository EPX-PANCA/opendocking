import angular from 'angular';
import controller from './oauth-provider-selector.controller';

angular.module('opendocking.oauth').component('oauthProvidersSelector', {
  templateUrl: './oauth-providers-selector.html',
  bindings: {
    onChange: '<',
    value: '<',
  },
  controller,
});
