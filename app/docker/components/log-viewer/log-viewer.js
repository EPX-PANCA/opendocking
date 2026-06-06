angular.module('opendocking.docker').component('logViewer', {
  templateUrl: './logViewer.html',
  controller: 'LogViewerController',
  bindings: {
    data: '=',
    displayTimestamps: '=',
    logCollectionChange: '<',
    sinceTimestamp: '=',
    lineCount: '=',
    resourceName: '<',
  },
});
