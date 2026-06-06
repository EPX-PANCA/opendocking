angular.module('opendocking.app').component('productItem', {
  templateUrl: './productItem.html',
  bindings: {
    model: '<',
    goTo: '<',
  },
});
