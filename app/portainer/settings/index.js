import angular from 'angular';

import authenticationModule from './authentication';
import generalModule from './general';

export default angular.module('opendocking.settings', [authenticationModule, generalModule]).name;
