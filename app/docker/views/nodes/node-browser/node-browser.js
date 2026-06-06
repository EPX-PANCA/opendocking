angular.module('opendocking.docker').component('nodeBrowserView', {
  templateUrl: './node-browser.html',
  controller: 'NodeBrowserController',
  bindings: {
    endpoint: '<',
  },
});
