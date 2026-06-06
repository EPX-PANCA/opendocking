import angular from 'angular';

import { VolumeBrowserController } from './volumeBrowserController';

angular.module('opendocking.agent').component('volumeBrowser', {
  templateUrl: './volumeBrowser.html',
  controller: VolumeBrowserController,
  bindings: {
    volumeId: '<',
    nodeName: '<',
    isUploadEnabled: '<',
    endpointId: '<',
  },
});
